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
	cell := Cell{false, false, false, false, 0}
	return cell
}
