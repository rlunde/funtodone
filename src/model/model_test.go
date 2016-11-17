package model

import (
	"encoding/json"
	"regexp"
	"testing"
	"time"
)

func TestTaskToString(t *testing.T) {
	status := Status{}
	task := NewTask("simple task", "do something", status, "d89275a6-4783-4091-863a-6ed5e361035f")
	s := TaskToString(task)
	// fmt.Println(s)
	expected :=
		`{"id":"d89275a6-4783-4091-863a-6ed5e361035f","parent":null,"children":null,"description":"simple task","summary":"do something","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}`
	if expected != s {
		t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
	}
}
func TestNewStatus(t *testing.T) {
	status, _ := NewStatus(false, false, nil)
	diff := time.Since(status.Created)
	if diff > time.Millisecond*10 {
		t.Errorf("Took > 10ms to create Status struct: %d", diff)
	}
	bs, _ := json.Marshal(status)
	s := string(bs)
	// fmt.Println(s)
	pattern := `^{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"[^"]*","modified":"[^"]*","completed":"0001-01-01T00:00:00Z"}$`
	var validStruct = regexp.MustCompile(pattern)
	if !validStruct.MatchString(s) {
		t.Errorf("pattern match expected was:\n%s\nbut got:\n%s", pattern, s)
	}
}
func TestNewStatusError(t *testing.T) {
	// should be an error if done is true but started is false
	status, err := NewStatus(true, false, nil)
	if status != nil {
		t.Errorf("returned a status when expected only an error")
	}
	if err == nil {
		t.Errorf("expected an error but didn't get one")
	}
	// decided not to call it an error if due date is in the past
}

// make sure that new task is created and has a valid status
func TestNewTask(t *testing.T) {
	status, _ := NewStatus(false, false, nil)
	task := NewTask("test task", "a test task", *status, "")
	if task.ID.String() == "" {
		t.Errorf("expected a new UUID but didn't get one")
	}
}

/*
 * TODO: Tests to add
 *  [ ] error tests for NewTask
 *  [ ] tests for StartTask
 *  [ ] tests for FinishTask
 *  [ ] error tests for TaskToString
 *  [ ] tests for AddTask
       [ ] nil node
			 [ ] nil newNode
			 [ ] add a valid child node
			 [ ] add a parent node to a node without a parent
			 [ ] error if add a parent node to a node with a parent
			 [ ] error if add a sibling to a node without a parent
			 [ ] add a sibling to a node
	     [ ] error if add a node with an invalid newNodeType

*/
