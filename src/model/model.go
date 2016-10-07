package funtodone

import (
	"encoding/json"
	"fmt"

	"github.com/satori/go.uuid"
)

/* this can be stand-alone, or a list, or a linear stack
 * or a tree. Or it can be some strange combination of lists
 * (using prev and next) and trees (using parent and children).
 *
 * In the short term, there will only be lists, cycles (lists that loop), and stacks
 * trees with only the highest child expanded.
 */
type Task struct {
	ID          UUID   `json:"id"`
	Parent      UUID   `json:"parent"`
	Children    []UUID `json:"children"`
	Previous    UUID   `json:"prev"`
	Next        UUID   `json:"next"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Level       int    `json:"level"`
}

func String(task *Task) string {
	// test and debug, for now, by returning a dummy value if nothing is passed in
	if !task {
		task := Task{
			ID:          uuid.NewV4(),
			Description: "an example description\nthis one has two lines",
			Summary:     "a task",
		}
	}
	taskstr, _ := json.Marshal(task)
	if GetEnv("DEBUG") != "" {
		fmt.Println(string(taskstr))
	}
	return string(taskstr)
}
