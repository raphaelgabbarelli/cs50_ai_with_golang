package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextFromQueue(t *testing.T) {
	queue := Queue{}

	queue.Add(1)
	queue.Add(2)
	queue.Add("three")

	assert.False(t, queue.IsEmpty(), "queue has 3 items - it's not empty!")
	v, err := queue.NextValue()
	assert.NoError(t, err, "Error while getting next value")
	assert.Equal(t, 1, v)
	assert.False(t, queue.IsEmpty(), "queue has 2 items - it's not empty!")
	v, err = queue.NextValue()
	assert.NoError(t, err, "Error while getting next value")
	assert.Equal(t, 2, v)
	assert.False(t, queue.IsEmpty(), "queue has 1 item - it's not empty!")
	v, err = queue.NextValue()
	assert.NoError(t, err, "Error while getting next value")
	assert.Equal(t, "three", v)

	assert.True(t, queue.IsEmpty(), "the queue is empty!")
	_, err = queue.NextValue()
	assert.Error(t, err, "NextValue on an empty queue should return an error, but none was thrown")
}
