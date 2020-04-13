package activities

import (
	"context"
	"go.uber.org/cadence/activity"
	"io/ioutil"
	"time"
)

type JSONResponse struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
}

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
func PersistResult(ctx context.Context, data string) (JSONResponse, error) { // Save in DB but for now saving in file
	activityInfo := activity.GetInfo(ctx)
	// taskToken := activityInfo.TaskToken
	runID := activityInfo.WorkflowExecution.RunID
	fileName := "/home/emumba/Desktop/cadence/" + runID
	time.Sleep(6 * time.Second)

	jsonResponse := JSONResponse{
		Value1: "value 1",
		Value2: "value 2",
	}
	return jsonResponse, ioutil.WriteFile(fileName, []byte(data+"\nRun ID: "+runID), 0666)
}
