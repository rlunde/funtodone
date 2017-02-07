package model

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	DropDatabase = true
)

func SetupTestTask(idStr string, desc string, summary string) *Task {
	status := Status{}
	id := bson.ObjectIdHex(idStr)
	idstr := id.Hex()
	task := NewTask(desc, summary, status, idstr)
	return task
}

func TestTaskToString(t *testing.T) {
	task := SetupTestTask("583f9a189e743bea858113ca", "simple task", "do something")
	s := task.TaskToString()
	// fmt.Println(s)
	expected :=
		`{"ID":"583f9a189e743bea858113ca","parent":null,"children":null,"description":"simple task","summary":"do something","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}`
	if expected != s {
		t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
	}
}
func TestStringToTask(t *testing.T) {
	status := Status{}
	id := bson.ObjectIdHex("583f9a189e743bea858113ca")
	idstr := id.Hex()
	task := NewTask("simple task", "do something", status, idstr)
	s := task.TaskToString()
	// fmt.Println(s)
	parsedTask := DecodeTask(s)
	if parsedTask.ID != id {
		t.Errorf("expected: %s\nbut got:\n%s", id, parsedTask.ID)
	}
	// TODO: compare more values here
	// s = task.TaskToString()
	// fmt.Println(s)
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
	if string(task.ID) == "" {
		t.Errorf("expected a new UUID but didn't get one")
	}
}

func TestStartTask(t *testing.T) {
	status, _ := NewStatus(false, false, nil)
	task := NewTask("test task", "a test task", *status, "")
	curtime := time.Now()
	StartTask(task)
	s := task.Status
	elapsed := s.Modified.Sub(curtime)
	if elapsed.Nanoseconds() < 0 {
		t.Errorf("expected task.Status.Modified time to be >= current time")
	}
	if !s.Started {
		t.Errorf("expected task.Status.Started to be true")
	}
}
func TestFinishTask(t *testing.T) {
	status, _ := NewStatus(false, false, nil)
	task := NewTask("test task", "a test task", *status, "")
	curtime := time.Now()
	StartTask(task)
	FinishTask(task)
	s := task.Status
	elapsed := s.Modified.Sub(curtime)
	if elapsed.Nanoseconds() < 0 {
		t.Errorf("expected task.Status.Modified time to be >= current time")
	}
	if !s.Done {
		t.Errorf("expected task.Status.Done to be true")
	}
	elapsed = s.Completed.Sub(s.Modified)
	if elapsed.Nanoseconds() != 0 {
		t.Errorf("expected task.Status.Modified time to be == task.Status.Completed time")
	}
}
func DropDatabaseOnce(session *mgo.Session) error {
	// Drop Database the first time this is called
	if DropDatabase {
		err := session.DB("test").DropDatabase()
		DropDatabase = false
		return err
	}
	return nil
}
func GetTestDatabaseConnection(t *testing.T) (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		t.Fatalf(err.Error())
		return nil, nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	// We are only running this one test, so drop the test database firstname
	err = DropDatabaseOnce(session)
	if err != nil {
		session.Close()
		t.Fatalf(err.Error())
		return nil, nil, err
	}
	c := session.DB("test").C("tasks")
	return session, c, nil
}
func TestSaveTask(t *testing.T) {
	session, c, err := GetTestDatabaseConnection(t)
	defer session.Close()
	status, _ := NewStatus(false, false, nil)
	task := NewTask("save this task", "save task to mongo database test, collection tasks", *status, "")
	err = SaveTask(task, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestFindTask(t *testing.T) {
	// TODO: figure out if we can just call TestSaveTask to set up task -- problem is how to pass back task
	session, c, err := GetTestDatabaseConnection(t)
	defer session.Close()
	status, _ := NewStatus(false, false, nil)
	task := NewTask("find this task", "find task in mongo database test, collection tasks", *status, "")
	err = SaveTask(task, c)
	id := task.ID
	if err != nil {
		t.Errorf(err.Error())
	} else {
		newtask, err := FindTaskByID(c, id)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			if newtask.Summary != task.Summary {
				t.Errorf("expected:\n%s\nbut got:\n%s", newtask.Summary, task.Summary)
			}
		}
	}
}

func TestAddChildTask(t *testing.T) {
	fmt.Println("creating parent node")
	parent := SetupTestTask("583f9a189e743bea858113ca", "parent task", "main task")
	fmt.Println("creating child node")
	child := SetupTestTask("583f9a189e743bea858113cb", "child task", "subtask")
	if parent == child {
		t.Errorf("something went wrong")
	}
	fmt.Println("adding child node to parent node")
	err := AddTask(child, parent, NodeParent)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("converting parent node to string")
	s := child.TaskToString()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(s)
		expected :=
			`{"ID":"583f9a189e743bea858113ca","parent":null,"children":null,"description":"simple task","summary":"do something","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}`
		if expected != s {
			t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
		}
	}
}

/*
 * TODO: Tests to add
 *  [ ] error tests for TaskToString
 *  [ ] tests for AddTask
       [ ] nil node
			 [ ] nil newNode
			 [ ] add a valid child node
			 [ ] add a parent node to a node without a parent
			 [ ] error if add a parent node to a node with a parent
	     [ ] error if add a node with an invalid newNodeType

*/
