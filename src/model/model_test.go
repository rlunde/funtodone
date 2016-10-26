package model

import (
	"encoding/json"
	"regexp"
	"testing"
	"time"
)

func TestTaskString(t *testing.T) {
	status := Status{}
	task := NewTask("simple task", "do something", status, "d89275a6-4783-4091-863a-6ed5e361035f")
	s := String(task)
	// fmt.Println(s)
	expected :=
		`{"id":"d89275a6-4783-4091-863a-6ed5e361035f","parent":null,"children":null,"prev":null,"next":null,"description":"simple task","summary":"do something","level":0,"status":{"done":false,"started":false,"due":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z","completed":"0001-01-01T00:00:00Z"}}`
	if expected != s {
		t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
	}
}
func TestNewStatus(t *testing.T) {
	status := NewStatus(false, false, nil)
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
