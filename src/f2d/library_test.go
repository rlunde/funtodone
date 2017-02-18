package main

import (
	"funtodone/model"
	"io/ioutil"
	"strconv"
	"testing"

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
		task := SetupTestTask("", taskstrs[i], "task"+strconv.Itoa(i))
		tasks = append(tasks, task)
		err := model.AddTask(parent, tasks[i], model.NodeChild)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	s := parent.TaskToString(true)

	baOut := []byte(s)
	err := ioutil.WriteFile("task.json", baOut, 0644)
	if err != nil {
		t.Errorf("error in WriteFile: %s\n", err.Error())
	}

	// now read it back in and see if we can parse it
	baIn, err := ioutil.ReadFile("task.json")
	if err != nil {
		t.Errorf("error in ReadFile: %s\n", err.Error())
	}

	newTask := model.DecodeTask(string(baIn))

	if newTask.Summary != "main task" {
		t.Errorf("round trip failed, expected summary to be \"main task\", but got: %s\n", newTask.Summary)
	}
}

/*
TODO:
For all of these, don't hard-code the object ID -- just generate one
x) create an in-memory version of the current funtodone.stack
x) print out the JSON form of that to a file
c) read in the JSON form from a file
d) save it to the database
e) read it from the database
*/
