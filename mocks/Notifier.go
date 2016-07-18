package mocks

import "github.com/stretchr/testify/mock"

type Notifier struct {
	mock.Mock
}

func (_m *Notifier) Notify() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
