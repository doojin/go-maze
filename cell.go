package maze

// Cell is a block maze consists of
type Cell struct {
	BorderLeft   bool
	BorderTop    bool
	BorderRight  bool
	BorderBottom bool
	Group        int
}

// NewCell returns new cell
func NewCell() Cell {
	cell := Cell{false, false, false, false, -1}
	return cell
}

// Copy returns copy of cell
func (cell Cell) Copy() Cell {
	copy := Cell{}
	copy.BorderLeft = cell.BorderLeft
	copy.BorderTop = cell.BorderTop
	copy.BorderRight = cell.BorderRight
	copy.BorderBottom = cell.BorderBottom
	copy.Group = cell.Group
	return copy
}
