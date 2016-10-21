package model

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
)

func TestTaskString(t *testing.T) {
	task := Task{
		ID:          uuid.NewV4(),
		Parent:      uuid.UUID{},
		Children:    []uuid.UUID{},
		Previous:    uuid.UUID{},
		Next:        uuid.UUID{},
		Description: "simple task",
		Summary:     "do something",
		Level:       0,
		Status:      Status{},
	}
	s := String(&task)
	fmt.Println(s)
}
