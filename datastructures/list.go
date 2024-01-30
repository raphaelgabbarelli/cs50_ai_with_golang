package datastructures

type Memory interface {
	NextValue() (any, error)
	IsEmpty() bool
	Contains(value any) bool
	Add(value any)
}

type List struct {
	first *node
	last  *node
}

func (l *List) Add(value any) {
	node := node{previous: l.last, Value: value}

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
		if node.Equal(value) {
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
