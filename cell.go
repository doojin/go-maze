package maze

// Cell is a block maze consists of
type Cell struct {
	BorderLeft   bool
	BorderTop    bool
	BorderRight  bool
	BorderBottom bool
	Group        int
}
