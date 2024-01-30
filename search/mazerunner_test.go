package search

import (
	"testing"

	ds "github.com/raphaelgabbarelli/cs50_ai_with_golang/datastructures"
	"github.com/stretchr/testify/assert"
)

func TestStepEquality(t *testing.T) {
	pointA := Point{X: 1, Y: 2}
	a := Step{CurrentPosition: &pointA}

	pointB := Point{X: 1, Y: 2}
	b := Step{CurrentPosition: &pointB}

	assert.True(t, a.Equal(b), "simple equality failed")

	var list ds.List
	list.Add(a)

	assert.True(t, list.Contains(b), "equality through list failed")

	pointC := Point{X: 1, Y: 3}
	c := Step{CurrentPosition: &pointC}

	assert.False(t, list.Contains(c), "equality through list failed in negative case")

}

func TestSolveSimpleMaze(t *testing.T) {

	var maze Maze
	err := maze.Load("maze1.txt")
	assert.NoError(t, err, "error while loading maze1")

	var fifo ds.Queue
	var depthFirstFrontier ds.Memory
	depthFirstFrontier = &fifo

	var lifo ds.Stack
	var breadthFirstFrontier ds.Memory
	breadthFirstFrontier = &lifo

	frontiers := []ds.Memory{depthFirstFrontier, breadthFirstFrontier}

	for _, frontier := range frontiers {

		hasSolution, lastStep := Solve(&maze, &frontier)

		assert.True(t, hasSolution, "maze1 has a soltion, but none was found")
		assert.NotNil(t, lastStep, "last step is nil")

		// backtracking the solution
		assert.True(t, lastStep.CurrentPosition.Equal(Point{5, 0}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{5, 1}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{5, 2}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{4, 2}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{4, 3}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{4, 4}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{3, 4}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{2, 4}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{1, 4}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{0, 4}))
		lastStep = lastStep.PreviousStep
		assert.True(t, lastStep.CurrentPosition.Equal(Point{0, 5}))
		lastStep = lastStep.PreviousStep
		assert.Nil(t, lastStep)
	}
}

func TestSolveLargeMaze(t *testing.T) {
	var maze Maze
	err := maze.Load("maze2.txt")
	assert.NoError(t, err, "error while loading maze2")

	var fifo ds.Queue
	var depthFirstFrontier ds.Memory
	depthFirstFrontier = &fifo

	var lifo ds.Stack
	var breadthFirstFrontier ds.Memory
	breadthFirstFrontier = &lifo

	frontiers := []ds.Memory{depthFirstFrontier, breadthFirstFrontier}

	for _, frontier := range frontiers {

		hasSolution, lastStep := Solve(&maze, &frontier)
		assert.True(t, hasSolution)
		assert.NotNil(t, lastStep)
	}
}
