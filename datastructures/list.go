package datastructures

import (
	"reflect"
)

type Memory interface {
	NextValue() any
	IsEmpty() bool
	Contains(value any) bool
}

type List struct {
	first *Node
	last  *Node
}

type Node struct {
	previous *Node
	next     *Node
	Value    any
}

func (l *List) Add(value any) {
	node := Node{previous: l.last, Value: value}

	if l.first == nil {
		l.first = &node
	}

	if l.last != nil {
		l.last.next = &node
	}
	l.last = &node
}

func (l List) IsEmpty() bool {
	return l.first == nil && l.last == nil
}

func (l List) Contains(value any) bool {

	for node := l.first; node != nil; node = node.next {

		if reflect.DeepEqual(value, node.Value) {
			return true
		}
	}

	return false
}

type EmptyListError struct {
	Message string
}

func (ele EmptyListError) Error() string {
	return ele.Message
}
