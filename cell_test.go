package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewCell

func TestShouldReturnCellWithDefaultValues(t *testing.T) {
	cell := NewCell()

	assert.Equal(t, cell.BorderLeft, false)
	assert.Equal(t, cell.BorderTop, false)
	assert.Equal(t, cell.BorderRight, false)
	assert.Equal(t, cell.BorderBottom, false)
	assert.Equal(t, cell.Group, -1)
}
