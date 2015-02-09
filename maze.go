package maze

import "errors"

// Maze consists of rows
type Maze struct {
	width  int
	height int
	Rows   []Row
}

// New returns an instance of Maze
func New(width int, height int) (maze Maze, err error) {
	maze.width = width
	maze.height = height
	if height < 1 {
		err = errors.New("Height cannot be smaller than 1")
	}
	if width < 1 {
		err = errors.New("Width cannot be smaller than 1")
	}
	return maze, err
}
