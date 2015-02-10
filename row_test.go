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

// SetGroups

func TestShouldSetCellGroupsCorrectlyIfAllCellGroupsAreEqualToMinusOne(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{Group: -1},
			Cell{Group: -1},
			Cell{Group: -1},
			Cell{Group: -1},
		},
	}

	row.SetGroups(5)

	assert.Equal(t, 5, row.Cells[0].Group)
	assert.Equal(t, 6, row.Cells[1].Group)
	assert.Equal(t, 7, row.Cells[2].Group)
	assert.Equal(t, 8, row.Cells[3].Group)
}

func TestShouldSetCellGroupsCorrectlyOnlyForCellsWhichGroupIsEqualToMinusOne(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{Group: -1},
			Cell{Group: 2},
			Cell{Group: -1},
			Cell{Group: 1},
		},
	}

	row.SetGroups(10)

	assert.Equal(t, 10, row.Cells[0].Group)
	assert.Equal(t, 2, row.Cells[1].Group)
	assert.Equal(t, 11, row.Cells[2].Group)
	assert.Equal(t, 1, row.Cells[3].Group)
}

// getGroupIndexes

func TestShouldReturnCorrectGroupIndexesForOneCellGroup(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{Group: 1},
			Cell{Group: 2},
		},
	}

	start, end := row.getGroupIndexes(0)

	assert.Equal(t, 0, start)
	assert.Equal(t, 0, end)
}

func TestShouldReturnCorrectGroupIndexesForMultipleCellGroup(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{Group: 1},
			Cell{Group: 2},
			Cell{Group: 2},
			Cell{Group: 2},
			Cell{Group: 3},
		},
	}

	start, end := row.getGroupIndexes(1)

	assert.Equal(t, 1, start)
	assert.Equal(t, 3, end)
}

func TestShouldReturnCorrectGroupIndexesWhenCalledInChain(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{Group: 1},
			Cell{Group: 2},
			Cell{Group: 2},
			Cell{Group: 2},
			Cell{Group: 3},
			Cell{Group: 4},
			Cell{Group: 4},
		},
	}

	start, end := row.getGroupIndexes(0)

	assert.Equal(t, 0, start)
	assert.Equal(t, 0, end)

	start, end = row.getGroupIndexes(end + 1)

	assert.Equal(t, 1, start)
	assert.Equal(t, 3, end)

	start, end = row.getGroupIndexes(end + 1)

	assert.Equal(t, 4, start)
	assert.Equal(t, 4, end)

	start, end = row.getGroupIndexes(end + 1)

	assert.Equal(t, 5, start)
	assert.Equal(t, 6, end)
}

// areCellBottomsLocked

func TestShouldReturnFalseIfAtLeastOneCellBottomBorrderIsEqualToFalse(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: false},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: false},
			Cell{BorderBottom: true},
			Cell{BorderBottom: false},
		},
	}

	assert.Equal(t, false, row.areCellBottomsLocked(1, 4))
}

func TestShouldReturnTrueIfAllBottomBorrdersAreEqualToTrue(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: false},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: false},
		},
	}

	assert.Equal(t, true, row.areCellBottomsLocked(1, 4))
}

func TestShouldReturnTrueIfCellBottomBorderIsEqualToTrue(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: false},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: false},
		},
	}

	assert.Equal(t, true, row.areCellBottomsLocked(1, 1))
}

func TestShouldReturnFalseIfCellBottomBorderIsEqualToFalse(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: false},
			Cell{BorderBottom: false},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: false},
		},
	}

	assert.Equal(t, false, row.areCellBottomsLocked(1, 1))
}

// breakRandomCellBottom

func TestShouldBreakBottomBorderOfCell(t *testing.T) {
	generatorMock := new(randomGeneratorMock)
	generatorMock.On("random", 1, 3).Return(2)
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
			Cell{BorderBottom: true},
		},
		generator: generatorMock,
	}

	row.breakRandomCellBottom(1, 3)

	assert.Equal(t, true, row.Cells[0].BorderBottom)
	assert.Equal(t, true, row.Cells[1].BorderBottom)
	assert.Equal(t, false, row.Cells[2].BorderBottom)
	assert.Equal(t, true, row.Cells[3].BorderBottom)
	assert.Equal(t, true, row.Cells[4].BorderBottom)
}

func TestShouldBreakBottomBorderOfSingleCell(t *testing.T) {
	generatorMock := new(randomGeneratorMock)
	generatorMock.On("random", 0, 0).Return(0)
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: true},
		},
		generator: generatorMock,
	}

	row.breakRandomCellBottom(0, 0)

	assert.Equal(t, false, row.Cells[0].BorderBottom)

}

// CreateRightBorders

func TestShouldCreateRightBorderAndMergeCellGroups(t *testing.T) {
	generatorMock := new(randomGeneratorMock)
	generatorMock.On("random50").Return(true)
	row := Row{
		Cells: []Cell{
			Cell{BorderRight: false, Group: 1},
			Cell{BorderRight: false, Group: 2},
			Cell{BorderRight: false, Group: 3},
			Cell{BorderRight: false, Group: 4},
			Cell{BorderRight: false, Group: 5},
		},
		generator: generatorMock,
	}

	row.CreateRightBorders()

	assert.Equal(t, true, row.Cells[0].BorderRight)
	assert.Equal(t, 1, row.Cells[0].Group)

	assert.Equal(t, true, row.Cells[1].BorderRight)
	assert.Equal(t, 2, row.Cells[1].Group)

	assert.Equal(t, true, row.Cells[2].BorderRight)
	assert.Equal(t, 3, row.Cells[2].Group)

	assert.Equal(t, true, row.Cells[3].BorderRight)
	assert.Equal(t, 4, row.Cells[3].Group)

	assert.Equal(t, false, row.Cells[4].BorderRight)
	assert.Equal(t, 5, row.Cells[4].Group)
}

// Copy

func TestShouldReturnCopyOfRow(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderLeft: true, Group: 1},
			Cell{BorderBottom: true, Group: 2},
		},
		generator: defaultGenerator,
	}

	copy := row.Copy()

	assert.Equal(t, Row{
		Cells: []Cell{
			Cell{BorderLeft: true, Group: 1},
			Cell{BorderBottom: true, Group: 2},
		},
		generator: defaultGenerator,
	}, copy)
}

func TestChangingCopyShouldNotChangeOriginal(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderLeft: true, Group: 1},
			Cell{BorderBottom: true, Group: 2},
		},
		generator: defaultGenerator,
	}

	copy := row.Copy()
	copy.Cells[1].Group = 3

	assert.Equal(t, 2, row.Cells[1].Group)
}

// RemoveRightBorders

func TestShoulRemoveAllButLastRightBorders(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderRight: true},
			Cell{BorderRight: true},
			Cell{BorderRight: false},
			Cell{BorderRight: true},
		},
	}

	row.RemoveRightBorders()

	assert.Equal(t, false, row.Cells[0].BorderRight)
	assert.Equal(t, false, row.Cells[1].BorderRight)
	assert.Equal(t, false, row.Cells[2].BorderRight)
	assert.Equal(t, true, row.Cells[3].BorderRight)
}

// RemoveGroupsAndBottomBordersIfNeed

func TestShoulRemoveCellGroupsAndBottomBorders(t *testing.T) {
	row := Row{
		Cells: []Cell{
			Cell{BorderBottom: true, Group: 1},
			Cell{BorderBottom: false, Group: 2},
			Cell{BorderBottom: true, Group: 3},
			Cell{BorderBottom: false, Group: 4},
		},
	}

	row.RemoveGroupsAndBottomBordersIfNeed()

	assert.Equal(t, Cell{BorderBottom: false, Group: -1}, row.Cells[0])
	assert.Equal(t, Cell{BorderBottom: false, Group: 2}, row.Cells[1])
	assert.Equal(t, Cell{BorderBottom: false, Group: -1}, row.Cells[2])
	assert.Equal(t, Cell{BorderBottom: false, Group: 4}, row.Cells[3])
}
