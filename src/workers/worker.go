package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/pborman/uuid"
	"github.com/syedmrizwan/orchestrator/src/workflows"
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

const (
	configFile = "config/development.yaml"
)

const ApplicationName = "helloWorldGroup"

func startWorkers(h *common.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers("test-domain", ApplicationName, workerOptions)
}

func startWorkflow(count int) {
	var h common.SampleHelper
	h.SetupServiceConfig()
	workflowclient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build Cadence Client")
		fmt.Println("error in initiating client, exiting main routing")
		return
	}
	for i := 0; i < count; i++ {
		workflowOptions := client.StartWorkflowOptions{
			ID:                              "helloworld_" + uuid.New(),
			TaskList:                        ApplicationName,
			ExecutionStartToCloseTimeout:    time.Minute * 1,
			DecisionTaskStartToCloseTimeout: time.Minute * 1,
		}
		wf, err := workflowclient.StartWorkflow(context.Background(), workflowOptions, workflows.DemoWorkflow)
		if err != nil {
			h.Logger.Error("Failed to build Cadence Client")
			fmt.Println("error in initiating workflow, exiting main routing")
			return
		}
		h.Logger.Info("Started Workflow", zap.String("WorkflowId", wf.ID), zap.String("RunId", wf.RunID))

		// Retrieve the workflow handler
		wfrun := workflowclient.GetWorkflow(context.Background(), wf.ID, wf.RunID)
		fmt.Println("GetRunId = ", wfrun.GetRunID())

		// Blocking API
		// var result string
		// err = wfrun.Get(context.Background(), result)
		// if err != nil {
		// 	fmt.Println("error in wfrun.Get, exiting main routing")
		// 	return
		// }
		// fmt.Println("workflow returned = ", result)
	}
}
func main() {
	var mode string
	var count int
	flag.StringVar(&mode, "m", "trigger", "Mode is worker or trigger.")
	flag.IntVar(&count, "c", 3, "Count of workflow to start")
	flag.Parse()

	var h common.SampleHelper
	h.SetupServiceConfig()

	switch mode {
	case "worker":
		startWorkers(&h)
		// The workers are supposed to be long running process that should not exit.
		// Use select{} to block indefinitely for samples, you can quit by CMD+C.
		select {}
	case "trigger":
		startWorkflow(count)
	}
}
