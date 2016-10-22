package model

import (
	"fmt"
	"testing"
)

func TestTaskString(t *testing.T) {
	status := Status{}
	task := NewTask("simple task", "do something", status)
	s := String(task)
	fmt.Println(s)
}
