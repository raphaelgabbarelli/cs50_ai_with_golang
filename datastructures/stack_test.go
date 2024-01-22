package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextFromStack(t *testing.T) {
	stack := Stack{}

	stack.Add(1)
	stack.Add(2)
	stack.Add("three")

	assert.False(t, stack.IsEmpty(), "There are 3 items in the stack, it's not empty!")
	v, err := stack.NextValue()
	assert.NoError(t, err, "error while getting next value")
	assert.Equal(t, "three", v)
	assert.False(t, stack.IsEmpty(), "There are 2 items in the stack, it's not empty!")
	v, err = stack.NextValue()
	assert.NoError(t, err, "error while getting next value")
	assert.Equal(t, 2, v)
	assert.False(t, stack.IsEmpty(), "there is 1 item in the stack, it's not empty")
	v, err = stack.NextValue()
	assert.NoError(t, err, "error while getting next value")
	assert.Equal(t, 1, v)

	assert.True(t, stack.IsEmpty(), "the stack is empty!")

	_, err = stack.NextValue()
	assert.Error(t, err, "NextValue on an empty stack should return an error, but none was thrown")
}
