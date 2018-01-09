package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	bson "gopkg.in/mgo.v2/bson"
)

const ( // iota is reset to 0
	//NodeNone - indicate a node that isn't a parent or child
	NodeNone = iota // == 0
	//NodeParent - indicate a parent node (e.g. when adding a node to another)
	NodeParent = iota // == 1
	//NodeChild - indicate a child node (e.g. when adding a node to another)
	NodeChild = iota // == 2
)

/*Task -- can be stand-alone, or a list, or a linear stack
 * or a tree. Or it can be some strange combination of lists
 * (using prev and next) and trees (using parent and children).
 *
 * In the short term, there will only be lists and stacks
 * trees with only the highest child expanded.
 *
 * We generate the ObjectId that mongo uses for _id ourselves, and also
 * use a pointer to keep track of references to other Tasks for efficient
 * implementation of in-memory stuff. Another approach might be to keep
 * a map of ptr -> objectId and of objectId -> ptr, and then only keep the
 * objectId in the object, but this current approach seems simpler, so far.
 */
type Task struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	ptr         *Task         // mongo didn't like (upper case) Ptr at all, even when I told it not to save/serialize it
	Parent      *Task         `json:"-"` // avoid infinite recursion when marshalling json
	Children    []*Task       `json:"children,omitempty"`
	Description string        `json:"description"`
	Summary     string        `json:"summary"`
	Level       int           `json:"level"`
	Status      Status        `json:"status"`
}

//NewTask - return a new top-level task with no parent or children
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

//StartTask - mark task as started. should this return an error if the task was already started? For now, it doesn't
func StartTask(t *Task) {
	t.Status.Started = true
	t.Status.Modified = time.Now()
}

//FinishTask - mark task as finished. should this return an error if the task was already done? For now, it doesn't
func FinishTask(t *Task) {
	t.Status.Done = true
	t.Status.Modified = time.Now()
	t.Status.Completed = t.Status.Modified
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

//NewStatus - create a status struct to use in a Task
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

/*User -- a placeholder for when we build in auth
So far, at least, we're thinking of just having collections tied to
a user, and only having tasks as part of collections. A user might move
a task from one collection to another, or copy a task from one to another,
but it would always be the same user.*/
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

//TaskStack - a "stack" is a collection of tasks with subtasks at different levels
type TaskStack struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	User        User   `json:"user"` // TODO: consider team or group
	Tasks       []Task `json:"tasks"`
}

//TaskList - a list of tasks, which aren't necessarily ordered
type TaskList struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	User        User   `json:"user"` // TODO: consider team or group
	Tasks       []Task `json:"tasks"`
}

//AddTask - Add a parent or child of a task.
func AddTask(node, newNode *Task, newNodeType int) error {
	if node == nil || newNode == nil {
		return errors.New("AddTask called with nil Task")
	}
	if newNodeType == NodeChild { // newNode is a child of node
		// fmt.Println("adding child node")
		newNode.Parent = node
		node.Children = append(node.Children, newNode)
	} else if newNodeType == NodeParent { // newNode is the parent of node
		if node.Parent != nil {
			return errors.New("AddTask can't add parent node: node already has parent")
		}
		node.Parent = newNode
		newNode.Children = append(newNode.Children, node)
	} else { // newNodeType is a NONE or bogus
		return fmt.Errorf("AddTask called with unknown newNodeType: %d", newNodeType)
	}
	return nil
}

// RemoveTask - pull a task out of links
// TODO: remove this node and change parent (if any)
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
			_ = RemoveTask(child, recursive) // we ignore errors -- should we?
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

//TaskToString - make a printable representation of a task
func (task *Task) TaskToString(prettyPrint bool) string {
	if task == nil {
		return ""
	}
	var taskstr []byte
	if prettyPrint {
		taskstr, _ = json.MarshalIndent(task, "", "    ")
	} else {
		taskstr, _ = json.Marshal(task)
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Println(string(taskstr))
	}
	return string(taskstr)
}

//DecodeTask - deserialize a task from a string
//TODO: return the error
func DecodeTask(jsonTask string) Task {
	dec := json.NewDecoder(strings.NewReader(jsonTask))
	var t Task
	err := dec.Decode(&t)
	if err != nil {
		log.Print(err)
	}
	return t
}
