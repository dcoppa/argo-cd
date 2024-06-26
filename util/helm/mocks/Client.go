// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	helm "github.com/dcoppa/argo-cd/v2/util/helm"
	io "github.com/dcoppa/argo-cd/v2/util/io"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CleanChartCache provides a mock function with given fields: chart, version
func (_m *Client) CleanChartCache(chart string, version string) error {
	ret := _m.Called(chart, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(chart, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtractChart provides a mock function with given fields: chart, version
func (_m *Client) ExtractChart(chart string, version string, passCredentials bool, manifestMaxExtractedSize int64, disableManifestMaxExtractedSize bool) (string, io.Closer, error) {
	ret := _m.Called(chart, version)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(chart, version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 io.Closer
	if rf, ok := ret.Get(1).(func(string, string) io.Closer); ok {
		r1 = rf(chart, version)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.Closer)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(chart, version)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetIndex provides a mock function with given fields: noCache
func (_m *Client) GetIndex(noCache bool, maxIndexSize int64) (*helm.Index, error) {
	ret := _m.Called(noCache)

	var r0 *helm.Index
	if rf, ok := ret.Get(0).(func(bool) *helm.Index); ok {
		r0 = rf(noCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*helm.Index)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(noCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: noCache
func (_m *Client) GetTags(chart string, noCache bool) (*helm.TagsList, error) {
	ret := _m.Called(chart, noCache)

	var r0 *helm.TagsList
	if rf, ok := ret.Get(0).(func(string, bool) *helm.TagsList); ok {
		r0 = rf(chart, noCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*helm.TagsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(chart, noCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestHelmOCI provides a mock function with given fields:
func (_m *Client) TestHelmOCI() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
