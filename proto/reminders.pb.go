// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reminders.proto

/*
Package reminders is a generated protocol buffer package.

It is generated from these files:
	reminders.proto

It has these top-level messages:
	ReminderConfig
	Reminder
	Empty
	ReminderList
	TaskList
*/
package reminders

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The state the task is in
type Reminder_ReminderState int32

const (
	Reminder_REPEATING  Reminder_ReminderState = 0
	Reminder_ASSIGNED   Reminder_ReminderState = 1
	Reminder_COMPLETE   Reminder_ReminderState = 2
	Reminder_UNASSIGNED Reminder_ReminderState = 3
)

var Reminder_ReminderState_name = map[int32]string{
	0: "REPEATING",
	1: "ASSIGNED",
	2: "COMPLETE",
	3: "UNASSIGNED",
}
var Reminder_ReminderState_value = map[string]int32{
	"REPEATING":  0,
	"ASSIGNED":   1,
	"COMPLETE":   2,
	"UNASSIGNED": 3,
}

func (x Reminder_ReminderState) String() string {
	return proto.EnumName(Reminder_ReminderState_name, int32(x))
}
func (Reminder_ReminderState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Reminder_ReminderPeriod int32

const (
	Reminder_WEEKLY      Reminder_ReminderPeriod = 0
	Reminder_MONTHLY     Reminder_ReminderPeriod = 1
	Reminder_YEARLY      Reminder_ReminderPeriod = 2
	Reminder_HALF_YEARLY Reminder_ReminderPeriod = 3
	Reminder_DAILY       Reminder_ReminderPeriod = 4
)

var Reminder_ReminderPeriod_name = map[int32]string{
	0: "WEEKLY",
	1: "MONTHLY",
	2: "YEARLY",
	3: "HALF_YEARLY",
	4: "DAILY",
}
var Reminder_ReminderPeriod_value = map[string]int32{
	"WEEKLY":      0,
	"MONTHLY":     1,
	"YEARLY":      2,
	"HALF_YEARLY": 3,
	"DAILY":       4,
}

func (x Reminder_ReminderPeriod) String() string {
	return proto.EnumName(Reminder_ReminderPeriod_name, int32(x))
}
func (Reminder_ReminderPeriod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type ReminderConfig struct {
	List  *ReminderList `protobuf:"bytes,1,opt,name=list" json:"list,omitempty"`
	Tasks []*TaskList   `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *ReminderConfig) Reset()                    { *m = ReminderConfig{} }
func (m *ReminderConfig) String() string            { return proto.CompactTextString(m) }
func (*ReminderConfig) ProtoMessage()               {}
func (*ReminderConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReminderConfig) GetList() *ReminderList {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReminderConfig) GetTasks() []*TaskList {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Reminder struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// Text day of the week
	DayOfWeek string `protobuf:"bytes,2,opt,name=day_of_week,json=dayOfWeek" json:"day_of_week,omitempty"`
	// The time this should next be run
	NextRunTime  int64                  `protobuf:"varint,3,opt,name=next_run_time,json=nextRunTime" json:"next_run_time,omitempty"`
	CurrentState Reminder_ReminderState `protobuf:"varint,4,opt,name=current_state,json=currentState,enum=reminders.Reminder_ReminderState" json:"current_state,omitempty"`
	// Assigned state for a github task
	GithubId string `protobuf:"bytes,5,opt,name=github_id,json=githubId" json:"github_id,omitempty"`
	// The component this should filed against in la github
	GithubComponent string                  `protobuf:"bytes,6,opt,name=github_component,json=githubComponent" json:"github_component,omitempty"`
	RepeatPeriod    Reminder_ReminderPeriod `protobuf:"varint,7,opt,name=repeatPeriod,enum=reminders.Reminder_ReminderPeriod" json:"repeatPeriod,omitempty"`
	Uid             int64                   `protobuf:"varint,8,opt,name=uid" json:"uid,omitempty"`
}

func (m *Reminder) Reset()                    { *m = Reminder{} }
func (m *Reminder) String() string            { return proto.CompactTextString(m) }
func (*Reminder) ProtoMessage()               {}
func (*Reminder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reminder) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Reminder) GetDayOfWeek() string {
	if m != nil {
		return m.DayOfWeek
	}
	return ""
}

func (m *Reminder) GetNextRunTime() int64 {
	if m != nil {
		return m.NextRunTime
	}
	return 0
}

func (m *Reminder) GetCurrentState() Reminder_ReminderState {
	if m != nil {
		return m.CurrentState
	}
	return Reminder_REPEATING
}

func (m *Reminder) GetGithubId() string {
	if m != nil {
		return m.GithubId
	}
	return ""
}

func (m *Reminder) GetGithubComponent() string {
	if m != nil {
		return m.GithubComponent
	}
	return ""
}

func (m *Reminder) GetRepeatPeriod() Reminder_ReminderPeriod {
	if m != nil {
		return m.RepeatPeriod
	}
	return Reminder_WEEKLY
}

func (m *Reminder) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReminderList struct {
	Reminders []*Reminder `protobuf:"bytes,1,rep,name=reminders" json:"reminders,omitempty"`
}

func (m *ReminderList) Reset()                    { *m = ReminderList{} }
func (m *ReminderList) String() string            { return proto.CompactTextString(m) }
func (*ReminderList) ProtoMessage()               {}
func (*ReminderList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReminderList) GetReminders() []*Reminder {
	if m != nil {
		return m.Reminders
	}
	return nil
}

type TaskList struct {
	Name  string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tasks *ReminderList `protobuf:"bytes,2,opt,name=tasks" json:"tasks,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TaskList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskList) GetTasks() *ReminderList {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*ReminderConfig)(nil), "reminders.ReminderConfig")
	proto.RegisterType((*Reminder)(nil), "reminders.Reminder")
	proto.RegisterType((*Empty)(nil), "reminders.Empty")
	proto.RegisterType((*ReminderList)(nil), "reminders.ReminderList")
	proto.RegisterType((*TaskList)(nil), "reminders.TaskList")
	proto.RegisterEnum("reminders.Reminder_ReminderState", Reminder_ReminderState_name, Reminder_ReminderState_value)
	proto.RegisterEnum("reminders.Reminder_ReminderPeriod", Reminder_ReminderPeriod_name, Reminder_ReminderPeriod_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Reminders service

type RemindersClient interface {
	AddReminder(ctx context.Context, in *Reminder, opts ...grpc.CallOption) (*Empty, error)
	ListReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReminderConfig, error)
	AddTaskList(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*Empty, error)
}

type remindersClient struct {
	cc *grpc.ClientConn
}

func NewRemindersClient(cc *grpc.ClientConn) RemindersClient {
	return &remindersClient{cc}
}

func (c *remindersClient) AddReminder(ctx context.Context, in *Reminder, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/reminders.Reminders/AddReminder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindersClient) ListReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReminderConfig, error) {
	out := new(ReminderConfig)
	err := grpc.Invoke(ctx, "/reminders.Reminders/ListReminders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindersClient) AddTaskList(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/reminders.Reminders/AddTaskList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reminders service

type RemindersServer interface {
	AddReminder(context.Context, *Reminder) (*Empty, error)
	ListReminders(context.Context, *Empty) (*ReminderConfig, error)
	AddTaskList(context.Context, *TaskList) (*Empty, error)
}

func RegisterRemindersServer(s *grpc.Server, srv RemindersServer) {
	s.RegisterService(&_Reminders_serviceDesc, srv)
}

func _Reminders_AddReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reminder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemindersServer).AddReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminders.Reminders/AddReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemindersServer).AddReminder(ctx, req.(*Reminder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reminders_ListReminders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemindersServer).ListReminders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminders.Reminders/ListReminders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemindersServer).ListReminders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reminders_AddTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemindersServer).AddTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminders.Reminders/AddTaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemindersServer).AddTaskList(ctx, req.(*TaskList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reminders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reminders.Reminders",
	HandlerType: (*RemindersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReminder",
			Handler:    _Reminders_AddReminder_Handler,
		},
		{
			MethodName: "ListReminders",
			Handler:    _Reminders_ListReminders_Handler,
		},
		{
			MethodName: "AddTaskList",
			Handler:    _Reminders_AddTaskList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reminders.proto",
}

func init() { proto.RegisterFile("reminders.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdb, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x2d, 0xcb, 0x27, 0x8d, 0x7c, 0x10, 0xd3, 0x8b, 0xaa, 0x29, 0x14, 0x57, 0x57, 0x0e,
	0xa5, 0x81, 0xba, 0xd0, 0xcb, 0x82, 0x70, 0xe4, 0xc4, 0xad, 0x7c, 0xe8, 0xda, 0x25, 0xf8, 0x4a,
	0x28, 0xd1, 0x3a, 0x59, 0x5c, 0x49, 0x46, 0x5a, 0xd3, 0xf8, 0xd9, 0xfa, 0x48, 0x7d, 0x89, 0xa0,
	0xd5, 0xc1, 0x0e, 0x11, 0xb9, 0x9b, 0xfd, 0xe7, 0xdf, 0x39, 0xc0, 0x37, 0xd0, 0x8b, 0xa8, 0xcf,
	0x02, 0x8f, 0x46, 0xf1, 0xc5, 0x2e, 0x0a, 0x79, 0x88, 0x4a, 0x21, 0x18, 0x0f, 0xd0, 0x25, 0xd9,
	0x63, 0x14, 0x06, 0x1b, 0x76, 0x8f, 0x9f, 0xa0, 0xf6, 0x87, 0xc5, 0x5c, 0x97, 0xfa, 0xd2, 0x40,
	0x1d, 0xbe, 0xbd, 0x38, 0x7e, 0xce, 0x8d, 0x36, 0x8b, 0x39, 0x11, 0x26, 0x3c, 0x87, 0x3a, 0x77,
	0xe3, 0x6d, 0xac, 0x57, 0xfb, 0xf2, 0x40, 0x1d, 0xbe, 0x39, 0x71, 0xaf, 0xdc, 0x78, 0x2b, 0x9c,
	0xa9, 0xc3, 0xf8, 0x2f, 0x43, 0x2b, 0xaf, 0x80, 0x08, 0x35, 0x4e, 0x1f, 0xd3, 0x26, 0x0a, 0x11,
	0x31, 0x7e, 0x00, 0xd5, 0x73, 0x0f, 0x4e, 0xb8, 0x71, 0xfe, 0x52, 0xba, 0xd5, 0xab, 0x22, 0xa5,
	0x78, 0xee, 0x61, 0xbe, 0xb9, 0xa1, 0x74, 0x8b, 0x06, 0x74, 0x02, 0xfa, 0xc8, 0x9d, 0x68, 0x1f,
	0x38, 0x9c, 0xf9, 0x54, 0x97, 0xfb, 0xd2, 0x40, 0x26, 0x6a, 0x22, 0x92, 0x7d, 0xb0, 0x62, 0x3e,
	0xc5, 0x31, 0x74, 0xee, 0xf6, 0x51, 0x44, 0x03, 0xee, 0xc4, 0xdc, 0xe5, 0x54, 0xaf, 0xf5, 0xa5,
	0x41, 0x77, 0xf8, 0xb1, 0x64, 0x8b, 0x22, 0x58, 0x26, 0x46, 0xd2, 0xce, 0xfe, 0x89, 0x17, 0xbe,
	0x07, 0xe5, 0x9e, 0xf1, 0x87, 0xfd, 0xad, 0xc3, 0x3c, 0xbd, 0x2e, 0x26, 0x69, 0xa5, 0xc2, 0xc4,
	0xc3, 0x73, 0xd0, 0xb2, 0xe4, 0x5d, 0xe8, 0xef, 0xc2, 0x80, 0x06, 0x5c, 0x6f, 0x08, 0x4f, 0x2f,
	0xd5, 0x47, 0xb9, 0x8c, 0x63, 0x68, 0x47, 0x74, 0x47, 0x5d, 0xbe, 0xa0, 0x11, 0x0b, 0x3d, 0xbd,
	0x29, 0xc6, 0x31, 0x5e, 0x1b, 0x27, 0x75, 0x92, 0x67, 0xff, 0x50, 0x03, 0x79, 0xcf, 0x3c, 0xbd,
	0x25, 0x36, 0x4e, 0x42, 0xe3, 0x07, 0x74, 0x9e, 0x2d, 0x80, 0x1d, 0x50, 0x88, 0xb5, 0xb0, 0xcc,
	0xd5, 0x64, 0x76, 0xa5, 0x55, 0xb0, 0x0d, 0x2d, 0x73, 0xb9, 0x9c, 0x5c, 0xcd, 0xac, 0x4b, 0x4d,
	0x4a, 0x5e, 0xa3, 0xf9, 0x74, 0x61, 0x5b, 0x2b, 0x4b, 0xab, 0x62, 0x17, 0xe0, 0xf7, 0xac, 0xc8,
	0xca, 0xc6, 0xaf, 0x23, 0x04, 0x59, 0x3f, 0x80, 0xc6, 0x8d, 0x65, 0xfd, 0xb4, 0xd7, 0x5a, 0x05,
	0x55, 0x68, 0x4e, 0xe7, 0xb3, 0xd5, 0xb5, 0xbd, 0xd6, 0xa4, 0x24, 0xb1, 0xb6, 0x4c, 0x62, 0xaf,
	0xb5, 0x2a, 0xf6, 0x40, 0xbd, 0x36, 0xed, 0xb1, 0x93, 0x09, 0x32, 0x2a, 0x50, 0xbf, 0x34, 0x27,
	0xf6, 0x5a, 0xab, 0x19, 0x4d, 0xa8, 0x5b, 0xfe, 0x8e, 0x1f, 0x0c, 0x13, 0xda, 0xa7, 0xdc, 0xe0,
	0x17, 0x38, 0xd2, 0xa7, 0x4b, 0x2f, 0xa8, 0xc9, 0xbd, 0xe4, 0x84, 0xd1, 0x29, 0xb4, 0x72, 0x98,
	0x12, 0x70, 0x02, 0xd7, 0xa7, 0x39, 0x38, 0x49, 0x8c, 0x9f, 0x8f, 0x10, 0xbe, 0x8a, 0x6c, 0xea,
	0x1a, 0xfe, 0x93, 0x40, 0xc9, 0xf5, 0x18, 0xbf, 0x81, 0x6a, 0x7a, 0x5e, 0x01, 0x66, 0xd9, 0x2c,
	0x67, 0xda, 0x89, 0x98, 0x6e, 0x55, 0xc1, 0xef, 0xd0, 0x11, 0x45, 0x8b, 0x42, 0x2f, 0x4c, 0x67,
	0xef, 0x4a, 0x6a, 0xa5, 0x47, 0x66, 0x54, 0xb2, 0xbe, 0xc5, 0x5e, 0x65, 0x97, 0x53, 0xd6, 0xf7,
	0xb6, 0x21, 0x4e, 0xf8, 0xeb, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x57, 0xae, 0x24, 0xd5,
	0x03, 0x00, 0x00,
}
