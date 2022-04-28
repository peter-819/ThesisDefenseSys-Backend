// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: teacher.proto

package teacher

import (
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

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{0}
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId string `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartTime  string `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    string `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *Schedule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Schedule) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Schedule) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type Defense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefenseId string `protobuf:"bytes,1,opt,name=defense_id,json=defenseId,proto3" json:"defense_id,omitempty"`
	StartTime string `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Defense) Reset() {
	*x = Defense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Defense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Defense) ProtoMessage() {}

func (x *Defense) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Defense.ProtoReflect.Descriptor instead.
func (*Defense) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *Defense) GetDefenseId() string {
	if x != nil {
		return x.DefenseId
	}
	return ""
}

func (x *Defense) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Defense) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsSecretary       string      `protobuf:"bytes,3,opt,name=is_secretary,json=isSecretary,proto3" json:"is_secretary,omitempty"`
	MaxDefensePerWeek int32       `protobuf:"varint,4,opt,name=max_defense_per_week,json=maxDefensePerWeek,proto3" json:"max_defense_per_week,omitempty"`
	PreferKeywords    []string    `protobuf:"bytes,5,rep,name=prefer_keywords,json=preferKeywords,proto3" json:"prefer_keywords,omitempty"`
	Schedules         []*Schedule `protobuf:"bytes,6,rep,name=schedules,proto3" json:"schedules,omitempty"`
	Defenses          []*Defense  `protobuf:"bytes,7,rep,name=defenses,proto3" json:"defenses,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *Teacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Teacher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Teacher) GetIsSecretary() string {
	if x != nil {
		return x.IsSecretary
	}
	return ""
}

func (x *Teacher) GetMaxDefensePerWeek() int32 {
	if x != nil {
		return x.MaxDefensePerWeek
	}
	return 0
}

func (x *Teacher) GetPreferKeywords() []string {
	if x != nil {
		return x.PreferKeywords
	}
	return nil
}

func (x *Teacher) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *Teacher) GetDefenses() []*Defense {
	if x != nil {
		return x.Defenses
	}
	return nil
}

type QueryTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QueryTeacherRequest) Reset() {
	*x = QueryTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTeacherRequest) ProtoMessage() {}

func (x *QueryTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTeacherRequest.ProtoReflect.Descriptor instead.
func (*QueryTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *QueryTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QueryTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teacher *Teacher `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
}

func (x *QueryTeacherResponse) Reset() {
	*x = QueryTeacherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTeacherResponse) ProtoMessage() {}

func (x *QueryTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTeacherResponse.ProtoReflect.Descriptor instead.
func (*QueryTeacherResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *QueryTeacherResponse) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type QueryAllTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryAllTeacherRequest) Reset() {
	*x = QueryAllTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllTeacherRequest) ProtoMessage() {}

func (x *QueryAllTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllTeacherRequest.ProtoReflect.Descriptor instead.
func (*QueryAllTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{6}
}

type QueryTeachersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Teacher `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *QueryTeachersResponse) Reset() {
	*x = QueryTeachersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTeachersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTeachersResponse) ProtoMessage() {}

func (x *QueryTeachersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTeachersResponse.ProtoReflect.Descriptor instead.
func (*QueryTeachersResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{7}
}

func (x *QueryTeachersResponse) GetList() []*Teacher {
	if x != nil {
		return x.List
	}
	return nil
}

type ModifyTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NewTeacher *Teacher `protobuf:"bytes,2,opt,name=new_teacher,json=newTeacher,proto3" json:"new_teacher,omitempty"`
}

func (x *ModifyTeacherRequest) Reset() {
	*x = ModifyTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyTeacherRequest) ProtoMessage() {}

func (x *ModifyTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyTeacherRequest.ProtoReflect.Descriptor instead.
func (*ModifyTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{8}
}

func (x *ModifyTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModifyTeacherRequest) GetNewTeacher() *Teacher {
	if x != nil {
		return x.NewTeacher
	}
	return nil
}

type RemoveTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveTeacherRequest) Reset() {
	*x = RemoveTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTeacherRequest) ProtoMessage() {}

func (x *RemoveTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTeacherRequest.ProtoReflect.Descriptor instead.
func (*RemoveTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teacher *Teacher `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
}

func (x *AddTeacherRequest) Reset() {
	*x = AddTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTeacherRequest) ProtoMessage() {}

func (x *AddTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTeacherRequest.ProtoReflect.Descriptor instead.
func (*AddTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{10}
}

func (x *AddTeacherRequest) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type QueryAvailableTeachersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords  []string `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Excluded  []string `protobuf:"bytes,2,rep,name=excluded,proto3" json:"excluded,omitempty"`
	StartTime string   `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string   `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *QueryAvailableTeachersRequest) Reset() {
	*x = QueryAvailableTeachersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAvailableTeachersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAvailableTeachersRequest) ProtoMessage() {}

func (x *QueryAvailableTeachersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAvailableTeachersRequest.ProtoReflect.Descriptor instead.
func (*QueryAvailableTeachersRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{11}
}

func (x *QueryAvailableTeachersRequest) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *QueryAvailableTeachersRequest) GetExcluded() []string {
	if x != nil {
		return x.Excluded
	}
	return nil
}

func (x *QueryAvailableTeachersRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *QueryAvailableTeachersRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

var File_teacher_proto protoreflect.FileDescriptor

var file_teacher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x0a, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x07, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x61, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x14, 0x6d,
	0x61, 0x78, 0x5f, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x77,
	0x65, 0x65, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x44, 0x65,
	0x66, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x65, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22,
	0x18, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x15, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x91, 0x01, 0x0a,
	0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0xe5, 0x03, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0d, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c,
	0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x12, 0x26, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teacher_proto_rawDescOnce sync.Once
	file_teacher_proto_rawDescData = file_teacher_proto_rawDesc
)

func file_teacher_proto_rawDescGZIP() []byte {
	file_teacher_proto_rawDescOnce.Do(func() {
		file_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_teacher_proto_rawDescData)
	})
	return file_teacher_proto_rawDescData
}

var file_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_teacher_proto_goTypes = []interface{}{
	(*EmptyResponse)(nil),                 // 0: teacher.EmptyResponse
	(*Schedule)(nil),                      // 1: teacher.Schedule
	(*Defense)(nil),                       // 2: teacher.Defense
	(*Teacher)(nil),                       // 3: teacher.Teacher
	(*QueryTeacherRequest)(nil),           // 4: teacher.QueryTeacherRequest
	(*QueryTeacherResponse)(nil),          // 5: teacher.QueryTeacherResponse
	(*QueryAllTeacherRequest)(nil),        // 6: teacher.QueryAllTeacherRequest
	(*QueryTeachersResponse)(nil),         // 7: teacher.QueryTeachersResponse
	(*ModifyTeacherRequest)(nil),          // 8: teacher.ModifyTeacherRequest
	(*RemoveTeacherRequest)(nil),          // 9: teacher.RemoveTeacherRequest
	(*AddTeacherRequest)(nil),             // 10: teacher.AddTeacherRequest
	(*QueryAvailableTeachersRequest)(nil), // 11: teacher.QueryAvailableTeachersRequest
}
var file_teacher_proto_depIdxs = []int32{
	1,  // 0: teacher.Teacher.schedules:type_name -> teacher.Schedule
	2,  // 1: teacher.Teacher.defenses:type_name -> teacher.Defense
	3,  // 2: teacher.QueryTeacherResponse.teacher:type_name -> teacher.Teacher
	3,  // 3: teacher.QueryTeachersResponse.list:type_name -> teacher.Teacher
	3,  // 4: teacher.ModifyTeacherRequest.new_teacher:type_name -> teacher.Teacher
	3,  // 5: teacher.AddTeacherRequest.teacher:type_name -> teacher.Teacher
	10, // 6: teacher.TeacherService.AddTeacher:input_type -> teacher.AddTeacherRequest
	9,  // 7: teacher.TeacherService.RemoveTeacher:input_type -> teacher.RemoveTeacherRequest
	8,  // 8: teacher.TeacherService.ModifyTeacher:input_type -> teacher.ModifyTeacherRequest
	6,  // 9: teacher.TeacherService.QueryAllTeacher:input_type -> teacher.QueryAllTeacherRequest
	4,  // 10: teacher.TeacherService.QueryTeacher:input_type -> teacher.QueryTeacherRequest
	11, // 11: teacher.TeacherService.QueryAvailableTeachers:input_type -> teacher.QueryAvailableTeachersRequest
	0,  // 12: teacher.TeacherService.AddTeacher:output_type -> teacher.EmptyResponse
	0,  // 13: teacher.TeacherService.RemoveTeacher:output_type -> teacher.EmptyResponse
	0,  // 14: teacher.TeacherService.ModifyTeacher:output_type -> teacher.EmptyResponse
	7,  // 15: teacher.TeacherService.QueryAllTeacher:output_type -> teacher.QueryTeachersResponse
	5,  // 16: teacher.TeacherService.QueryTeacher:output_type -> teacher.QueryTeacherResponse
	7,  // 17: teacher.TeacherService.QueryAvailableTeachers:output_type -> teacher.QueryTeachersResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_teacher_proto_init() }
func file_teacher_proto_init() {
	if File_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Defense); i {
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
		file_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
		file_teacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTeacherRequest); i {
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
		file_teacher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTeacherResponse); i {
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
		file_teacher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllTeacherRequest); i {
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
		file_teacher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTeachersResponse); i {
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
		file_teacher_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyTeacherRequest); i {
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
		file_teacher_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTeacherRequest); i {
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
		file_teacher_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTeacherRequest); i {
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
		file_teacher_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAvailableTeachersRequest); i {
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
			RawDescriptor: file_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teacher_proto_goTypes,
		DependencyIndexes: file_teacher_proto_depIdxs,
		MessageInfos:      file_teacher_proto_msgTypes,
	}.Build()
	File_teacher_proto = out.File
	file_teacher_proto_rawDesc = nil
	file_teacher_proto_goTypes = nil
	file_teacher_proto_depIdxs = nil
}
