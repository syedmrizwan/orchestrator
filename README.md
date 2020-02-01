
# Orchestrator

Orchestrator using Uber Cadence

# Installation

1. Install docker for your distribution

2. Download [docker-compose file](https://github.com/uber/cadence/blob/master/docker/docker-compose.yml) for Cadence Service 

3. Following steps will bring up the docker container running cadence server along with all its dependencies (cassandra, statsd, graphite). Exposes cadence frontend on port 7933 and grafana metrics frontend on port 8080. 
	```
	sudo docker-compose up
	```
	View metrics at localhost:8080/dashboard  
	View Cadence-Web at localhost:8088

4. make an alias for the Cadence cli instance :
	```
	alias cadence="sudo docker run --network=host --rm ubercadence/cli:master "
	```
5. Register a domain with Cadence Service:
	```
	cadence --do test-domain domain register -rd 1
	```
6. Describe a domain with Cadence cli
	```
	cadence --domain test-domain domain describe
	```
7. For WebUI use the following default url http://localhost:8088

  
  

# usage example

1. Run a worker process for executing workflows by compiling source in src directory
	```
	go build -o ./orchestrator

	./orchestrator -m worker
	```
2. There worker process will output verbose messages and will wait for work. Dont close window

3. Open an other terminal for accessing Cadence cli

4. Submit a workflow request 

	```

	cadence --domain test-domain workflow start --wt github.com/syedmrizwan/orchestrator/src/workflows.DemoWorkflow --tl helloWorldGroup -et 300

	```

5. Optional: Use the WebUI for viewing the Workflow execution steps