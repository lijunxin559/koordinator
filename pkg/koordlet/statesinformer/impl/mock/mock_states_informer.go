// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/koordlet/statesinformer/api.go

// Package impl is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	v1alpha10 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	statesinformer "github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	v1 "k8s.io/api/core/v1"
)

// MockStatesInformer is a mock of StatesInformer interface.
type MockStatesInformer struct {
	ctrl     *gomock.Controller
	recorder *MockStatesInformerMockRecorder
}

// MockStatesInformerMockRecorder is the mock recorder for MockStatesInformer.
type MockStatesInformerMockRecorder struct {
	mock *MockStatesInformer
}

// NewMockStatesInformer creates a new mock instance.
func NewMockStatesInformer(ctrl *gomock.Controller) *MockStatesInformer {
	mock := &MockStatesInformer{ctrl: ctrl}
	mock.recorder = &MockStatesInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatesInformer) EXPECT() *MockStatesInformerMockRecorder {
	return m.recorder
}

// GetAllPods mocks base method.
func (m *MockStatesInformer) GetAllPods() []*statesinformer.PodMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPods")
	ret0, _ := ret[0].([]*statesinformer.PodMeta)
	return ret0
}

// GetAllPods indicates an expected call of GetAllPods.
func (mr *MockStatesInformerMockRecorder) GetAllPods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPods", reflect.TypeOf((*MockStatesInformer)(nil).GetAllPods))
}

// GetNode mocks base method.
func (m *MockStatesInformer) GetNode() *v1.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode")
	ret0, _ := ret[0].(*v1.Node)
	return ret0
}

// GetNode indicates an expected call of GetNode.
func (mr *MockStatesInformerMockRecorder) GetNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockStatesInformer)(nil).GetNode))
}

// GetNodeMetricSpec mocks base method.
func (m *MockStatesInformer) GetNodeMetricSpec() *v1alpha10.NodeMetricSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeMetricSpec")
	ret0, _ := ret[0].(*v1alpha10.NodeMetricSpec)
	return ret0
}

// GetNodeMetricSpec indicates an expected call of GetNodeMetricSpec.
func (mr *MockStatesInformerMockRecorder) GetNodeMetricSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeMetricSpec", reflect.TypeOf((*MockStatesInformer)(nil).GetNodeMetricSpec))
}

// GetNodeSLO mocks base method.
func (m *MockStatesInformer) GetNodeSLO() *v1alpha10.NodeSLO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeSLO")
	ret0, _ := ret[0].(*v1alpha10.NodeSLO)
	return ret0
}

// GetNodeSLO indicates an expected call of GetNodeSLO.
func (mr *MockStatesInformerMockRecorder) GetNodeSLO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeSLO", reflect.TypeOf((*MockStatesInformer)(nil).GetNodeSLO))
}

// GetNodeTopo mocks base method.
func (m *MockStatesInformer) GetNodeTopo() *v1alpha1.NodeResourceTopology {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeTopo")
	ret0, _ := ret[0].(*v1alpha1.NodeResourceTopology)
	return ret0
}

// GetNodeTopo indicates an expected call of GetNodeTopo.
func (mr *MockStatesInformerMockRecorder) GetNodeTopo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeTopo", reflect.TypeOf((*MockStatesInformer)(nil).GetNodeTopo))
}

// GetVolumeName mocks base method.
func (m *MockStatesInformer) GetVolumeName(pvcNamespace, pvcName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeName", pvcNamespace, pvcName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVolumeName indicates an expected call of GetVolumeName.
func (mr *MockStatesInformerMockRecorder) GetVolumeName(pvcNamespace, pvcName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeName", reflect.TypeOf((*MockStatesInformer)(nil).GetVolumeName), pvcNamespace, pvcName)
}

// HasSynced mocks base method.
func (m *MockStatesInformer) HasSynced() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSynced")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSynced indicates an expected call of HasSynced.
func (mr *MockStatesInformerMockRecorder) HasSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSynced", reflect.TypeOf((*MockStatesInformer)(nil).HasSynced))
}

// RegisterCallbacks mocks base method.
func (m *MockStatesInformer) RegisterCallbacks(objType statesinformer.RegisterType, name, description string, callbackFn statesinformer.UpdateCbFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterCallbacks", objType, name, description, callbackFn)
}

// RegisterCallbacks indicates an expected call of RegisterCallbacks.
func (mr *MockStatesInformerMockRecorder) RegisterCallbacks(objType, name, description, callbackFn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCallbacks", reflect.TypeOf((*MockStatesInformer)(nil).RegisterCallbacks), objType, name, description, callbackFn)
}

// Run mocks base method.
func (m *MockStatesInformer) Run(stopCh <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", stopCh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockStatesInformerMockRecorder) Run(stopCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockStatesInformer)(nil).Run), stopCh)
}
