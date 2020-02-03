build:
	go build -i -o bin/orchestrator src/workers/worker.go

run:
	bin/orchestrator -m worker