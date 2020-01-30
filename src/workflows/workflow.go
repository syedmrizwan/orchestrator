package workflows

import (
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"time"
)

func init() {
	activity.Register(getNameActivity)
	activity.Register(sayHello)
	workflow.Register(DemoWorkflow)
}
func getNameActivity() (string, error) {
	return "Cadence", nil
}

func sayHello(name string) (string, error) {
	return "Hello " + name + "!", nil
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

	workflow.GetLogger(ctx).Info("Result: " + result)

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
