package datastructures

import "testing"

func TestAdd(t *testing.T) {
	list := List{}
	list.Add(1)

	if list.first == nil {
		t.Errorf("fist node is nil")
	}

	if list.first.Value != 1 {
		t.Errorf("unexpected value of first node: expected %d , obtained %d", 1, list.first.Value)
	}

	if list.last == nil {
		t.Errorf("last node is nil")
	}

	if list.last.Value != 1 {
		t.Errorf("unexpected value of last node: expected %d , obtained %d", 1, list.last.Value)
	}

	list.Add(2)

	if list.first.Value != 1 {
		t.Errorf("unexpected value of first node: expected %d , obtained %d", 1, list.first.Value)
	}
	if list.last.Value != 2 {
		t.Errorf("unexpected value of last node: expected %d , obtained %d", 2, list.last.Value)
	}
}
