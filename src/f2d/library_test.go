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
	var id bson.ObjectId
	if idStr == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(idStr)
	}
	idstr := id.Hex()
	task := model.NewTask(desc, summary, status, idstr)
	return task
}

/*
Reproduce this:
{"step":"1","description":"write tests of library - regenerate this list"},
{"step":"2","description":"write tests for REST APIs"},
{"step":"3","description":"write REST APIs"},
{"step":"4","description":"write tests of an initial web UI"},
{"step":"5","description":"design initial responsive web UI"},
{"step":"6","description":"tie UI to REST APIs"}
{"step":"7","description":"add user management / auth"}
{"step":"8","description":"make it minimally good looking"}
*/
func TestMakeInitialStack(t *testing.T) {
	taskstrs := []string{
		"write tests of library - regenerate this list",
		"write tests for REST APIs",
		"write REST APIs",
		"write tests of an initial web UI",
		"design initial responsive web UI",
		"tie UI to REST APIs",
		"add user management / auth",
		"make it minimally good looking",
	}
	var tasks []*model.Task
	parent := SetupTestTask("", "parent task", "main task")
	for i := 0; i < len(taskstrs); i++ {
		task := SetupTestTask("", taskstrs[i], "task"+string(i))
		tasks = append(tasks, task)
		err := model.AddTask(parent, tasks[i], model.NodeChild)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	s := parent.TaskToString()

	expected :=
		`{"ID":"583f9a189e743bea858113ca","children":[{"ID":"583f9a189e743bea858113cb","description":"child task","summary":"subtask","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}],"description":"parent task","summary":"main task","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}`
	if expected != s {
		t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
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
