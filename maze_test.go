package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// New

func TestMazeRowCountShouldBeEqualToMazeHeightArgument(t *testing.T) {
	maze := New(3, 4)

	assert.Equal(t, 4, len(maze.Rows))
}

func TestMazeRowsShouldHaveCellCountEqualToMazeWidthArgument(t *testing.T) {
	maze := New(3, 4)

	assert.Equal(t, 3, len(maze.Rows[0].Cells))
	assert.Equal(t, 3, len(maze.Rows[1].Cells))
	assert.Equal(t, 3, len(maze.Rows[2].Cells))
	assert.Equal(t, 3, len(maze.Rows[3].Cells))
}
