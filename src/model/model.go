package funtodone

import (
"encoding/json"
 "fmt"
 "os"
 "github.com/satori/go.uuid"
)
type Task struct {
	ID     UUID     `json:"is"`
	Description string `json:"description"`
  Summary string `json:"summary"`
}
func sample() {
  task1 := &Task{
        ID:   uuid.NewV4(),
        Description: "an example description\nthis one has two lines"
        Summary: "a task"
      }
    taskstr, _ := json.Marshal(task1)
    fmt.Println(string(taskstr))
}
