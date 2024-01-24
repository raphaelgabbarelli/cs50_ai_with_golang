package search

import (
	"bufio"
	"os"
)

type Point struct {
	X, Y int
}

type Maze struct {
	Start, Goal Point
	Walls       [][]bool
}

type MalformedMaze struct {
	Message string
}

func (m MalformedMaze) Error() string {
	return m.Message
}

func (maze *Maze) Load(filepath string) error {

	file, err := os.Open(filepath)

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineLength := -1
	hasStart := false
	hasGoal := false

	var walls [][]bool
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if lineLength == -1 {
			lineLength = len(line)
		}

		if len(line) != lineLength {
			return MalformedMaze{Message: "All lines must be of the same length"}
		}

		row := make([]bool, lineLength)

		for i, c := range line {

			row[i] = false // value for start ('A'), goal ('B'), or wall ('#')

			if c == 'A' {
				if hasStart {
					return MalformedMaze{Message: "Too many starts: only one start (A) is acceptable"}
				}
				hasStart = true
				maze.Start = Point{X: rowIndex, Y: i}
			} else if c == 'B' {
				if hasGoal {
					return MalformedMaze{Message: "Too many goals: only one goal (B) is acceptable"}
				}
				hasGoal = true
				maze.Goal = Point{X: rowIndex, Y: i}
			} else if c == '#' {
				row[i] = true
			}
		}
		walls = append(walls, row)
		rowIndex++
	}

	if !hasStart {
		return MalformedMaze{Message: "No start (A) defined"}
	}

	if !hasGoal {
		return MalformedMaze{Message: "No goal (B) defined"}
	}

	maze.Walls = walls

	return nil
}
