package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	list := List{}
	list.Add(1)

	assert.NotNil(t, list.first, "first node is nil")
	assert.Equal(t, 1, list.first.Value)
	assert.NotNil(t, list.last, "last node is nil")
	assert.Equal(t, 1, list.last.Value)

	list.Add(2)

	assert.Equal(t, 1, list.first.Value)
	assert.Equal(t, 2, list.last.Value)
}

func TestContains(t *testing.T) {
	list := List{}

	elements := [4]any{1, 2, 3, "four"}

	for _, value := range elements {
		list.Add(value)
	}

	for _, value := range elements {
		assert.True(t, list.Contains(value), "List denies containing %v", value)
	}
}
