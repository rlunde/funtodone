package main

import (
	"testing"

	"funtodone/model"

	"gopkg.in/mgo.v2/bson"
)

var (
	DropDatabase = true
)

func SetupTestTask(idStr string, desc string, summary string) *model.Task {
	status := model.Status{}
	id := bson.ObjectIdHex(idStr)
	idstr := id.Hex()
	task := model.NewTask(desc, summary, status, idstr)
	return task
}
func TestAddChildTask(t *testing.T) {
	parent := SetupTestTask("583f9a189e743bea858113ca", "parent task", "main task")
	child := SetupTestTask("583f9a189e743bea858113cb", "child task", "subtask")
	if parent == child {
		t.Errorf("something went wrong")
	}
	err := model.AddTask(parent, child, model.NodeChild)
	if err != nil {
		t.Errorf(err.Error())
	}
	s := parent.TaskToString()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		expected :=
			`{"ID":"583f9a189e743bea858113ca","children":[{"ID":"583f9a189e743bea858113cb","description":"child task","summary":"subtask","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}],"description":"parent task","summary":"main task","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}`
		if expected != s {
			t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
		}
	}
}

/*
TODO:
For all of these, don't hard-code the object ID -- just generate one
a) create an in-memory version of the current funtodone.stack
b) print out the JSON form of that to a file
c) read in the JSON form from a file
d) save it to the database
e) read it from the database
*/
