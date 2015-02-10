package maze

import "github.com/stretchr/testify/mock"

type randomGeneratorMock struct {
	mock.Mock
}

func (mock *randomGeneratorMock) random(min int, max int) int {
	args := mock.Called(min, max)
	return args.Int(0)
}

func (mock *randomGeneratorMock) random50() bool {
	args := mock.Called()
	return args.Bool(0)
}
