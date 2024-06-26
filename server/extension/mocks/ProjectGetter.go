// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/dcoppa/argo-cd/v2/pkg/apis/application/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// ProjectGetter is an autogenerated mock type for the ProjectGetter type
type ProjectGetter struct {
	mock.Mock
}

// Get provides a mock function with given fields: name
func (_m *ProjectGetter) Get(name string) (*v1alpha1.AppProject, error) {
	ret := _m.Called(name)

	var r0 *v1alpha1.AppProject
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.AppProject); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.AppProject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusters provides a mock function with given fields: project
func (_m *ProjectGetter) GetClusters(project string) ([]*v1alpha1.Cluster, error) {
	ret := _m.Called(project)

	var r0 []*v1alpha1.Cluster
	if rf, ok := ret.Get(0).(func(string) []*v1alpha1.Cluster); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProjectGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewProjectGetter creates a new instance of ProjectGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProjectGetter(t mockConstructorTestingTNewProjectGetter) *ProjectGetter {
	mock := &ProjectGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
