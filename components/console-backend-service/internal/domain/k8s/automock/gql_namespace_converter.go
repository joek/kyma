// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// gqlNamespaceConverter is an autogenerated mock type for the gqlNamespaceConverter type
type gqlNamespaceConverter struct {
	mock.Mock
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlNamespaceConverter) ToGQLs(in []*v1.Namespace) ([]gqlschema.Namespace, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.Namespace
	if rf, ok := ret.Get(0).(func([]*v1.Namespace) []gqlschema.Namespace); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1.Namespace) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
