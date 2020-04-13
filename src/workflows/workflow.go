package workflows

import (
	"github.com/syedmrizwan/orchestrator/src/activities"
	"go.uber.org/cadence/workflow"
	"time"
)

func init() {
	workflow.Register(DemoWorkflow)
}

//DemoWorkflow executes three demo purpose activities
func DemoWorkflow(ctx workflow.Context) (*activities.JSONResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 10 * time.Second,
		StartToCloseTimeout:    10 * time.Second,
		//ScheduleToCloseTimeout: 10 * time.Second,
		// HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var name string
	err := retry(func() error {
		return workflow.ExecuteActivity(ctx, activities.GetNameActivity).Get(ctx, &name)
	})
	if err != nil {
		return nil, err
	}

	var result string
	err = retry(func() error {
		return workflow.ExecuteActivity(ctx, activities.SayHello, name).Get(ctx, &result)
	})
	if err != nil {
		return nil, err
	}

	var persistResult activities.JSONResponse
	err = retry(func() error {
		return workflow.ExecuteActivity(ctx, activities.PersistResult, result).Get(ctx, &persistResult)
	})
	if err != nil {
		return nil, err
	}

	logger := workflow.GetLogger(ctx)
	logger.Info("Result: " + result)
	logger.Info("Workflow completed for WF_RegisterDevicePollerMap")

	return &persistResult, nil

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
