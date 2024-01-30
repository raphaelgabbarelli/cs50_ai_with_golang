package search

import (
	ds "github.com/raphaelgabbarelli/cs50_ai_with_golang/datastructures"
)

type CantMoveThere struct {
	Message string
}

func (cmt CantMoveThere) Error() string {
	return cmt.Message
}

func move(maze *Maze, whereAmI Point, vectorX int, vectorY int) (Point, error) {

	nextPosition := Point{X: whereAmI.X + vectorX, Y: whereAmI.Y + vectorY}

	if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X >= len(maze.Walls[0]) || nextPosition.Y >= len(maze.Walls) {
		return whereAmI, CantMoveThere{Message: "index out of bounds"}
	}

	if maze.Walls[nextPosition.Y][nextPosition.X] {
		return whereAmI, CantMoveThere{Message: "hit a wall"}
	}

	whereAmI = nextPosition

	return whereAmI, nil
}

type Move func(*Maze, Point) (Point, error)

func up(maze *Maze, whereAmI Point) (Point, error) {
	return move(maze, whereAmI, 0, -1)
}

func down(maze *Maze, whereAmI Point) (Point, error) {
	return move(maze, whereAmI, 0, 1)
}

func left(maze *Maze, whereAmI Point) (Point, error) {
	return move(maze, whereAmI, -1, 0)
}

func right(maze *Maze, whereAmI Point) (Point, error) {
	return move(maze, whereAmI, 1, 0)
}

func Solve(maze *Maze, frontier *ds.Memory) (bool, *Step) {

	var explored ds.List
	(*frontier).Add(Step{CurrentPosition: &maze.Start})
	actions := [4]Move{up, down, left, right}

	for {
		// is the frontier empty?
		if (*frontier).IsEmpty() {
			return false, nil // TODO handle dead end
		}

		step, err := (*frontier).NextValue()
		if err != nil {
			return false, nil // handle error here
		}
		whereAmI := step.(Step)
		if whereAmI.CurrentPosition.Equal(maze.Goal) {
			return true, &whereAmI // handle solution
		}
		if explored.Contains(whereAmI) {
			continue
		}

		for _, action := range actions {
			nextStep, err := action(maze, *whereAmI.CurrentPosition)

			if err == nil {
				(*frontier).Add(Step{CurrentPosition: &nextStep, PreviousStep: &whereAmI})
			}
		}
		explored.Add(whereAmI)
	}
}
