package datastructures

import "testing"

func TestNextFromQueue(t *testing.T) {
	queue := Queue{}

	queue.Add(1)
	queue.Add(2)
	queue.Add("three")

	if queue.IsEmpty() {
		t.Errorf("queue is empty, expected 3 items")
	}
	v, err := queue.NextValue()
	if err != nil {
		t.Errorf("Error while getting next value: %s", err.Error())
	}
	if v, ok := v.(int32); ok && v != 1 {
		t.Errorf("Unexpected value taken from queue: expected %d, received %d", 1, v)
	}

	if queue.IsEmpty() {
		t.Errorf("queue is empty, expected 2 items")
	}
	v, err = queue.NextValue()
	if err != nil {
		t.Errorf("Error while getting next value: %s", err.Error())
	}
	if v, ok := v.(int32); ok && v != 2 {
		t.Errorf("Unexpected value taken from queue: expected %d, received %d", 2, v)
	}

	if queue.IsEmpty() {
		t.Errorf("queue is empty, expected 1 items")
	}
	v, err = queue.NextValue()
	if err != nil {
		t.Errorf("Error while getting next value: %s", err.Error())
	}
	if v, ok := v.(string); ok && v != "three" {
		t.Errorf("Unexpected value taken from queue: expected %s, received %s", "three", v)
	}

	if !queue.IsEmpty() {
		t.Errorf("queue is not empty, but all the values have been dequeued!")
	}

	if _, err := queue.NextValue(); err == nil {
		t.Errorf("NextValue on an empty queue should return an error, but none was thrown")
	}
}
