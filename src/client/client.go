package client

import (
	"github.com/syedmrizwan/orchestrator/env"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

var HostPort = env.Env.Host
var Domain = env.Env.Domain
var TaskListName = env.Env.ApplicationName
var ClientName = env.Env.ClientName
var CadenceService = env.Env.Service

func buildCadenceServiceClient() (workflowserviceclient.Interface, error) {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		return nil, err
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: ch.NewSingleOutbound(HostPort)},
		},
	})

	if err := dispatcher.Start(); err != nil {
		return nil, err
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(CadenceService)), nil
}

func GetNewCadenceClient() (client.Client, error) {
	service, err := buildCadenceServiceClient()
	if err != nil {
		return nil, err
	}
	return client.NewClient(service, Domain, &client.Options{}), nil
}

