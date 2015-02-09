package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewRow

func TestRowShouldHaveRowCountEqualToSizeArgument(t *testing.T) {
	row := NewRow(5)

	assert.Equal(t, 5, len(row.Cells))
}

func TestFirstRowShouldHaveLeftBorderLastRowShouldHaveRightBorder(t *testing.T) {
	row := NewRow(5)

	assert.Equal(t, true, row.Cells[0].BorderLeft)
	assert.Equal(t, true, row.Cells[4].BorderRight)
}

// SetBorderTop

func TestShouldSetTopBorderOfAllCellsToTrue(t *testing.T) {
	row := NewRow(3)

	assert.Equal(t, false, row.Cells[0].BorderTop)
	assert.Equal(t, false, row.Cells[1].BorderTop)
	assert.Equal(t, false, row.Cells[2].BorderTop)

	row.SetBorderTop()

	assert.Equal(t, true, row.Cells[0].BorderTop)
	assert.Equal(t, true, row.Cells[1].BorderTop)
	assert.Equal(t, true, row.Cells[2].BorderTop)
}

// SetBorderBot

func TestShouldSetBottomBorderOfAllCellsToTrue(t *testing.T) {
	row := NewRow(3)

	assert.Equal(t, false, row.Cells[0].BorderBottom)
	assert.Equal(t, false, row.Cells[1].BorderBottom)
	assert.Equal(t, false, row.Cells[2].BorderBottom)

	row.SetBorderBottom()

	assert.Equal(t, true, row.Cells[0].BorderBottom)
	assert.Equal(t, true, row.Cells[1].BorderBottom)
	assert.Equal(t, true, row.Cells[2].BorderBottom)
}
