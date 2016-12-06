// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/quintilesims/layer0/common/aws/ecs (interfaces: Provider)

package mock_ecs

import (
	gomock "github.com/golang/mock/gomock"
	ecs "github.com/quintilesims/layer0/common/aws/ecs"
)

// Mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *_MockProviderRecorder
}

// Recorder for MockProvider (not exported)
type _MockProviderRecorder struct {
	mock *MockProvider
}

func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &_MockProviderRecorder{mock}
	return mock
}

func (_m *MockProvider) EXPECT() *_MockProviderRecorder {
	return _m.recorder
}

func (_m *MockProvider) CreateCluster(_param0 string) (*ecs.Cluster, error) {
	ret := _m.ctrl.Call(_m, "CreateCluster", _param0)
	ret0, _ := ret[0].(*ecs.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateCluster", arg0)
}

func (_m *MockProvider) CreateService(_param0 string, _param1 string, _param2 string, _param3 int64, _param4 []*ecs.LoadBalancer, _param5 *string) (*ecs.Service, error) {
	ret := _m.ctrl.Call(_m, "CreateService", _param0, _param1, _param2, _param3, _param4, _param5)
	ret0, _ := ret[0].(*ecs.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) CreateService(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateService", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockProvider) DeleteCluster(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteCluster", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) DeleteCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCluster", arg0)
}

func (_m *MockProvider) DeleteService(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteService", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) DeleteService(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteService", arg0, arg1)
}

func (_m *MockProvider) DeleteTaskDefinition(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTaskDefinition", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) DeleteTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTaskDefinition", arg0)
}

func (_m *MockProvider) DescribeCluster(_param0 string) (*ecs.Cluster, error) {
	ret := _m.ctrl.Call(_m, "DescribeCluster", _param0)
	ret0, _ := ret[0].(*ecs.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeCluster", arg0)
}

func (_m *MockProvider) DescribeContainerInstances(_param0 string, _param1 []*string) ([]*ecs.ContainerInstance, error) {
	ret := _m.ctrl.Call(_m, "DescribeContainerInstances", _param0, _param1)
	ret0, _ := ret[0].([]*ecs.ContainerInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeContainerInstances(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeContainerInstances", arg0, arg1)
}

func (_m *MockProvider) DescribeService(_param0 string, _param1 string) (*ecs.Service, error) {
	ret := _m.ctrl.Call(_m, "DescribeService", _param0, _param1)
	ret0, _ := ret[0].(*ecs.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeService(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeService", arg0, arg1)
}

func (_m *MockProvider) DescribeServices(_param0 string, _param1 []*string) ([]*ecs.Service, error) {
	ret := _m.ctrl.Call(_m, "DescribeServices", _param0, _param1)
	ret0, _ := ret[0].([]*ecs.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeServices(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeServices", arg0, arg1)
}

func (_m *MockProvider) DescribeTaskDefinition(_param0 string) (*ecs.TaskDefinition, error) {
	ret := _m.ctrl.Call(_m, "DescribeTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTaskDefinition", arg0)
}

func (_m *MockProvider) DescribeTasks(_param0 string, _param1 []*string) ([]*ecs.Task, error) {
	ret := _m.ctrl.Call(_m, "DescribeTasks", _param0, _param1)
	ret0, _ := ret[0].([]*ecs.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeTasks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTasks", arg0, arg1)
}

func (_m *MockProvider) Helper_DescribeClusters() ([]*ecs.Cluster, error) {
	ret := _m.ctrl.Call(_m, "Helper_DescribeClusters")
	ret0, _ := ret[0].([]*ecs.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) Helper_DescribeClusters() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Helper_DescribeClusters")
}

func (_m *MockProvider) Helper_DescribeServices(_param0 string) ([]*ecs.Service, error) {
	ret := _m.ctrl.Call(_m, "Helper_DescribeServices", _param0)
	ret0, _ := ret[0].([]*ecs.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) Helper_DescribeServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Helper_DescribeServices", arg0)
}

func (_m *MockProvider) Helper_DescribeTaskDefinitions(_param0 string) ([]*ecs.TaskDefinition, error) {
	ret := _m.ctrl.Call(_m, "Helper_DescribeTaskDefinitions", _param0)
	ret0, _ := ret[0].([]*ecs.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) Helper_DescribeTaskDefinitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Helper_DescribeTaskDefinitions", arg0)
}

func (_m *MockProvider) Helper_ListTaskDefinitions(_param0 string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "Helper_ListTaskDefinitions", _param0)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) Helper_ListTaskDefinitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Helper_ListTaskDefinitions", arg0)
}

func (_m *MockProvider) ListClusters() ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListClusters")
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) ListClusters() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListClusters")
}

func (_m *MockProvider) ListContainerInstances(_param0 string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListContainerInstances", _param0)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) ListContainerInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainerInstances", arg0)
}

func (_m *MockProvider) ListServices(_param0 string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListServices", _param0)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) ListServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListServices", arg0)
}

func (_m *MockProvider) ListTaskDefinitionFamilies(_param0 string, _param1 *string) ([]*string, *string, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionFamilies", _param0, _param1)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(*string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockProviderRecorder) ListTaskDefinitionFamilies(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionFamilies", arg0, arg1)
}

func (_m *MockProvider) ListTaskDefinitionFamiliesPages(_param0 string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionFamiliesPages", _param0)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) ListTaskDefinitionFamiliesPages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionFamiliesPages", arg0)
}

func (_m *MockProvider) ListTaskDefinitions(_param0 string, _param1 *string) ([]*string, *string, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitions", _param0, _param1)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(*string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockProviderRecorder) ListTaskDefinitions(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitions", arg0, arg1)
}

func (_m *MockProvider) ListTaskDefinitionsPages(_param0 string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionsPages", _param0)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) ListTaskDefinitionsPages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionsPages", arg0)
}

func (_m *MockProvider) ListTasks(_param0 string, _param1 *string, _param2 *string, _param3 *string, _param4 *string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListTasks", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) ListTasks(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockProvider) RegisterTaskDefinition(_param0 string, _param1 string, _param2 string, _param3 []*ecs.ContainerDefinition, _param4 []*ecs.Volume) (*ecs.TaskDefinition, error) {
	ret := _m.ctrl.Call(_m, "RegisterTaskDefinition", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(*ecs.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) RegisterTaskDefinition(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterTaskDefinition", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockProvider) RunTask(_param0 string, _param1 string, _param2 int64, _param3 *string, _param4 []*ecs.ContainerOverride) ([]*ecs.Task, []*ecs.FailedTask, error) {
	ret := _m.ctrl.Call(_m, "RunTask", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].([]*ecs.Task)
	ret1, _ := ret[1].([]*ecs.FailedTask)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockProviderRecorder) RunTask(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunTask", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockProvider) StartTask(_param0 string, _param1 string, _param2 *ecs.TaskOverride, _param3 []*string, _param4 *string) error {
	ret := _m.ctrl.Call(_m, "StartTask", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) StartTask(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTask", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockProvider) StopTask(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "StopTask", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) StopTask(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopTask", arg0, arg1, arg2)
}

func (_m *MockProvider) UpdateService(_param0 string, _param1 string, _param2 *string, _param3 *int64) error {
	ret := _m.ctrl.Call(_m, "UpdateService", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) UpdateService(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateService", arg0, arg1, arg2, arg3)
}
