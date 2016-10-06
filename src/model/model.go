package funtodone

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
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
}

func sample() {
	task1 := &Task{
		ID:          uuid.NewV4(),
		Description: "an example description\nthis one has two lines",
		Summary:     "a task",
	}
	taskstr, _ := json.Marshal(task1)
	fmt.Println(string(taskstr))
}
