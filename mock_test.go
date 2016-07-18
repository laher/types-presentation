package types

import (
	"testing"

	"github.com/laher/types/mocks"
	"github.com/stretchr/testify/assert"
)

type Container struct {
	Notifier
}

func (c Container) Do() string {
	return c.Notify()
}

func TestMock(t *testing.T) {
	mockObj := &mocks.Notifier{}
	mockObj.On("Notify").Return("mock")
	c := Container{
		Notifier: mockObj,
	}
	result := c.Do()
	mockObj.AssertCalled(t, "Notify")
	assert.Equal(t, "mock", result)
}
