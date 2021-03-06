// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: reminders.proto

package reminders

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The state the task is in
type Reminder_ReminderState int32

const (
	Reminder_REPEATING  Reminder_ReminderState = 0
	Reminder_ASSIGNED   Reminder_ReminderState = 1
	Reminder_COMPLETE   Reminder_ReminderState = 2
	Reminder_UNASSIGNED Reminder_ReminderState = 3
)

// Enum value maps for Reminder_ReminderState.
var (
	Reminder_ReminderState_name = map[int32]string{
		0: "REPEATING",
		1: "ASSIGNED",
		2: "COMPLETE",
		3: "UNASSIGNED",
	}
	Reminder_ReminderState_value = map[string]int32{
		"REPEATING":  0,
		"ASSIGNED":   1,
		"COMPLETE":   2,
		"UNASSIGNED": 3,
	}
)

func (x Reminder_ReminderState) Enum() *Reminder_ReminderState {
	p := new(Reminder_ReminderState)
	*p = x
	return p
}

func (x Reminder_ReminderState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reminder_ReminderState) Descriptor() protoreflect.EnumDescriptor {
	return file_reminders_proto_enumTypes[0].Descriptor()
}

func (Reminder_ReminderState) Type() protoreflect.EnumType {
	return &file_reminders_proto_enumTypes[0]
}

func (x Reminder_ReminderState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reminder_ReminderState.Descriptor instead.
func (Reminder_ReminderState) EnumDescriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{1, 0}
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

// Enum value maps for Reminder_ReminderPeriod.
var (
	Reminder_ReminderPeriod_name = map[int32]string{
		0: "WEEKLY",
		1: "MONTHLY",
		2: "YEARLY",
		3: "HALF_YEARLY",
		4: "DAILY",
		5: "BIWEEKLY",
	}
	Reminder_ReminderPeriod_value = map[string]int32{
		"WEEKLY":      0,
		"MONTHLY":     1,
		"YEARLY":      2,
		"HALF_YEARLY": 3,
		"DAILY":       4,
		"BIWEEKLY":    5,
	}
)

func (x Reminder_ReminderPeriod) Enum() *Reminder_ReminderPeriod {
	p := new(Reminder_ReminderPeriod)
	*p = x
	return p
}

func (x Reminder_ReminderPeriod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reminder_ReminderPeriod) Descriptor() protoreflect.EnumDescriptor {
	return file_reminders_proto_enumTypes[1].Descriptor()
}

func (Reminder_ReminderPeriod) Type() protoreflect.EnumType {
	return &file_reminders_proto_enumTypes[1]
}

func (x Reminder_ReminderPeriod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reminder_ReminderPeriod.Descriptor instead.
func (Reminder_ReminderPeriod) EnumDescriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{1, 1}
}

type ReminderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  *ReminderList `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
	Tasks []*TaskList   `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ReminderConfig) Reset() {
	*x = ReminderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderConfig) ProtoMessage() {}

func (x *ReminderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderConfig.ProtoReflect.Descriptor instead.
func (*ReminderConfig) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{0}
}

func (x *ReminderConfig) GetList() *ReminderList {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ReminderConfig) GetTasks() []*TaskList {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Reminder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text   string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Server string `protobuf:"bytes,11,opt,name=server,proto3" json:"server,omitempty"`
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
	RepeatPeriodInSeconds int64 `protobuf:"varint,9,opt,name=repeat_period_in_seconds,json=repeatPeriodInSeconds,proto3" json:"repeat_period_in_seconds,omitempty"`
	// To enable alert silencing
	Silences []string `protobuf:"bytes,10,rep,name=silences,proto3" json:"silences,omitempty"`
}

func (x *Reminder) Reset() {
	*x = Reminder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reminder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reminder) ProtoMessage() {}

func (x *Reminder) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reminder.ProtoReflect.Descriptor instead.
func (*Reminder) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{1}
}

func (x *Reminder) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Reminder) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *Reminder) GetDayOfWeek() string {
	if x != nil {
		return x.DayOfWeek
	}
	return ""
}

func (x *Reminder) GetNextRunTime() int64 {
	if x != nil {
		return x.NextRunTime
	}
	return 0
}

func (x *Reminder) GetCurrentState() Reminder_ReminderState {
	if x != nil {
		return x.CurrentState
	}
	return Reminder_REPEATING
}

func (x *Reminder) GetGithubId() string {
	if x != nil {
		return x.GithubId
	}
	return ""
}

func (x *Reminder) GetGithubComponent() string {
	if x != nil {
		return x.GithubComponent
	}
	return ""
}

func (x *Reminder) GetRepeatPeriod() Reminder_ReminderPeriod {
	if x != nil {
		return x.RepeatPeriod
	}
	return Reminder_WEEKLY
}

func (x *Reminder) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Reminder) GetRepeatPeriodInSeconds() int64 {
	if x != nil {
		return x.RepeatPeriodInSeconds
	}
	return 0
}

func (x *Reminder) GetSilences() []string {
	if x != nil {
		return x.Silences
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{2}
}

type ReminderList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reminders []*Reminder `protobuf:"bytes,1,rep,name=reminders,proto3" json:"reminders,omitempty"`
}

func (x *ReminderList) Reset() {
	*x = ReminderList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderList) ProtoMessage() {}

func (x *ReminderList) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderList.ProtoReflect.Descriptor instead.
func (*ReminderList) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{3}
}

func (x *ReminderList) GetReminders() []*Reminder {
	if x != nil {
		return x.Reminders
	}
	return nil
}

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tasks *ReminderList `protobuf:"bytes,2,opt,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{4}
}

func (x *TaskList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskList) GetTasks() *ReminderList {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{6}
}

type ReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveRequest) Reset() {
	*x = ReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveRequest) ProtoMessage() {}

func (x *ReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveRequest.ProtoReflect.Descriptor instead.
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{7}
}

type ReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveResponse) Reset() {
	*x = ReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reminders_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveResponse) ProtoMessage() {}

func (x *ReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reminders_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveResponse.ProtoReflect.Descriptor instead.
func (*ReceiveResponse) Descriptor() ([]byte, []int) {
	return file_reminders_proto_rawDescGZIP(), []int{8}
}

var File_reminders_proto protoreflect.FileDescriptor

var file_reminders_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x68, 0x0a, 0x0e,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0xe6, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0b, 0x64, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12,
	0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0c, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x37, 0x0a,
	0x18, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69,
	0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x22, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x4e, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x22, 0x5f,
	0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x45, 0x4b, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x41,
	0x52, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x41, 0x4c, 0x46, 0x5f, 0x59, 0x45,
	0x41, 0x52, 0x4c, 0x59, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x49, 0x57, 0x45, 0x45, 0x4b, 0x4c, 0x59, 0x10, 0x05, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x10, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x10, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x80, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x72, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e,
	0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x56, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reminders_proto_rawDescOnce sync.Once
	file_reminders_proto_rawDescData = file_reminders_proto_rawDesc
)

func file_reminders_proto_rawDescGZIP() []byte {
	file_reminders_proto_rawDescOnce.Do(func() {
		file_reminders_proto_rawDescData = protoimpl.X.CompressGZIP(file_reminders_proto_rawDescData)
	})
	return file_reminders_proto_rawDescData
}

var file_reminders_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_reminders_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_reminders_proto_goTypes = []interface{}{
	(Reminder_ReminderState)(0),  // 0: reminders.Reminder.ReminderState
	(Reminder_ReminderPeriod)(0), // 1: reminders.Reminder.ReminderPeriod
	(*ReminderConfig)(nil),       // 2: reminders.ReminderConfig
	(*Reminder)(nil),             // 3: reminders.Reminder
	(*Empty)(nil),                // 4: reminders.Empty
	(*ReminderList)(nil),         // 5: reminders.ReminderList
	(*TaskList)(nil),             // 6: reminders.TaskList
	(*DeleteRequest)(nil),        // 7: reminders.DeleteRequest
	(*DeleteResponse)(nil),       // 8: reminders.DeleteResponse
	(*ReceiveRequest)(nil),       // 9: reminders.ReceiveRequest
	(*ReceiveResponse)(nil),      // 10: reminders.ReceiveResponse
}
var file_reminders_proto_depIdxs = []int32{
	5,  // 0: reminders.ReminderConfig.list:type_name -> reminders.ReminderList
	6,  // 1: reminders.ReminderConfig.tasks:type_name -> reminders.TaskList
	0,  // 2: reminders.Reminder.current_state:type_name -> reminders.Reminder.ReminderState
	1,  // 3: reminders.Reminder.repeatPeriod:type_name -> reminders.Reminder.ReminderPeriod
	3,  // 4: reminders.ReminderList.reminders:type_name -> reminders.Reminder
	5,  // 5: reminders.TaskList.tasks:type_name -> reminders.ReminderList
	3,  // 6: reminders.Reminders.AddReminder:input_type -> reminders.Reminder
	4,  // 7: reminders.Reminders.ListReminders:input_type -> reminders.Empty
	6,  // 8: reminders.Reminders.AddTaskList:input_type -> reminders.TaskList
	7,  // 9: reminders.Reminders.DeleteTask:input_type -> reminders.DeleteRequest
	9,  // 10: reminders.ReminderReceiver.Receive:input_type -> reminders.ReceiveRequest
	4,  // 11: reminders.Reminders.AddReminder:output_type -> reminders.Empty
	2,  // 12: reminders.Reminders.ListReminders:output_type -> reminders.ReminderConfig
	4,  // 13: reminders.Reminders.AddTaskList:output_type -> reminders.Empty
	8,  // 14: reminders.Reminders.DeleteTask:output_type -> reminders.DeleteResponse
	10, // 15: reminders.ReminderReceiver.Receive:output_type -> reminders.ReceiveResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_reminders_proto_init() }
func file_reminders_proto_init() {
	if File_reminders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reminders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reminder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reminders_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reminders_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_reminders_proto_goTypes,
		DependencyIndexes: file_reminders_proto_depIdxs,
		EnumInfos:         file_reminders_proto_enumTypes,
		MessageInfos:      file_reminders_proto_msgTypes,
	}.Build()
	File_reminders_proto = out.File
	file_reminders_proto_rawDesc = nil
	file_reminders_proto_goTypes = nil
	file_reminders_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

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
	cc grpc.ClientConnInterface
}

func NewRemindersClient(cc grpc.ClientConnInterface) RemindersClient {
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

// UnimplementedRemindersServer can be embedded to have forward compatible implementations.
type UnimplementedRemindersServer struct {
}

func (*UnimplementedRemindersServer) AddReminder(context.Context, *Reminder) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReminder not implemented")
}
func (*UnimplementedRemindersServer) ListReminders(context.Context, *Empty) (*ReminderConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReminders not implemented")
}
func (*UnimplementedRemindersServer) AddTaskList(context.Context, *TaskList) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTaskList not implemented")
}
func (*UnimplementedRemindersServer) DeleteTask(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
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

// ReminderReceiverClient is the client API for ReminderReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReminderReceiverClient interface {
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
}

type reminderReceiverClient struct {
	cc grpc.ClientConnInterface
}

func NewReminderReceiverClient(cc grpc.ClientConnInterface) ReminderReceiverClient {
	return &reminderReceiverClient{cc}
}

func (c *reminderReceiverClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/reminders.ReminderReceiver/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReminderReceiverServer is the server API for ReminderReceiver service.
type ReminderReceiverServer interface {
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
}

// UnimplementedReminderReceiverServer can be embedded to have forward compatible implementations.
type UnimplementedReminderReceiverServer struct {
}

func (*UnimplementedReminderReceiverServer) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}

func RegisterReminderReceiverServer(s *grpc.Server, srv ReminderReceiverServer) {
	s.RegisterService(&_ReminderReceiver_serviceDesc, srv)
}

func _ReminderReceiver_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderReceiverServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminders.ReminderReceiver/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderReceiverServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReminderReceiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reminders.ReminderReceiver",
	HandlerType: (*ReminderReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _ReminderReceiver_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reminders.proto",
}
