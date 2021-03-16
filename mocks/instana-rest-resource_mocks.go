// Code generated by MockGen. DO NOT EDIT.
// Source: instana/restapi/instana-rest-resource.go

// Package mocks is a generated GoMock package.
package mocks

import (
	restapi "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInstanaDataObject is a mock of InstanaDataObject interface
type MockInstanaDataObject struct {
	ctrl     *gomock.Controller
	recorder *MockInstanaDataObjectMockRecorder
}

// MockInstanaDataObjectMockRecorder is the mock recorder for MockInstanaDataObject
type MockInstanaDataObjectMockRecorder struct {
	mock *MockInstanaDataObject
}

// NewMockInstanaDataObject creates a new mock instance
func NewMockInstanaDataObject(ctrl *gomock.Controller) *MockInstanaDataObject {
	mock := &MockInstanaDataObject{ctrl: ctrl}
	mock.recorder = &MockInstanaDataObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstanaDataObject) EXPECT() *MockInstanaDataObjectMockRecorder {
	return m.recorder
}

// GetIDForResourcePath mocks base method
func (m *MockInstanaDataObject) GetIDForResourcePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDForResourcePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIDForResourcePath indicates an expected call of GetIDForResourcePath
func (mr *MockInstanaDataObjectMockRecorder) GetIDForResourcePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDForResourcePath", reflect.TypeOf((*MockInstanaDataObject)(nil).GetIDForResourcePath))
}

// Validate mocks base method
func (m *MockInstanaDataObject) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockInstanaDataObjectMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockInstanaDataObject)(nil).Validate))
}

// MockRestResource is a mock of RestResource interface
type MockRestResource struct {
	ctrl     *gomock.Controller
	recorder *MockRestResourceMockRecorder
}

// MockRestResourceMockRecorder is the mock recorder for MockRestResource
type MockRestResourceMockRecorder struct {
	mock *MockRestResource
}

// NewMockRestResource creates a new mock instance
func NewMockRestResource(ctrl *gomock.Controller) *MockRestResource {
	mock := &MockRestResource{ctrl: ctrl}
	mock.recorder = &MockRestResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRestResource) EXPECT() *MockRestResourceMockRecorder {
	return m.recorder
}

// GetOne mocks base method
func (m *MockRestResource) GetOne(id string) (restapi.InstanaDataObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(restapi.InstanaDataObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockRestResourceMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockRestResource)(nil).GetOne), id)
}

// Create mocks base method
func (m *MockRestResource) Create(data restapi.InstanaDataObject) (restapi.InstanaDataObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", data)
	ret0, _ := ret[0].(restapi.InstanaDataObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRestResourceMockRecorder) Create(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRestResource)(nil).Create), data)
}

// Update mocks base method
func (m *MockRestResource) Update(data restapi.InstanaDataObject) (restapi.InstanaDataObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", data)
	ret0, _ := ret[0].(restapi.InstanaDataObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockRestResourceMockRecorder) Update(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRestResource)(nil).Update), data)
}

// Delete mocks base method
func (m *MockRestResource) Delete(data restapi.InstanaDataObject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRestResourceMockRecorder) Delete(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRestResource)(nil).Delete), data)
}

// DeleteByID mocks base method
func (m *MockRestResource) DeleteByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockRestResourceMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockRestResource)(nil).DeleteByID), id)
}

// MockReadOnlyRestResource is a mock of ReadOnlyRestResource interface
type MockReadOnlyRestResource struct {
	ctrl     *gomock.Controller
	recorder *MockReadOnlyRestResourceMockRecorder
}

// MockReadOnlyRestResourceMockRecorder is the mock recorder for MockReadOnlyRestResource
type MockReadOnlyRestResourceMockRecorder struct {
	mock *MockReadOnlyRestResource
}

// NewMockReadOnlyRestResource creates a new mock instance
func NewMockReadOnlyRestResource(ctrl *gomock.Controller) *MockReadOnlyRestResource {
	mock := &MockReadOnlyRestResource{ctrl: ctrl}
	mock.recorder = &MockReadOnlyRestResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadOnlyRestResource) EXPECT() *MockReadOnlyRestResourceMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockReadOnlyRestResource) GetAll() (*[]restapi.InstanaDataObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]restapi.InstanaDataObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockReadOnlyRestResourceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockReadOnlyRestResource)(nil).GetAll))
}

// GetOne mocks base method
func (m *MockReadOnlyRestResource) GetOne(id string) (restapi.InstanaDataObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(restapi.InstanaDataObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockReadOnlyRestResourceMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockReadOnlyRestResource)(nil).GetOne), id)
}

// MockJSONUnmarshaller is a mock of JSONUnmarshaller interface
type MockJSONUnmarshaller struct {
	ctrl     *gomock.Controller
	recorder *MockJSONUnmarshallerMockRecorder
}

// MockJSONUnmarshallerMockRecorder is the mock recorder for MockJSONUnmarshaller
type MockJSONUnmarshallerMockRecorder struct {
	mock *MockJSONUnmarshaller
}

// NewMockJSONUnmarshaller creates a new mock instance
func NewMockJSONUnmarshaller(ctrl *gomock.Controller) *MockJSONUnmarshaller {
	mock := &MockJSONUnmarshaller{ctrl: ctrl}
	mock.recorder = &MockJSONUnmarshallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJSONUnmarshaller) EXPECT() *MockJSONUnmarshallerMockRecorder {
	return m.recorder
}

// Unmarshal mocks base method
func (m *MockJSONUnmarshaller) Unmarshal(data []byte) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockJSONUnmarshallerMockRecorder) Unmarshal(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockJSONUnmarshaller)(nil).Unmarshal), data)
}
