package workflows

import (
	"context"
	"io/ioutil"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
)

func init() {
	activity.Register(getNameActivity)
	activity.Register(sayHello)
	activity.Register(persistResult)
	workflow.Register(DemoWorkflow)
}
func getNameActivity() (string, error) {
	return "Cadence", nil
}

func sayHello(name string) (string, error) {
	return "Hello " + name + "!", nil
}
func persistResult(ctx context.Context, data string) error { // Save in DB but for now saving in file
	activityInfo := activity.GetInfo(ctx)
	// taskToken := activityInfo.TaskToken
	runID := activityInfo.WorkflowExecution.RunID
	fileName := "/home/emumba/Desktop/cadence/" + runID
	return ioutil.WriteFile(fileName, []byte(runID+"_"+data), 0666)
}
func DemoWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout:    3 * time.Second,
		// ScheduleToCloseTimeout: 10 * time.Second,
		// HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var name string
	err := retry(func() error {
		return workflow.ExecuteActivity(ctx, getNameActivity).Get(ctx, &name)
	})
	if err != nil {
		return err
	}

	var result string
	err = retry(func() error {
		return workflow.ExecuteActivity(ctx, sayHello, name).Get(ctx, &result)
	})
	if err != nil {
		return err
	}

	err = retry(func() error {
		return workflow.ExecuteActivity(ctx, persistResult, result).Get(ctx, nil)
	})
	if err != nil {
		return err
	}

	logger := workflow.GetLogger(ctx)
	logger.Info("Result: " + result)
	logger.Info("Workflow completed for WF_RegisterDevicePollerMap")

	return nil

}

func retry(op func() error) error {
	var err error
	for i := 0; i < 10; i++ {
		if err = op(); err == nil {
			return nil
		}
		// switch.err.(type){
		// 	case *workflow.Error
		// }
	}
	return err
}
