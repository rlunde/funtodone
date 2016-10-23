package model

import "testing"

func TestTaskString(t *testing.T) {
	status := Status{}
	task := NewTask("simple task", "do something", status, "d89275a6-4783-4091-863a-6ed5e361035f")
	s := String(task)
	// fmt.Println(s)
	expected := "{\"id\":\"d89275a6-4783-4091-863a-6ed5e361035f\",\"parent\":null,\"children\":null,\"prev\":null,\"next\":null,\"description\":\"simple task\",\"summary\":\"do something\",\"level\":0,\"status\":{\"done\":false,\"started\":\"0001-01-01T00:00:00Z\",\"due\":\"0001-01-01T00:00:00Z\",\"created\":\"0001-01-01T00:00:00Z\",\"modified\":\"0001-01-01T00:00:00Z\",\"completed\":\"0001-01-01T00:00:00Z\"}}"
	if expected != s {
		t.Errorf("expected:\n%s\nbut got:\n%s", expected, s)
	}
}
