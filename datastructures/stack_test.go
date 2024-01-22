package datastructures

import "testing"

func TestNextFromStack(t *testing.T) {
	stack := Stack{}

	stack.Add(1)
	stack.Add(2)
	stack.Add("three")

	if stack.IsEmpty() {
		t.Errorf("queue is empty, expected 3 items")
	}
	v, err := stack.NextValue()
	if err != nil {
		t.Errorf("Error while getting next value: %s", err.Error())
	}
	if v, ok := v.(string); ok && v != "three" {
		t.Errorf("Unexpected value taken from stack: expected %s, received %s", "three", v)
	}

	if stack.IsEmpty() {
		t.Errorf("stack is empty, expected 2 items")
	}
	v, err = stack.NextValue()
	if err != nil {
		t.Errorf("Error while getting next value: %s", err.Error())
	}
	if v, ok := v.(int32); ok && v != 2 {
		t.Errorf("Unexpected value taken from stack: expected %d, received %d", 2, v)
	}

	if stack.IsEmpty() {
		t.Errorf("stack is empty, expected 1 items")
	}
	v, err = stack.NextValue()
	if err != nil {
		t.Errorf("Error while getting next value: %s", err.Error())
	}
	if v, ok := v.(int32); ok && v != 1 {
		t.Errorf("Unexpected value taken from stack: expected %d, received %d", 1, v)
	}

	if !stack.IsEmpty() {
		t.Errorf("stacl is not empty, but all the values have been popped!")
	}

	if _, err := stack.NextValue(); err == nil {
		t.Errorf("NextValue on an empty stack should return an error, but none was thrown")
	}
}
