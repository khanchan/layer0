// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/quintilesims/layer0/api/logic (interfaces: ServiceLogic)

package mock_logic

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/quintilesims/layer0/common/models"
)

// Mock of ServiceLogic interface
type MockServiceLogic struct {
	ctrl     *gomock.Controller
	recorder *_MockServiceLogicRecorder
}

// Recorder for MockServiceLogic (not exported)
type _MockServiceLogicRecorder struct {
	mock *MockServiceLogic
}

func NewMockServiceLogic(ctrl *gomock.Controller) *MockServiceLogic {
	mock := &MockServiceLogic{ctrl: ctrl}
	mock.recorder = &_MockServiceLogicRecorder{mock}
	return mock
}

func (_m *MockServiceLogic) EXPECT() *_MockServiceLogicRecorder {
	return _m.recorder
}

func (_m *MockServiceLogic) CreateService(_param0 models.CreateServiceRequest) (*models.Service, error) {
	ret := _m.ctrl.Call(_m, "CreateService", _param0)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceLogicRecorder) CreateService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateService", arg0)
}

func (_m *MockServiceLogic) DeleteService(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteService", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceLogicRecorder) DeleteService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteService", arg0)
}

func (_m *MockServiceLogic) GetService(_param0 string) (*models.Service, error) {
	ret := _m.ctrl.Call(_m, "GetService", _param0)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceLogicRecorder) GetService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetService", arg0)
}

func (_m *MockServiceLogic) GetServiceLogs(_param0 string, _param1 int) ([]*models.LogFile, error) {
	ret := _m.ctrl.Call(_m, "GetServiceLogs", _param0, _param1)
	ret0, _ := ret[0].([]*models.LogFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceLogicRecorder) GetServiceLogs(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServiceLogs", arg0, arg1)
}

func (_m *MockServiceLogic) ListServices() ([]*models.ServiceSummary, error) {
	ret := _m.ctrl.Call(_m, "ListServices")
	ret0, _ := ret[0].([]*models.ServiceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceLogicRecorder) ListServices() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListServices")
}

func (_m *MockServiceLogic) ScaleService(_param0 string, _param1 int) (*models.Service, error) {
	ret := _m.ctrl.Call(_m, "ScaleService", _param0, _param1)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceLogicRecorder) ScaleService(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScaleService", arg0, arg1)
}

func (_m *MockServiceLogic) UpdateService(_param0 string, _param1 models.UpdateServiceRequest) (*models.Service, error) {
	ret := _m.ctrl.Call(_m, "UpdateService", _param0, _param1)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceLogicRecorder) UpdateService(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateService", arg0, arg1)
}
