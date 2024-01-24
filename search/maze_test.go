package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadMaze1(t *testing.T) {
	var maze Maze
	err := maze.Load("maze1.txt")
	assert.NoError(t, err)
	assert.NotNil(t, maze.Start) // 5;0
	assert.Equal(t, 5, maze.Start.X, "wrong starting x")
	assert.Equal(t, 0, maze.Start.Y, "wrong starting y")
	assert.NotNil(t, maze.Goal) // 0;5
	assert.Equal(t, 0, maze.Goal.X, "wrong goal x")
	assert.Equal(t, 5, maze.Goal.Y, "wrong goal y")
	assert.NotNil(t, maze.Walls)

	expectedRowValues := [7]bool{true, true, true, true, true, false, true}
	for i, v := range maze.Walls[0] {
		assert.Equal(t, expectedRowValues[i], v, "unexpected value at index %d", i)
	}

	expectedRowValues = [7]bool{true, true, true, true, false, false, true}
	for i, v := range maze.Walls[2] {
		assert.Equal(t, expectedRowValues[i], v, "unexpected value at index %d", i)
	}

}
