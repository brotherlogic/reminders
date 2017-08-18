// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reminders.proto

/*
Package reminders is a generated protocol buffer package.

It is generated from these files:
	reminders.proto

It has these top-level messages:
	Reminder
	Empty
	ReminderList
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

type Reminder struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// Text day of the week
	DayOfWeek string `protobuf:"bytes,2,opt,name=day_of_week,json=dayOfWeek" json:"day_of_week,omitempty"`
	// The time this should next be run
	NextRunTime int64 `protobuf:"varint,3,opt,name=next_run_time,json=nextRunTime" json:"next_run_time,omitempty"`
}

func (m *Reminder) Reset()                    { *m = Reminder{} }
func (m *Reminder) String() string            { return proto.CompactTextString(m) }
func (*Reminder) ProtoMessage()               {}
func (*Reminder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReminderList struct {
	Reminders []*Reminder `protobuf:"bytes,1,rep,name=reminders" json:"reminders,omitempty"`
}

func (m *ReminderList) Reset()                    { *m = ReminderList{} }
func (m *ReminderList) String() string            { return proto.CompactTextString(m) }
func (*ReminderList) ProtoMessage()               {}
func (*ReminderList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReminderList) GetReminders() []*Reminder {
	if m != nil {
		return m.Reminders
	}
	return nil
}

func init() {
	proto.RegisterType((*Reminder)(nil), "reminders.Reminder")
	proto.RegisterType((*Empty)(nil), "reminders.Empty")
	proto.RegisterType((*ReminderList)(nil), "reminders.ReminderList")
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
	ListReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReminderList, error)
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

func (c *remindersClient) ListReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReminderList, error) {
	out := new(ReminderList)
	err := grpc.Invoke(ctx, "/reminders.Reminders/ListReminders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reminders service

type RemindersServer interface {
	AddReminder(context.Context, *Reminder) (*Empty, error)
	ListReminders(context.Context, *Empty) (*ReminderList, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reminders.proto",
}

func init() { proto.RegisterFile("reminders.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0xcd, 0xcd,
	0xcc, 0x4b, 0x49, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x25, 0x71, 0x71, 0x04, 0x41, 0x39, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x1c, 0x17, 0x77, 0x4a, 0x62, 0x65, 0x7c, 0x7e,
	0x5a, 0x7c, 0x79, 0x6a, 0x6a, 0xb6, 0x04, 0x13, 0x58, 0x8a, 0x33, 0x25, 0xb1, 0xd2, 0x3f, 0x2d,
	0x3c, 0x35, 0x35, 0x5b, 0x48, 0x89, 0x8b, 0x37, 0x2f, 0xb5, 0xa2, 0x24, 0xbe, 0xa8, 0x34, 0x2f,
	0xbe, 0x24, 0x33, 0x37, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x88, 0x1b, 0x24, 0x18, 0x54,
	0x9a, 0x17, 0x92, 0x99, 0x9b, 0xaa, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0xa9, 0xe4,
	0xc8, 0xc5, 0x03, 0xb3, 0xcc, 0x27, 0xb3, 0xb8, 0x44, 0xc8, 0x90, 0x0b, 0xe1, 0x12, 0x09, 0x46,
	0x05, 0x66, 0x0d, 0x6e, 0x23, 0x61, 0x3d, 0x84, 0x63, 0x61, 0x6a, 0x83, 0x10, 0xaa, 0x8c, 0x1a,
	0x19, 0xb9, 0x38, 0x61, 0xe2, 0xc5, 0x42, 0x66, 0x5c, 0xdc, 0x8e, 0x29, 0x29, 0x70, 0x0f, 0x60,
	0xd3, 0x2c, 0x25, 0x80, 0x24, 0x08, 0x71, 0x06, 0x83, 0x90, 0x0d, 0x17, 0x2f, 0xc8, 0x01, 0x08,
	0x83, 0x30, 0x14, 0x49, 0x89, 0x63, 0x31, 0x0b, 0xa4, 0x47, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x8a,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x36, 0x9c, 0xd4, 0x58, 0x01, 0x00, 0x00,
}
