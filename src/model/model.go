package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2"
)

const ( // iota is reset to 0
	NODE_NONE    = iota // == 0
	NODE_PARENT  = iota // == 1
	NODE_CHILD   = iota // == 2
	NODE_SIBLING = iota // == 3
)

/* Task can be stand-alone, or a list, or a linear stack
 * or a tree. Or it can be some strange combination of lists
 * (using prev and next) and trees (using parent and children).
 *
 * In the short term, there will only be lists, cycles (lists that loop), and stacks
 * trees with only the highest child expanded.
 */

type Task struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	ptr         *Task
	Parent      *Task   `json:"parent"`
	Children    []*Task `json:"children"`
	Description string  `json:"description"`
	Summary     string  `json:"summary"`
	Level       int     `json:"level"`
	Status      Status  `json:"status"`
}

//return a new top-level task with no parent, children, or siblings
//generate a new UUID if one isn't passed in
func NewTask(desc string, summary string, status Status, idstr string) *Task {
	var id bson.ObjectId
	if idstr == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(idstr)
	}
	task := Task{
		ID:          id,
		ptr:         nil,
		Parent:      nil,
		Children:    nil,
		Description: desc,
		Summary:     summary,
		Level:       0,
		Status:      status,
	}
	task.ptr = &task
	return &task
}

// should this return an error if the task was already started? For now, it doesn't
func StartTask(t *Task) {
	t.Status.Started = true
	t.Status.Modified = time.Now()
}

// should this return an error if the task was already done? For now, it doesn't
func FinishTask(t *Task) {
	t.Status.Done = true
	t.Status.Modified = time.Now()
	t.Status.Completed = t.Status.Modified
}

func SaveTask(t *Task, c *mgo.Collection) error {
	err := c.Insert(t)
	return err
}

// Status keeps track of what state a task is in
type Status struct {
	Done      bool      `json:"done"`
	Started   bool      `json:"started"`
	Due       time.Time `json:"due"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
	Completed time.Time `json:"completed"`
}

func NewStatus(done, started bool, due *time.Time) (*Status, error) {
	now := time.Now()
	// error if it's done but not started
	if done && !started {
		err := errors.New("Status can't be done if it wasn't started")
		return nil, err
	}
	status := Status{
		Done:     done,
		Started:  started,
		Created:  now,
		Modified: now,
	}
	if due != nil {
		status.Due = *due
	}

	return &status, nil
}

/* User is a placeholder for when we build in auth */
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type TaskStack struct {
	User  User   `json:"user"` // TODO: consider team or group
	Tasks []Task `json:"tasks"`
}

type TaskList struct {
	User  User   `json:"user"` // TODO: consider team or group
	Tasks []Task `json:"tasks"`
}

type TaskCycle struct {
	User  User   `json:"user"` // TODO: consider team or group
	Tasks []Task `json:"tasks"`
}

/* Add a parent, child, or sibling of a task.
 */
func AddTask(node, newNode *Task, newNodeType int) error {
	if node == nil || newNode == nil {
		return errors.New("AddTask called with nil Task")
	}
	if newNodeType == NODE_CHILD { // newNode is a child of node
		newNode.Parent = node
		node.Children = append(node.Children, newNode)
	} else if newNodeType == NODE_PARENT { // newNode is the parent of node
		if node.Parent != nil {
			return errors.New("AddTask can't add parent node: node already has parent")
		}
		node.Parent = newNode
	} else if newNodeType == NODE_SIBLING { // newNode is a sibling of node
		if node.Parent == nil {
			return errors.New("AddTask can't add sibling node: node has no parent")
		}
		node.Parent.Children = append(node.Parent.Children, newNode)
	} else { // newNodeType is a NONE or bogus
		return fmt.Errorf("AddTask called with unknown newNodeType: %d", newNodeType)
	}
	return nil
}

// TODO: remove this node and change parent (if any) and siblings (if any)
// and children (if any -- or should this take a recursive flag?)
func RemoveTask(node *Task, recursive bool) error {
	if node == nil {
		return errors.New("RemoveTask called with nil Task")
	}
	if node.Children != nil && len(node.Children) > 0 && !recursive {
		return errors.New("RemoveTask called on Task with children, but recursive is not specified")
	}
	// remove children recursively, if any
	if len(node.Children) > 0 {
		for _, child := range node.Children {
			RemoveTask(child, recursive)
		}
	}
	parent := node.Parent
	if parent == nil {
		return nil
	}
	if len(parent.Children) == 1 {
		parent.Children = nil
	}
	// cut the node out of parent's children
	newKids := make([]*Task, len(parent.Children)-1)
	for _, child := range parent.Children {
		if child != node {
			newKids = append(newKids, child)
		}
	}
	parent.Children = newKids
	return nil
}

// TODO: make this a method
func TaskToString(task *Task) string {
	if task == nil {
		return ""
	}
	taskstr, _ := json.Marshal(task)
	if os.Getenv("DEBUG") != "" {
		fmt.Println(string(taskstr))
	}
	return string(taskstr)
}
