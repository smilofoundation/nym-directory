// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/nymtech/nym-directory/models"

// IDb is an autogenerated mock type for the IDb type
type IDb struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *IDb) Add(_a0 models.PersistedMixMetric) {
	_m.Called(_a0)
}

// List provides a mock function with given fields:
func (_m *IDb) List() []models.PersistedMixMetric {
	ret := _m.Called()

	var r0 []models.PersistedMixMetric
	if rf, ok := ret.Get(0).(func() []models.PersistedMixMetric); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PersistedMixMetric)
		}
	}

	return r0
}