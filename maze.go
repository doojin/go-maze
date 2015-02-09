package maze

// Maze consists of rows
type Maze struct {
	Rows []Row
}

// New returns an instance of Maze
func New(width int, height int) Maze {
	maze := Maze{}
	for i := 0; i < height; i++ {
		maze.Rows = append(maze.Rows, NewRow(width))
	}
	return maze
}
