package maze

// Row is a set of cells
type Row struct {
	Cells     []Cell
	generator randomGeneratorI
}

// NewRow returns a row
func NewRow(size int) Row {
	row := Row{generator: defaultGenerator}
	for i := 0; i < size; i++ {
		row.Cells = append(row.Cells, NewCell())
	}
	row.Cells[0].BorderLeft = true
	row.Cells[size-1].BorderRight = true
	return row
}

// SetBorderTop sets top border of all cells to true
func (row *Row) SetBorderTop() {
	for i := range row.Cells {
		row.Cells[i].BorderTop = true
	}
}

// SetBorderBottom sets bottom border of all cells to true
func (row *Row) SetBorderBottom() {
	for i := range row.Cells {
		row.Cells[i].BorderBottom = true
	}
}

// SetGroups updates group of cell if it is equal to -1
func (row *Row) SetGroups(startGroup int) {
	for i, cell := range row.Cells {
		if cell.Group == -1 {
			row.Cells[i].Group = startGroup
			startGroup++
		}
	}
}

// CreateRightBorders with probability of 50% will create right border between cells
func (row *Row) CreateRightBorders() {
	for i := 0; i < len(row.Cells)-1; i++ {
		if row.Cells[i].Group == row.Cells[i+1].Group {
			row.Cells[i].BorderRight = true
			return
		}
		// 50%
		if row.generator.random50() {
			row.Cells[i].BorderRight = true
		} else {
			row.Cells[i+1].Group = row.Cells[i].Group
		}
	}
}

// CreateBottomBorders with probability of 50% will create bottom border for a cell
func (row *Row) CreateBottomBorders() {
	for i := range row.Cells {
		if row.generator.random50() {
			row.Cells[i].BorderBottom = true
		}
	}
	row.destroyBottomBorderIfNeed()
}

// destroyBottomBorderIfNeed destroys one borrom border of cell group if all cells of group have a bottom border
func (row *Row) destroyBottomBorderIfNeed() {
	end := -1
	for end < len(row.Cells)-1 {
		end++
		start, end := row.getGroupIndexes(end)
		if row.areCellBottomsLocked(start, end) {
			row.breakRandomCellBottom(start, end)
		}
	}
}

// getGroup returns start end end indexes of cell group starting with index argument
func (row *Row) getGroupIndexes(index int) (int, int) {
	currentGroup := row.Cells[index].Group
	var endIndex int
	for endIndex = index + 1; endIndex < len(row.Cells); endIndex++ {
		if row.Cells[endIndex].Group != currentGroup {
			break
		}
	}
	endIndex--
	return index, endIndex
}

// areCellBottomsLocked returns true if all cell (from start to end
func (row *Row) areCellBottomsLocked(start int, end int) bool {
	for i := start; i <= end; i++ {
		if !row.Cells[i].BorderBottom {
			return false
		}
	}
	return true
}

// breakRandomCellBottom unsets borderBottom to a random cell on interval from start to end
func (row *Row) breakRandomCellBottom(start int, end int) {
	index := row.generator.random(start, end)
	row.Cells[index].BorderBottom = false
}

// Copy creates a copy of row
func (row Row) Copy() Row {
	copy := Row{}
	copy.Cells = []Cell{}
	for _, cell := range row.Cells {
		copy.Cells = append(copy.Cells, cell.Copy())
	}
	copy.generator = defaultGenerator
	return copy
}

// RemoveRightBorders removes all right borders of cells
func (row *Row) RemoveRightBorders() {
	for i := range row.Cells {
		if i != len(row.Cells)-1 {
			row.Cells[i].BorderRight = false
		}
	}
}

// RemoveGroupsAndBottomBordersIfNeed removes groups from cells where bottom border is set
func (row *Row) RemoveGroupsAndBottomBordersIfNeed() {
	for i, cell := range row.Cells {
		if cell.BorderBottom {
			row.Cells[i].Group = -1
			row.Cells[i].BorderBottom = false
		}
	}
}
