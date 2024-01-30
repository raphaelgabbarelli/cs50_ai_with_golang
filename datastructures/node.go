package datastructures

type node struct {
	previous *node
	next     *node
	Value    any
}

func (a node) Equal(b any) bool {
	aEquatable, aOK := a.Value.(Equatable)
	bEquatable, bOk := b.(Equatable)

	if aOK && bOk {
		return aEquatable.Equal(bEquatable)
	} else {
		return a.Value == b
	}
}
