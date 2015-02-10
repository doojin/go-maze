package maze

import (
	"errors"
	"fmt"
)

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

// Generate generates a random maze
func (maze *Maze) Generate() {
	// Generating first row
	row := NewRow(maze.width)
	row.SetGroups(0)
	row.CreateRightBorders()
	row.CreateBottomBorders()
	maze.Rows = append(maze.Rows, row)

	// Generating other rows
	for i := 1; i < maze.height; i++ {
		nextRow := maze.Rows[i-1].Copy()
		nextRow.RemoveRightBorders()
		nextRow.RemoveGroupsAndBottomBordersIfNeed()
		nextRow.SetGroups(i * maze.width)
		nextRow.CreateRightBorders()
		nextRow.CreateBottomBorders()
		maze.Rows = append(maze.Rows, nextRow)
	}

	maze.Rows[len(maze.Rows)-1].SetBorderBottom()
	row.SetBorderTop()
	fmt.Println(maze)
}
