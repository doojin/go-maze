package maze

// Row is a set of cells
type Row struct {
	Cells []Cell
}

// NewRow returns a row
func NewRow(size int) Row {
	row := Row{}
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
