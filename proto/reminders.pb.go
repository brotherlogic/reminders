// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reminders.proto

package reminders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

func (Reminder_ReminderState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{1, 0}
}

type Reminder_ReminderPeriod int32

const (
	Reminder_WEEKLY      Reminder_ReminderPeriod = 0
	Reminder_MONTHLY     Reminder_ReminderPeriod = 1
	Reminder_YEARLY      Reminder_ReminderPeriod = 2
	Reminder_HALF_YEARLY Reminder_ReminderPeriod = 3
	Reminder_DAILY       Reminder_ReminderPeriod = 4
	Reminder_BIWEEKLY    Reminder_ReminderPeriod = 5
)

var Reminder_ReminderPeriod_name = map[int32]string{
	0: "WEEKLY",
	1: "MONTHLY",
	2: "YEARLY",
	3: "HALF_YEARLY",
	4: "DAILY",
	5: "BIWEEKLY",
}

var Reminder_ReminderPeriod_value = map[string]int32{
	"WEEKLY":      0,
	"MONTHLY":     1,
	"YEARLY":      2,
	"HALF_YEARLY": 3,
	"DAILY":       4,
	"BIWEEKLY":    5,
}

func (x Reminder_ReminderPeriod) String() string {
	return proto.EnumName(Reminder_ReminderPeriod_name, int32(x))
}

func (Reminder_ReminderPeriod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{1, 1}
}

type ReminderConfig struct {
	List                 *ReminderList `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
	Tasks                []*TaskList   `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReminderConfig) Reset()         { *m = ReminderConfig{} }
func (m *ReminderConfig) String() string { return proto.CompactTextString(m) }
func (*ReminderConfig) ProtoMessage()    {}
func (*ReminderConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{0}
}

func (m *ReminderConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReminderConfig.Unmarshal(m, b)
}
func (m *ReminderConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReminderConfig.Marshal(b, m, deterministic)
}
func (m *ReminderConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReminderConfig.Merge(m, src)
}
func (m *ReminderConfig) XXX_Size() int {
	return xxx_messageInfo_ReminderConfig.Size(m)
}
func (m *ReminderConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ReminderConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ReminderConfig proto.InternalMessageInfo

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
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Text day of the week
	DayOfWeek string `protobuf:"bytes,2,opt,name=day_of_week,json=dayOfWeek,proto3" json:"day_of_week,omitempty"`
	// The time this should next be run
	NextRunTime  int64                  `protobuf:"varint,3,opt,name=next_run_time,json=nextRunTime,proto3" json:"next_run_time,omitempty"`
	CurrentState Reminder_ReminderState `protobuf:"varint,4,opt,name=current_state,json=currentState,proto3,enum=reminders.Reminder_ReminderState" json:"current_state,omitempty"`
	// Assigned state for a github task
	GithubId string `protobuf:"bytes,5,opt,name=github_id,json=githubId,proto3" json:"github_id,omitempty"`
	// The component this should filed against in la github
	GithubComponent string                  `protobuf:"bytes,6,opt,name=github_component,json=githubComponent,proto3" json:"github_component,omitempty"`
	RepeatPeriod    Reminder_ReminderPeriod `protobuf:"varint,7,opt,name=repeatPeriod,proto3,enum=reminders.Reminder_ReminderPeriod" json:"repeatPeriod,omitempty"`
	Uid             int64                   `protobuf:"varint,8,opt,name=uid,proto3" json:"uid,omitempty"`
	// For something more freeform
	RepeatPeriodInSeconds int64    `protobuf:"varint,9,opt,name=repeat_period_in_seconds,json=repeatPeriodInSeconds,proto3" json:"repeat_period_in_seconds,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Reminder) Reset()         { *m = Reminder{} }
func (m *Reminder) String() string { return proto.CompactTextString(m) }
func (*Reminder) ProtoMessage()    {}
func (*Reminder) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{1}
}

func (m *Reminder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reminder.Unmarshal(m, b)
}
func (m *Reminder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reminder.Marshal(b, m, deterministic)
}
func (m *Reminder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reminder.Merge(m, src)
}
func (m *Reminder) XXX_Size() int {
	return xxx_messageInfo_Reminder.Size(m)
}
func (m *Reminder) XXX_DiscardUnknown() {
	xxx_messageInfo_Reminder.DiscardUnknown(m)
}

var xxx_messageInfo_Reminder proto.InternalMessageInfo

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

func (m *Reminder) GetRepeatPeriodInSeconds() int64 {
	if m != nil {
		return m.RepeatPeriodInSeconds
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ReminderList struct {
	Reminders            []*Reminder `protobuf:"bytes,1,rep,name=reminders,proto3" json:"reminders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReminderList) Reset()         { *m = ReminderList{} }
func (m *ReminderList) String() string { return proto.CompactTextString(m) }
func (*ReminderList) ProtoMessage()    {}
func (*ReminderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{3}
}

func (m *ReminderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReminderList.Unmarshal(m, b)
}
func (m *ReminderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReminderList.Marshal(b, m, deterministic)
}
func (m *ReminderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReminderList.Merge(m, src)
}
func (m *ReminderList) XXX_Size() int {
	return xxx_messageInfo_ReminderList.Size(m)
}
func (m *ReminderList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReminderList.DiscardUnknown(m)
}

var xxx_messageInfo_ReminderList proto.InternalMessageInfo

func (m *ReminderList) GetReminders() []*Reminder {
	if m != nil {
		return m.Reminders
	}
	return nil
}

type TaskList struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tasks                *ReminderList `protobuf:"bytes,2,opt,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{4}
}

func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (m *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(m, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

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

type DeleteRequest struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{5}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c182e1a34f0cb2e5, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("reminders.Reminder_ReminderState", Reminder_ReminderState_name, Reminder_ReminderState_value)
	proto.RegisterEnum("reminders.Reminder_ReminderPeriod", Reminder_ReminderPeriod_name, Reminder_ReminderPeriod_value)
	proto.RegisterType((*ReminderConfig)(nil), "reminders.ReminderConfig")
	proto.RegisterType((*Reminder)(nil), "reminders.Reminder")
	proto.RegisterType((*Empty)(nil), "reminders.Empty")
	proto.RegisterType((*ReminderList)(nil), "reminders.ReminderList")
	proto.RegisterType((*TaskList)(nil), "reminders.TaskList")
	proto.RegisterType((*DeleteRequest)(nil), "reminders.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "reminders.DeleteResponse")
}

func init() { proto.RegisterFile("reminders.proto", fileDescriptor_c182e1a34f0cb2e5) }

var fileDescriptor_c182e1a34f0cb2e5 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4b, 0x4f, 0xdb, 0x40,
	0x10, 0xc7, 0xe3, 0x3c, 0x20, 0x1e, 0x27, 0xc1, 0x9a, 0xaa, 0xea, 0x96, 0x4a, 0x55, 0xf0, 0x29,
	0xa8, 0x2a, 0x52, 0x53, 0xa9, 0xbd, 0x55, 0x72, 0x83, 0x81, 0xb4, 0x26, 0x20, 0x27, 0x15, 0xca,
	0xc9, 0x32, 0x78, 0x00, 0x2b, 0x64, 0x9d, 0x7a, 0x37, 0x2a, 0xdc, 0xfa, 0x99, 0xfa, 0x09, 0x2b,
	0xaf, 0x1f, 0x49, 0x44, 0xc4, 0x6d, 0x76, 0xe6, 0xb7, 0xf3, 0xf2, 0x7f, 0x0d, 0x7b, 0x09, 0xcd,
	0x23, 0x1e, 0x52, 0x22, 0x8e, 0x16, 0x49, 0x2c, 0x63, 0xd4, 0x4b, 0x87, 0x75, 0x0f, 0x1d, 0x2f,
	0x3f, 0x0c, 0x62, 0x7e, 0x1b, 0xdd, 0xe1, 0x07, 0xa8, 0x3f, 0x44, 0x42, 0x32, 0xad, 0xab, 0xf5,
	0x8c, 0xfe, 0x9b, 0xa3, 0xd5, 0xe5, 0x02, 0x74, 0x23, 0x21, 0x3d, 0x05, 0xe1, 0x21, 0x34, 0x64,
	0x20, 0x66, 0x82, 0x55, 0xbb, 0xb5, 0x9e, 0xd1, 0x7f, 0xb5, 0x46, 0x4f, 0x02, 0x31, 0x53, 0x64,
	0x46, 0x58, 0xff, 0xea, 0xd0, 0x2c, 0x32, 0x20, 0x42, 0x5d, 0xd2, 0x63, 0x56, 0x44, 0xf7, 0x94,
	0x8d, 0xef, 0xc1, 0x08, 0x83, 0x27, 0x3f, 0xbe, 0xf5, 0xff, 0x10, 0xcd, 0x58, 0x55, 0x85, 0xf4,
	0x30, 0x78, 0xba, 0xb8, 0xbd, 0x22, 0x9a, 0xa1, 0x05, 0x6d, 0x4e, 0x8f, 0xd2, 0x4f, 0x96, 0xdc,
	0x97, 0xd1, 0x9c, 0x58, 0xad, 0xab, 0xf5, 0x6a, 0x9e, 0x91, 0x3a, 0xbd, 0x25, 0x9f, 0x44, 0x73,
	0xc2, 0x13, 0x68, 0xdf, 0x2c, 0x93, 0x84, 0xb8, 0xf4, 0x85, 0x0c, 0x24, 0xb1, 0x7a, 0x57, 0xeb,
	0x75, 0xfa, 0x07, 0x5b, 0xa6, 0x28, 0x8d, 0x71, 0x0a, 0x7a, 0xad, 0xfc, 0x9e, 0x3a, 0xe1, 0x3b,
	0xd0, 0xef, 0x22, 0x79, 0xbf, 0xbc, 0xf6, 0xa3, 0x90, 0x35, 0x54, 0x27, 0xcd, 0xcc, 0x31, 0x0c,
	0xf1, 0x10, 0xcc, 0x3c, 0x78, 0x13, 0xcf, 0x17, 0x31, 0x27, 0x2e, 0xd9, 0x8e, 0x62, 0xf6, 0x32,
	0xff, 0xa0, 0x70, 0xe3, 0x09, 0xb4, 0x12, 0x5a, 0x50, 0x20, 0x2f, 0x29, 0x89, 0xe2, 0x90, 0xed,
	0xaa, 0x76, 0xac, 0x97, 0xda, 0xc9, 0x48, 0x6f, 0xe3, 0x1e, 0x9a, 0x50, 0x5b, 0x46, 0x21, 0x6b,
	0xaa, 0x89, 0x53, 0x13, 0xbf, 0x02, 0xcb, 0x08, 0x7f, 0xa1, 0x10, 0x3f, 0xe2, 0xbe, 0xa0, 0x9b,
	0x98, 0x87, 0x82, 0xe9, 0x0a, 0x7b, 0xbd, 0x9e, 0x61, 0xc8, 0xc7, 0x59, 0xd0, 0xfa, 0x01, 0xed,
	0x8d, 0xc9, 0xb1, 0x0d, 0xba, 0xe7, 0x5c, 0x3a, 0xf6, 0x64, 0x38, 0x3a, 0x35, 0x2b, 0xd8, 0x82,
	0xa6, 0x3d, 0x1e, 0x0f, 0x4f, 0x47, 0xce, 0xb1, 0xa9, 0xa5, 0xa7, 0xc1, 0xc5, 0xf9, 0xa5, 0xeb,
	0x4c, 0x1c, 0xb3, 0x8a, 0x1d, 0x80, 0x5f, 0xa3, 0x32, 0x5a, 0xb3, 0xfc, 0x95, 0x7a, 0xf2, 0x46,
	0x01, 0x76, 0xae, 0x1c, 0xe7, 0xa7, 0x3b, 0x35, 0x2b, 0x68, 0xc0, 0xee, 0xf9, 0xc5, 0x68, 0x72,
	0xe6, 0x4e, 0x4d, 0x2d, 0x0d, 0x4c, 0x1d, 0xdb, 0x73, 0xa7, 0x66, 0x15, 0xf7, 0xc0, 0x38, 0xb3,
	0xdd, 0x13, 0x3f, 0x77, 0xd4, 0x50, 0x87, 0xc6, 0xb1, 0x3d, 0x74, 0xa7, 0x66, 0x3d, 0x2d, 0xf8,
	0x7d, 0x98, 0xa7, 0x68, 0x58, 0xbb, 0xd0, 0x70, 0xe6, 0x0b, 0xf9, 0x64, 0xd9, 0xd0, 0x5a, 0x97,
	0x1f, 0x7e, 0x82, 0x95, 0x88, 0x99, 0xf6, 0x4c, 0x7c, 0x05, 0xeb, 0xad, 0x49, 0xfd, 0x1c, 0x9a,
	0x85, 0x26, 0x53, 0xfd, 0xf1, 0x60, 0x4e, 0x85, 0xfe, 0x52, 0x1b, 0x3f, 0xae, 0xb4, 0xfc, 0xa2,
	0xf2, 0x73, 0x3d, 0x1f, 0x40, 0xfb, 0x98, 0x1e, 0x48, 0x92, 0x47, 0xbf, 0x97, 0x24, 0x64, 0xf1,
	0x8d, 0xb4, 0xf2, 0x1b, 0x59, 0x26, 0x74, 0x0a, 0x44, 0x2c, 0x62, 0x2e, 0xa8, 0xff, 0xb7, 0x0a,
	0x7a, 0x91, 0x4c, 0xe0, 0x17, 0x30, 0xec, 0x30, 0x2c, 0x1f, 0xc5, 0xb6, 0x01, 0xf6, 0xcd, 0x35,
	0x67, 0xb6, 0x8a, 0x0a, 0x7e, 0x83, 0xb6, 0xea, 0xa4, 0x4c, 0xf4, 0x0c, 0xda, 0x7f, 0xbb, 0x25,
	0x57, 0xf6, 0xc0, 0xad, 0x4a, 0x5e, 0xb7, 0x5c, 0xc6, 0xb6, 0x57, 0xbb, 0xb5, 0xee, 0x00, 0x20,
	0x9b, 0x27, 0xa5, 0x90, 0xad, 0x11, 0x1b, 0x9b, 0xd8, 0x28, 0xbe, 0xb9, 0x00, 0xab, 0x72, 0xbd,
	0xa3, 0xfe, 0x41, 0x9f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x3e, 0x12, 0x7d, 0x96, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemindersClient is the client API for Reminders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemindersClient interface {
	AddReminder(ctx context.Context, in *Reminder, opts ...grpc.CallOption) (*Empty, error)
	ListReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReminderConfig, error)
	AddTaskList(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*Empty, error)
	DeleteTask(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type remindersClient struct {
	cc *grpc.ClientConn
}

func NewRemindersClient(cc *grpc.ClientConn) RemindersClient {
	return &remindersClient{cc}
}

func (c *remindersClient) AddReminder(ctx context.Context, in *Reminder, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/reminders.Reminders/AddReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindersClient) ListReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReminderConfig, error) {
	out := new(ReminderConfig)
	err := c.cc.Invoke(ctx, "/reminders.Reminders/ListReminders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindersClient) AddTaskList(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/reminders.Reminders/AddTaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindersClient) DeleteTask(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/reminders.Reminders/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemindersServer is the server API for Reminders service.
type RemindersServer interface {
	AddReminder(context.Context, *Reminder) (*Empty, error)
	ListReminders(context.Context, *Empty) (*ReminderConfig, error)
	AddTaskList(context.Context, *TaskList) (*Empty, error)
	DeleteTask(context.Context, *DeleteRequest) (*DeleteResponse, error)
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

func _Reminders_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemindersServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminders.Reminders/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemindersServer).DeleteTask(ctx, req.(*DeleteRequest))
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
		{
			MethodName: "DeleteTask",
			Handler:    _Reminders_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reminders.proto",
}
