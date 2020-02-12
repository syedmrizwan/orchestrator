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

//GetNameActivity is an activty that just returns name
func GetNameActivity() (string, error) {
	return "Cadence Activity", nil
}

//SayHello is an activty that appends Hello to a string
func SayHello(name string) (string, error) {
	return "Hello " + name + "!", nil
}

//PersistResult is an activity that will be add some information to a file
func PersistResult(ctx context.Context, data string) error { // Save in DB but for now saving in file
	activityInfo := activity.GetInfo(ctx)
	// taskToken := activityInfo.TaskToken
	runID := activityInfo.WorkflowExecution.RunID
	fileName := "/home/emumba/Desktop/cadence/" + runID
	return ioutil.WriteFile(fileName, []byte(data+"\nRun ID: "+runID), 0666)
}
