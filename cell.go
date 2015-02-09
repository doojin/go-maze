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
func NewCell(group int) Cell {
	cell := Cell{false, false, false, false, group}
	return cell
}
