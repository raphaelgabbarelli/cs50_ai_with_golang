package datastructures

type Stack struct {
	List
}

func (s *Stack) NextValue() (any, error) {

	if s.IsEmpty() {
		return nil, EmptyListError{Message: "Empty stack"}
	}

	nextValue := s.last.Value

	if s.last.previous != nil {
		nextNode := s.last.previous
		s.last = nextNode
		nextNode.next = nil
	} else {
		s.last = nil
		s.first = nil
	}

	return nextValue, nil
}
