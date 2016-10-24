package model

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/satori/go.uuid"
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
	Previous    *Task     `json:"prev"`
	Next        *Task     `json:"next"`
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
		Previous:    nil,
		Next:        nil,
		Description: desc,
		Summary:     summary,
		Level:       0,
		Status:      status,
	}
	return &task
}

// Status keeps track of what state a task is in
type Status struct {
	Done      bool      `json:"done"`
	Started   time.Time `json:"started"`
	Due       time.Time `json:"due"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
	Completed time.Time `json:"completed"`
}

func NewStatus(done bool, due, started time.Time) {
	now := time.Now()
	status := Status{
		Done: done,
		Created: now,
		Modified: now,
	}
	if due != time.Time{} {
		status.Due = due
	}
	if started != time.Time{} {
		status.Started = started
	}
}

/* User is a placeholder for when we build in auth */
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Stack struct {
	User  User   `json:"user"` // TODO: consider team or group
	Tasks []Task `json:"tasks"`
}

func String(task *Task) string {
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
