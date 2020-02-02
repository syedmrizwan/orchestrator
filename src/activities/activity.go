package activities

import (
	"context"
	"go.uber.org/cadence/activity"
	"io/ioutil"
)

func init() {
	activity.Register(GetNameActivity)
	activity.Register(SayHello)
	activity.Register(PersistResult)
}
func GetNameActivity() (string, error) {
	return "Cadence", nil
}

func SayHello(name string) (string, error) {
	return "Hello " + name + "!", nil
}
func PersistResult(ctx context.Context, data string) error { // Save in DB but for now saving in file
	activityInfo := activity.GetInfo(ctx)
	// taskToken := activityInfo.TaskToken
	runID := activityInfo.WorkflowExecution.RunID
	fileName := "/home/emumba/Desktop/cadence/" + runID
	return ioutil.WriteFile(fileName, []byte(runID+"_"+data), 0666)
}
