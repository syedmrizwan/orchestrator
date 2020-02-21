package main

import (
	"context"
	"net/http"
	"github.com/syedmrizwan/orchestrator/util"
	"time"
	cadenceClient "github.com/syedmrizwan/orchestrator/src/client"
	"github.com/gin-gonic/gin"
	"go.uber.org/cadence/client"
	"github.com/syedmrizwan/orchestrator/src/workflows"
	"github.com/pborman/uuid"
)

func main() {
	r := gin.New()

	r.GET("/workflow", executeWorkflow)

	r.Run(":" + "8001")

	//r.Run()  listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func executeWorkflow(c *gin.Context) {
	logger := util.GetLogger()
	cadClient, err := cadenceClient.GetNewCadenceClient()

	if err != nil {
		logger.Error(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "helloworld_" + uuid.New(),
		TaskList:                        cadenceClient.TaskListName,
		ExecutionStartToCloseTimeout:    time.Minute * 1,
		DecisionTaskStartToCloseTimeout: time.Minute * 1,
	}

	r, err := cadClient.ExecuteWorkflow(context.Background(),
		workflowOptions,
		workflows.DemoWorkflow)

	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// for synchronous wait

	//if err := r.Get(context.Background(), nil) ; err != nil {
	//	logger.Error(err.Error())
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{
		"message":         "Demo workflow triggered successfully",
		"workflow_id":     r.GetID(),
		"workflow_run_id": r.GetRunID(),
	})

}
