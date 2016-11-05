package model

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/satori/go.uuid"
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
	ID          uuid.UUID `json:"id"`
	Parent      *Task     `json:"parent"`
	Children    []*Task   `json:"children"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	Level       int       `json:"level"`
	Status      Status    `json:"status"`
}

//return a new top-level task with no parent, children, or siblings
//generate a new UUID if one isn't passed in
func NewTask(desc string, summary string, status Status, uuidstr string) *Task {
	var uu uuid.UUID
	if uuidstr == "" {
		uu = uuid.NewV4()
	} else {
		uu, _ = uuid.FromString(uuidstr)
	}
	task := Task{
		ID:          uu,
		Parent:      nil,
		Children:    nil,
		Description: desc,
		Summary:     summary,
		Level:       0,
		Status:      status,
	}
	return &task
}

func StartTask(t *Task) {
	t.Status.Started = true
	t.Status.Modified = time.Now()
}

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

func NewStatus(done, started bool, due *time.Time) Status {
	now := time.Now()
	status := Status{
		Done:     done,
		Started:  started,
		Created:  now,
		Modified: now,
	}
	if due != nil {
		status.Due = *due
	}

	return status
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

/* Add a parent, child, or sibling of a task. If child, and there are other
 * children already, add this to the end of the siblings.
 */
func AddTask(node, newNode *Task, newNodeType int) {
	if node == nil || newNode == nil {
		// TODO: log an error
		return
	}
	if newNodeType == NODE_CHILD { // newNode is a child of node
		newNode.Parent = node
		node.Children = append(node.Children, newNode)
	} else if newNodeType == NODE_PARENT { // newNode is the parent of node
		// TODO: finish this -- error if it already has a parent
	} else if newNodeType == NODE_SIBLING { // newNode is a sibling of node
		// TODO: finish this
	} else { // newNodeType is a NONE or bogus
		// TODO: log an error
	}
}

// TODO: remove this node and change parent (if any) and siblings (if any)
// and children (if any -- or should this take a recursive flag?)
func RemoveTask(node *Task) {
}

// TODO: make this a method
func TaskToString(task *Task) string {
	// test and debug, for now, by returning a dummy value if nothing is passed in
	if task == nil {
		task = &Task{
			ID:          uuid.NewV4(),
			Description: "an example description\nthis one has two lines",
			Summary:     "a task",
			Status: Status{
				Done: false,
			},
		}
	}
	taskstr, _ := json.Marshal(task)
	if os.Getenv("DEBUG") != "" {
		fmt.Println(string(taskstr))
	}
	return string(taskstr)
}
