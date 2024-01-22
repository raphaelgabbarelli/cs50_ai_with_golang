package datastructures

type Queue struct {
	List
}

func (q *Queue) NextValue() (any, error) {

	if q.IsEmpty() {
		return nil, EmptyListError{Message: "Empty queue"}
	}

	nextValue := q.first.Value

	if q.first.next != nil {
		nextNode := q.first.next
		q.first = nextNode
		nextNode.previous = nil
	} else {
		q.first = nil
		q.last = nil
	}

	return nextValue, nil
}
