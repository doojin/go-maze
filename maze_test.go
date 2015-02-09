package maze

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// New

func TestIfWidthIsSmallerThanOneShouldReturnError(t *testing.T) {
	_, err := New(0, 4)

	assert.Equal(t, errors.New("Width cannot be smaller than 1"), err)
}

func TestIfHeightIsSmallerThanOneShouldReturnError(t *testing.T) {
	_, err := New(4, 0)

	assert.Equal(t, errors.New("Height cannot be smaller than 1"), err)
}
