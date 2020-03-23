// Code generated by MockGen. DO NOT EDIT.
// Source: api.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUpClient is a mock of UpClient interface
type MockUpClient struct {
	ctrl     *gomock.Controller
	recorder *MockUpClientMockRecorder
}

// MockUpClientMockRecorder is the mock recorder for MockUpClient
type MockUpClientMockRecorder struct {
	mock *MockUpClient
}

// NewMockUpClient creates a new mock instance
func NewMockUpClient(ctrl *gomock.Controller) *MockUpClient {
	mock := &MockUpClient{ctrl: ctrl}
	mock.recorder = &MockUpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpClient) EXPECT() *MockUpClientMockRecorder {
	return m.recorder
}

// UpArcs mocks base method
func (m *MockUpClient) UpArcs(ctx context.Context, in *UpArcsReq, opts ...grpc.CallOption) (*UpArcsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpArcs", varargs...)
	ret0, _ := ret[0].(*UpArcsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpArcs indicates an expected call of UpArcs
func (mr *MockUpClientMockRecorder) UpArcs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpArcs", reflect.TypeOf((*MockUpClient)(nil).UpArcs), varargs...)
}

// UpsArcs mocks base method
func (m *MockUpClient) UpsArcs(ctx context.Context, in *UpsArcsReq, opts ...grpc.CallOption) (*UpsArcsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsArcs", varargs...)
	ret0, _ := ret[0].(*UpsArcsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsArcs indicates an expected call of UpsArcs
func (mr *MockUpClientMockRecorder) UpsArcs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsArcs", reflect.TypeOf((*MockUpClient)(nil).UpsArcs), varargs...)
}

// UpCount mocks base method
func (m *MockUpClient) UpCount(ctx context.Context, in *UpCountReq, opts ...grpc.CallOption) (*UpCountReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpCount", varargs...)
	ret0, _ := ret[0].(*UpCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpCount indicates an expected call of UpCount
func (mr *MockUpClientMockRecorder) UpCount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpCount", reflect.TypeOf((*MockUpClient)(nil).UpCount), varargs...)
}

// UpsCount mocks base method
func (m *MockUpClient) UpsCount(ctx context.Context, in *UpsCountReq, opts ...grpc.CallOption) (*UpsCountReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsCount", varargs...)
	ret0, _ := ret[0].(*UpsCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsCount indicates an expected call of UpsCount
func (mr *MockUpClientMockRecorder) UpsCount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsCount", reflect.TypeOf((*MockUpClient)(nil).UpsCount), varargs...)
}

// UpsAidPubTime mocks base method
func (m *MockUpClient) UpsAidPubTime(ctx context.Context, in *UpsArcsReq, opts ...grpc.CallOption) (*UpsAidPubTimeReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsAidPubTime", varargs...)
	ret0, _ := ret[0].(*UpsAidPubTimeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsAidPubTime indicates an expected call of UpsAidPubTime
func (mr *MockUpClientMockRecorder) UpsAidPubTime(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsAidPubTime", reflect.TypeOf((*MockUpClient)(nil).UpsAidPubTime), varargs...)
}

// AddUpPassedCacheByStaff mocks base method
func (m *MockUpClient) AddUpPassedCacheByStaff(ctx context.Context, in *UpCacheReq, opts ...grpc.CallOption) (*NoReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUpPassedCacheByStaff", varargs...)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUpPassedCacheByStaff indicates an expected call of AddUpPassedCacheByStaff
func (mr *MockUpClientMockRecorder) AddUpPassedCacheByStaff(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpPassedCacheByStaff", reflect.TypeOf((*MockUpClient)(nil).AddUpPassedCacheByStaff), varargs...)
}

// AddUpPassedCache mocks base method
func (m *MockUpClient) AddUpPassedCache(ctx context.Context, in *UpCacheReq, opts ...grpc.CallOption) (*NoReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUpPassedCache", varargs...)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUpPassedCache indicates an expected call of AddUpPassedCache
func (mr *MockUpClientMockRecorder) AddUpPassedCache(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpPassedCache", reflect.TypeOf((*MockUpClient)(nil).AddUpPassedCache), varargs...)
}

// DelUpPassedCacheByStaff mocks base method
func (m *MockUpClient) DelUpPassedCacheByStaff(ctx context.Context, in *UpCacheReq, opts ...grpc.CallOption) (*NoReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelUpPassedCacheByStaff", varargs...)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelUpPassedCacheByStaff indicates an expected call of DelUpPassedCacheByStaff
func (mr *MockUpClientMockRecorder) DelUpPassedCacheByStaff(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelUpPassedCacheByStaff", reflect.TypeOf((*MockUpClient)(nil).DelUpPassedCacheByStaff), varargs...)
}

// DelUpPassedCache mocks base method
func (m *MockUpClient) DelUpPassedCache(ctx context.Context, in *UpCacheReq, opts ...grpc.CallOption) (*NoReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelUpPassedCache", varargs...)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelUpPassedCache indicates an expected call of DelUpPassedCache
func (mr *MockUpClientMockRecorder) DelUpPassedCache(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelUpPassedCache", reflect.TypeOf((*MockUpClient)(nil).DelUpPassedCache), varargs...)
}

// UpInfoActivitys mocks base method
func (m *MockUpClient) UpInfoActivitys(ctx context.Context, in *UpListByLastIDReq, opts ...grpc.CallOption) (*UpActivityListReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpInfoActivitys", varargs...)
	ret0, _ := ret[0].(*UpActivityListReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpInfoActivitys indicates an expected call of UpInfoActivitys
func (mr *MockUpClientMockRecorder) UpInfoActivitys(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpInfoActivitys", reflect.TypeOf((*MockUpClient)(nil).UpInfoActivitys), varargs...)
}

// UpSpecial mocks base method
func (m *MockUpClient) UpSpecial(ctx context.Context, in *UpSpecialReq, opts ...grpc.CallOption) (*UpSpecialReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpSpecial", varargs...)
	ret0, _ := ret[0].(*UpSpecialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpSpecial indicates an expected call of UpSpecial
func (mr *MockUpClientMockRecorder) UpSpecial(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpSpecial", reflect.TypeOf((*MockUpClient)(nil).UpSpecial), varargs...)
}

// UpsSpecial mocks base method
func (m *MockUpClient) UpsSpecial(ctx context.Context, in *UpsSpecialReq, opts ...grpc.CallOption) (*UpsSpecialReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsSpecial", varargs...)
	ret0, _ := ret[0].(*UpsSpecialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsSpecial indicates an expected call of UpsSpecial
func (mr *MockUpClientMockRecorder) UpsSpecial(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsSpecial", reflect.TypeOf((*MockUpClient)(nil).UpsSpecial), varargs...)
}

// UpGroups mocks base method
func (m *MockUpClient) UpGroups(ctx context.Context, in *NoArgReq, opts ...grpc.CallOption) (*UpGroupsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpGroups", varargs...)
	ret0, _ := ret[0].(*UpGroupsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpGroups indicates an expected call of UpGroups
func (mr *MockUpClientMockRecorder) UpGroups(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpGroups", reflect.TypeOf((*MockUpClient)(nil).UpGroups), varargs...)
}

// UpGroupMids mocks base method
func (m *MockUpClient) UpGroupMids(ctx context.Context, in *UpGroupMidsReq, opts ...grpc.CallOption) (*UpGroupMidsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpGroupMids", varargs...)
	ret0, _ := ret[0].(*UpGroupMidsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpGroupMids indicates an expected call of UpGroupMids
func (mr *MockUpClientMockRecorder) UpGroupMids(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpGroupMids", reflect.TypeOf((*MockUpClient)(nil).UpGroupMids), varargs...)
}

// UpAttr mocks base method
func (m *MockUpClient) UpAttr(ctx context.Context, in *UpAttrReq, opts ...grpc.CallOption) (*UpAttrReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpAttr", varargs...)
	ret0, _ := ret[0].(*UpAttrReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpAttr indicates an expected call of UpAttr
func (mr *MockUpClientMockRecorder) UpAttr(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpAttr", reflect.TypeOf((*MockUpClient)(nil).UpAttr), varargs...)
}

// UpBaseStats mocks base method
func (m *MockUpClient) UpBaseStats(ctx context.Context, in *UpStatReq, opts ...grpc.CallOption) (*UpBaseStatReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpBaseStats", varargs...)
	ret0, _ := ret[0].(*UpBaseStatReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpBaseStats indicates an expected call of UpBaseStats
func (mr *MockUpClientMockRecorder) UpBaseStats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpBaseStats", reflect.TypeOf((*MockUpClient)(nil).UpBaseStats), varargs...)
}

// SetUpSwitch mocks base method
func (m *MockUpClient) SetUpSwitch(ctx context.Context, in *UpSwitchReq, opts ...grpc.CallOption) (*NoReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetUpSwitch", varargs...)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUpSwitch indicates an expected call of SetUpSwitch
func (mr *MockUpClientMockRecorder) SetUpSwitch(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpSwitch", reflect.TypeOf((*MockUpClient)(nil).SetUpSwitch), varargs...)
}

// UpSwitch mocks base method
func (m *MockUpClient) UpSwitch(ctx context.Context, in *UpSwitchReq, opts ...grpc.CallOption) (*UpSwitchReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpSwitch", varargs...)
	ret0, _ := ret[0].(*UpSwitchReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpSwitch indicates an expected call of UpSwitch
func (mr *MockUpClientMockRecorder) UpSwitch(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpSwitch", reflect.TypeOf((*MockUpClient)(nil).UpSwitch), varargs...)
}

// GetHighAllyUps mocks base method
func (m *MockUpClient) GetHighAllyUps(ctx context.Context, in *HighAllyUpsReq, opts ...grpc.CallOption) (*HighAllyUpsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHighAllyUps", varargs...)
	ret0, _ := ret[0].(*HighAllyUpsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighAllyUps indicates an expected call of GetHighAllyUps
func (mr *MockUpClientMockRecorder) GetHighAllyUps(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighAllyUps", reflect.TypeOf((*MockUpClient)(nil).GetHighAllyUps), varargs...)
}

// MockUpServer is a mock of UpServer interface
type MockUpServer struct {
	ctrl     *gomock.Controller
	recorder *MockUpServerMockRecorder
}

// MockUpServerMockRecorder is the mock recorder for MockUpServer
type MockUpServerMockRecorder struct {
	mock *MockUpServer
}

// NewMockUpServer creates a new mock instance
func NewMockUpServer(ctrl *gomock.Controller) *MockUpServer {
	mock := &MockUpServer{ctrl: ctrl}
	mock.recorder = &MockUpServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpServer) EXPECT() *MockUpServerMockRecorder {
	return m.recorder
}

// UpArcs mocks base method
func (m *MockUpServer) UpArcs(arg0 context.Context, arg1 *UpArcsReq) (*UpArcsReply, error) {
	ret := m.ctrl.Call(m, "UpArcs", arg0, arg1)
	ret0, _ := ret[0].(*UpArcsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpArcs indicates an expected call of UpArcs
func (mr *MockUpServerMockRecorder) UpArcs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpArcs", reflect.TypeOf((*MockUpServer)(nil).UpArcs), arg0, arg1)
}

// UpsArcs mocks base method
func (m *MockUpServer) UpsArcs(arg0 context.Context, arg1 *UpsArcsReq) (*UpsArcsReply, error) {
	ret := m.ctrl.Call(m, "UpsArcs", arg0, arg1)
	ret0, _ := ret[0].(*UpsArcsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsArcs indicates an expected call of UpsArcs
func (mr *MockUpServerMockRecorder) UpsArcs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsArcs", reflect.TypeOf((*MockUpServer)(nil).UpsArcs), arg0, arg1)
}

// UpCount mocks base method
func (m *MockUpServer) UpCount(arg0 context.Context, arg1 *UpCountReq) (*UpCountReply, error) {
	ret := m.ctrl.Call(m, "UpCount", arg0, arg1)
	ret0, _ := ret[0].(*UpCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpCount indicates an expected call of UpCount
func (mr *MockUpServerMockRecorder) UpCount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpCount", reflect.TypeOf((*MockUpServer)(nil).UpCount), arg0, arg1)
}

// UpsCount mocks base method
func (m *MockUpServer) UpsCount(arg0 context.Context, arg1 *UpsCountReq) (*UpsCountReply, error) {
	ret := m.ctrl.Call(m, "UpsCount", arg0, arg1)
	ret0, _ := ret[0].(*UpsCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsCount indicates an expected call of UpsCount
func (mr *MockUpServerMockRecorder) UpsCount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsCount", reflect.TypeOf((*MockUpServer)(nil).UpsCount), arg0, arg1)
}

// UpsAidPubTime mocks base method
func (m *MockUpServer) UpsAidPubTime(arg0 context.Context, arg1 *UpsArcsReq) (*UpsAidPubTimeReply, error) {
	ret := m.ctrl.Call(m, "UpsAidPubTime", arg0, arg1)
	ret0, _ := ret[0].(*UpsAidPubTimeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsAidPubTime indicates an expected call of UpsAidPubTime
func (mr *MockUpServerMockRecorder) UpsAidPubTime(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsAidPubTime", reflect.TypeOf((*MockUpServer)(nil).UpsAidPubTime), arg0, arg1)
}

// AddUpPassedCacheByStaff mocks base method
func (m *MockUpServer) AddUpPassedCacheByStaff(arg0 context.Context, arg1 *UpCacheReq) (*NoReply, error) {
	ret := m.ctrl.Call(m, "AddUpPassedCacheByStaff", arg0, arg1)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUpPassedCacheByStaff indicates an expected call of AddUpPassedCacheByStaff
func (mr *MockUpServerMockRecorder) AddUpPassedCacheByStaff(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpPassedCacheByStaff", reflect.TypeOf((*MockUpServer)(nil).AddUpPassedCacheByStaff), arg0, arg1)
}

// AddUpPassedCache mocks base method
func (m *MockUpServer) AddUpPassedCache(arg0 context.Context, arg1 *UpCacheReq) (*NoReply, error) {
	ret := m.ctrl.Call(m, "AddUpPassedCache", arg0, arg1)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUpPassedCache indicates an expected call of AddUpPassedCache
func (mr *MockUpServerMockRecorder) AddUpPassedCache(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpPassedCache", reflect.TypeOf((*MockUpServer)(nil).AddUpPassedCache), arg0, arg1)
}

// DelUpPassedCacheByStaff mocks base method
func (m *MockUpServer) DelUpPassedCacheByStaff(arg0 context.Context, arg1 *UpCacheReq) (*NoReply, error) {
	ret := m.ctrl.Call(m, "DelUpPassedCacheByStaff", arg0, arg1)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelUpPassedCacheByStaff indicates an expected call of DelUpPassedCacheByStaff
func (mr *MockUpServerMockRecorder) DelUpPassedCacheByStaff(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelUpPassedCacheByStaff", reflect.TypeOf((*MockUpServer)(nil).DelUpPassedCacheByStaff), arg0, arg1)
}

// DelUpPassedCache mocks base method
func (m *MockUpServer) DelUpPassedCache(arg0 context.Context, arg1 *UpCacheReq) (*NoReply, error) {
	ret := m.ctrl.Call(m, "DelUpPassedCache", arg0, arg1)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelUpPassedCache indicates an expected call of DelUpPassedCache
func (mr *MockUpServerMockRecorder) DelUpPassedCache(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelUpPassedCache", reflect.TypeOf((*MockUpServer)(nil).DelUpPassedCache), arg0, arg1)
}

// UpInfoActivitys mocks base method
func (m *MockUpServer) UpInfoActivitys(arg0 context.Context, arg1 *UpListByLastIDReq) (*UpActivityListReply, error) {
	ret := m.ctrl.Call(m, "UpInfoActivitys", arg0, arg1)
	ret0, _ := ret[0].(*UpActivityListReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpInfoActivitys indicates an expected call of UpInfoActivitys
func (mr *MockUpServerMockRecorder) UpInfoActivitys(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpInfoActivitys", reflect.TypeOf((*MockUpServer)(nil).UpInfoActivitys), arg0, arg1)
}

// UpSpecial mocks base method
func (m *MockUpServer) UpSpecial(arg0 context.Context, arg1 *UpSpecialReq) (*UpSpecialReply, error) {
	ret := m.ctrl.Call(m, "UpSpecial", arg0, arg1)
	ret0, _ := ret[0].(*UpSpecialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpSpecial indicates an expected call of UpSpecial
func (mr *MockUpServerMockRecorder) UpSpecial(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpSpecial", reflect.TypeOf((*MockUpServer)(nil).UpSpecial), arg0, arg1)
}

// UpsSpecial mocks base method
func (m *MockUpServer) UpsSpecial(arg0 context.Context, arg1 *UpsSpecialReq) (*UpsSpecialReply, error) {
	ret := m.ctrl.Call(m, "UpsSpecial", arg0, arg1)
	ret0, _ := ret[0].(*UpsSpecialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsSpecial indicates an expected call of UpsSpecial
func (mr *MockUpServerMockRecorder) UpsSpecial(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsSpecial", reflect.TypeOf((*MockUpServer)(nil).UpsSpecial), arg0, arg1)
}

// UpGroups mocks base method
func (m *MockUpServer) UpGroups(arg0 context.Context, arg1 *NoArgReq) (*UpGroupsReply, error) {
	ret := m.ctrl.Call(m, "UpGroups", arg0, arg1)
	ret0, _ := ret[0].(*UpGroupsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpGroups indicates an expected call of UpGroups
func (mr *MockUpServerMockRecorder) UpGroups(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpGroups", reflect.TypeOf((*MockUpServer)(nil).UpGroups), arg0, arg1)
}

// UpGroupMids mocks base method
func (m *MockUpServer) UpGroupMids(arg0 context.Context, arg1 *UpGroupMidsReq) (*UpGroupMidsReply, error) {
	ret := m.ctrl.Call(m, "UpGroupMids", arg0, arg1)
	ret0, _ := ret[0].(*UpGroupMidsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpGroupMids indicates an expected call of UpGroupMids
func (mr *MockUpServerMockRecorder) UpGroupMids(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpGroupMids", reflect.TypeOf((*MockUpServer)(nil).UpGroupMids), arg0, arg1)
}

// UpAttr mocks base method
func (m *MockUpServer) UpAttr(arg0 context.Context, arg1 *UpAttrReq) (*UpAttrReply, error) {
	ret := m.ctrl.Call(m, "UpAttr", arg0, arg1)
	ret0, _ := ret[0].(*UpAttrReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpAttr indicates an expected call of UpAttr
func (mr *MockUpServerMockRecorder) UpAttr(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpAttr", reflect.TypeOf((*MockUpServer)(nil).UpAttr), arg0, arg1)
}

// UpBaseStats mocks base method
func (m *MockUpServer) UpBaseStats(arg0 context.Context, arg1 *UpStatReq) (*UpBaseStatReply, error) {
	ret := m.ctrl.Call(m, "UpBaseStats", arg0, arg1)
	ret0, _ := ret[0].(*UpBaseStatReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpBaseStats indicates an expected call of UpBaseStats
func (mr *MockUpServerMockRecorder) UpBaseStats(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpBaseStats", reflect.TypeOf((*MockUpServer)(nil).UpBaseStats), arg0, arg1)
}

// SetUpSwitch mocks base method
func (m *MockUpServer) SetUpSwitch(arg0 context.Context, arg1 *UpSwitchReq) (*NoReply, error) {
	ret := m.ctrl.Call(m, "SetUpSwitch", arg0, arg1)
	ret0, _ := ret[0].(*NoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUpSwitch indicates an expected call of SetUpSwitch
func (mr *MockUpServerMockRecorder) SetUpSwitch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpSwitch", reflect.TypeOf((*MockUpServer)(nil).SetUpSwitch), arg0, arg1)
}

// UpSwitch mocks base method
func (m *MockUpServer) UpSwitch(arg0 context.Context, arg1 *UpSwitchReq) (*UpSwitchReply, error) {
	ret := m.ctrl.Call(m, "UpSwitch", arg0, arg1)
	ret0, _ := ret[0].(*UpSwitchReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpSwitch indicates an expected call of UpSwitch
func (mr *MockUpServerMockRecorder) UpSwitch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpSwitch", reflect.TypeOf((*MockUpServer)(nil).UpSwitch), arg0, arg1)
}

// GetHighAllyUps mocks base method
func (m *MockUpServer) GetHighAllyUps(arg0 context.Context, arg1 *HighAllyUpsReq) (*HighAllyUpsReply, error) {
	ret := m.ctrl.Call(m, "GetHighAllyUps", arg0, arg1)
	ret0, _ := ret[0].(*HighAllyUpsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighAllyUps indicates an expected call of GetHighAllyUps
func (mr *MockUpServerMockRecorder) GetHighAllyUps(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighAllyUps", reflect.TypeOf((*MockUpServer)(nil).GetHighAllyUps), arg0, arg1)
}
