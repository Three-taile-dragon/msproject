//grpc 登陆服务配置文件

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: task_service.proto

package task

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

type TaskReqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId    int64  `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	ProjectCode string `protobuf:"bytes,2,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	Page        int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	StageCode   string `protobuf:"bytes,5,opt,name=stageCode,proto3" json:"stageCode,omitempty"`
	Name        string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	AssignTo    string `protobuf:"bytes,7,opt,name=assignTo,proto3" json:"assignTo,omitempty"`
}

func (x *TaskReqMessage) Reset() {
	*x = TaskReqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReqMessage) ProtoMessage() {}

func (x *TaskReqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReqMessage.ProtoReflect.Descriptor instead.
func (*TaskReqMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{0}
}

func (x *TaskReqMessage) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *TaskReqMessage) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *TaskReqMessage) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TaskReqMessage) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TaskReqMessage) GetStageCode() string {
	if x != nil {
		return x.StageCode
	}
	return ""
}

func (x *TaskReqMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskReqMessage) GetAssignTo() string {
	if x != nil {
		return x.AssignTo
	}
	return ""
}

type TaskStagesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectCode string `protobuf:"bytes,3,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	Sort        int32  `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime  string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Deleted     int32  `protobuf:"varint,7,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Id          int32  `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskStagesMessage) Reset() {
	*x = TaskStagesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStagesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStagesMessage) ProtoMessage() {}

func (x *TaskStagesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStagesMessage.ProtoReflect.Descriptor instead.
func (*TaskStagesMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{1}
}

func (x *TaskStagesMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TaskStagesMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskStagesMessage) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *TaskStagesMessage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *TaskStagesMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskStagesMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *TaskStagesMessage) GetDeleted() int32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *TaskStagesMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskStagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*TaskStagesMessage `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TaskStagesResponse) Reset() {
	*x = TaskStagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStagesResponse) ProtoMessage() {}

func (x *TaskStagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStagesResponse.ProtoReflect.Descriptor instead.
func (*TaskStagesResponse) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{2}
}

func (x *TaskStagesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskStagesResponse) GetList() []*TaskStagesMessage {
	if x != nil {
		return x.List
	}
	return nil
}

type MemberProjectMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar     string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	MemberCode int64  `protobuf:"varint,3,opt,name=memberCode,proto3" json:"memberCode,omitempty"`
	Code       string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	IsOwner    int32  `protobuf:"varint,6,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
}

func (x *MemberProjectMessage) Reset() {
	*x = MemberProjectMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberProjectMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberProjectMessage) ProtoMessage() {}

func (x *MemberProjectMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberProjectMessage.ProtoReflect.Descriptor instead.
func (*MemberProjectMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{3}
}

func (x *MemberProjectMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberProjectMessage) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *MemberProjectMessage) GetMemberCode() int64 {
	if x != nil {
		return x.MemberCode
	}
	return 0
}

func (x *MemberProjectMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MemberProjectMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberProjectMessage) GetIsOwner() int32 {
	if x != nil {
		return x.IsOwner
	}
	return 0
}

type MemberProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*MemberProjectMessage `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MemberProjectResponse) Reset() {
	*x = MemberProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberProjectResponse) ProtoMessage() {}

func (x *MemberProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberProjectResponse.ProtoReflect.Descriptor instead.
func (*MemberProjectResponse) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{4}
}

func (x *MemberProjectResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemberProjectResponse) GetList() []*MemberProjectMessage {
	if x != nil {
		return x.List
	}
	return nil
}

type TaskMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64            `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ProjectCode   string           `protobuf:"bytes,2,opt,name=ProjectCode,proto3" json:"ProjectCode,omitempty"`
	Name          string           `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Pri           int32            `protobuf:"varint,4,opt,name=Pri,proto3" json:"Pri,omitempty"`
	ExecuteStatus string           `protobuf:"bytes,5,opt,name=ExecuteStatus,proto3" json:"ExecuteStatus,omitempty"`
	Description   string           `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	CreateBy      string           `protobuf:"bytes,7,opt,name=CreateBy,proto3" json:"CreateBy,omitempty"`
	DoneBy        string           `protobuf:"bytes,8,opt,name=DoneBy,proto3" json:"DoneBy,omitempty"`
	DoneTime      string           `protobuf:"bytes,9,opt,name=DoneTime,proto3" json:"DoneTime,omitempty"`
	CreateTime    string           `protobuf:"bytes,10,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	AssignTo      string           `protobuf:"bytes,11,opt,name=AssignTo,proto3" json:"AssignTo,omitempty"`
	Deleted       int32            `protobuf:"varint,12,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	StageCode     string           `protobuf:"bytes,13,opt,name=StageCode,proto3" json:"StageCode,omitempty"`
	TaskTag       string           `protobuf:"bytes,14,opt,name=TaskTag,proto3" json:"TaskTag,omitempty"`
	Done          int32            `protobuf:"varint,15,opt,name=Done,proto3" json:"Done,omitempty"`
	BeginTime     string           `protobuf:"bytes,16,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
	EndTime       string           `protobuf:"bytes,17,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	RemindTime    string           `protobuf:"bytes,18,opt,name=RemindTime,proto3" json:"RemindTime,omitempty"`
	Pcode         string           `protobuf:"bytes,19,opt,name=Pcode,proto3" json:"Pcode,omitempty"`
	Sort          int32            `protobuf:"varint,20,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Like          int32            `protobuf:"varint,21,opt,name=Like,proto3" json:"Like,omitempty"`
	Star          int32            `protobuf:"varint,22,opt,name=Star,proto3" json:"Star,omitempty"`
	DeletedTime   string           `protobuf:"bytes,23,opt,name=DeletedTime,proto3" json:"DeletedTime,omitempty"`
	Private       int32            `protobuf:"varint,24,opt,name=Private,proto3" json:"Private,omitempty"`
	IdNum         int32            `protobuf:"varint,25,opt,name=IdNum,proto3" json:"IdNum,omitempty"`
	Path          string           `protobuf:"bytes,26,opt,name=Path,proto3" json:"Path,omitempty"`
	Schedule      int32            `protobuf:"varint,27,opt,name=Schedule,proto3" json:"Schedule,omitempty"`
	VersionCode   string           `protobuf:"bytes,28,opt,name=VersionCode,proto3" json:"VersionCode,omitempty"`
	FeaturesCode  string           `protobuf:"bytes,29,opt,name=FeaturesCode,proto3" json:"FeaturesCode,omitempty"`
	WorkTime      int32            `protobuf:"varint,30,opt,name=WorkTime,proto3" json:"WorkTime,omitempty"`
	Status        int32            `protobuf:"varint,31,opt,name=Status,proto3" json:"Status,omitempty"`
	Code          string           `protobuf:"bytes,32,opt,name=code,proto3" json:"code,omitempty"`
	CanRead       int32            `protobuf:"varint,33,opt,name=canRead,proto3" json:"canRead,omitempty"`
	Executor      *ExecutorMessage `protobuf:"bytes,34,opt,name=executor,proto3" json:"executor,omitempty"`
}

func (x *TaskMessage) Reset() {
	*x = TaskMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskMessage) ProtoMessage() {}

func (x *TaskMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskMessage.ProtoReflect.Descriptor instead.
func (*TaskMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{5}
}

func (x *TaskMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskMessage) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *TaskMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskMessage) GetPri() int32 {
	if x != nil {
		return x.Pri
	}
	return 0
}

func (x *TaskMessage) GetExecuteStatus() string {
	if x != nil {
		return x.ExecuteStatus
	}
	return ""
}

func (x *TaskMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskMessage) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *TaskMessage) GetDoneBy() string {
	if x != nil {
		return x.DoneBy
	}
	return ""
}

func (x *TaskMessage) GetDoneTime() string {
	if x != nil {
		return x.DoneTime
	}
	return ""
}

func (x *TaskMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *TaskMessage) GetAssignTo() string {
	if x != nil {
		return x.AssignTo
	}
	return ""
}

func (x *TaskMessage) GetDeleted() int32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *TaskMessage) GetStageCode() string {
	if x != nil {
		return x.StageCode
	}
	return ""
}

func (x *TaskMessage) GetTaskTag() string {
	if x != nil {
		return x.TaskTag
	}
	return ""
}

func (x *TaskMessage) GetDone() int32 {
	if x != nil {
		return x.Done
	}
	return 0
}

func (x *TaskMessage) GetBeginTime() string {
	if x != nil {
		return x.BeginTime
	}
	return ""
}

func (x *TaskMessage) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *TaskMessage) GetRemindTime() string {
	if x != nil {
		return x.RemindTime
	}
	return ""
}

func (x *TaskMessage) GetPcode() string {
	if x != nil {
		return x.Pcode
	}
	return ""
}

func (x *TaskMessage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *TaskMessage) GetLike() int32 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *TaskMessage) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *TaskMessage) GetDeletedTime() string {
	if x != nil {
		return x.DeletedTime
	}
	return ""
}

func (x *TaskMessage) GetPrivate() int32 {
	if x != nil {
		return x.Private
	}
	return 0
}

func (x *TaskMessage) GetIdNum() int32 {
	if x != nil {
		return x.IdNum
	}
	return 0
}

func (x *TaskMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TaskMessage) GetSchedule() int32 {
	if x != nil {
		return x.Schedule
	}
	return 0
}

func (x *TaskMessage) GetVersionCode() string {
	if x != nil {
		return x.VersionCode
	}
	return ""
}

func (x *TaskMessage) GetFeaturesCode() string {
	if x != nil {
		return x.FeaturesCode
	}
	return ""
}

func (x *TaskMessage) GetWorkTime() int32 {
	if x != nil {
		return x.WorkTime
	}
	return 0
}

func (x *TaskMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TaskMessage) GetCanRead() int32 {
	if x != nil {
		return x.CanRead
	}
	return 0
}

func (x *TaskMessage) GetExecutor() *ExecutorMessage {
	if x != nil {
		return x.Executor
	}
	return nil
}

type ExecutorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Avatar string `protobuf:"bytes,2,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *ExecutorMessage) Reset() {
	*x = ExecutorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorMessage) ProtoMessage() {}

func (x *ExecutorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorMessage.ProtoReflect.Descriptor instead.
func (*ExecutorMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{6}
}

func (x *ExecutorMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecutorMessage) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*TaskMessage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{7}
}

func (x *TaskListResponse) GetList() []*TaskMessage {
	if x != nil {
		return x.List
	}
	return nil
}

var File_task_service_proto protoreflect.FileDescriptor

var file_task_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x54, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x54, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xa6, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x22, 0x68, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x39, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xb1, 0x07, 0x0a, 0x0b,
	0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x50, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6e, 0x65, 0x42,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x61, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53,
	0x6f, 0x72, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4c,
	0x69, 0x6b, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x53, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x61, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x3c, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x22,
	0x3d, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x44,
	0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x32, 0x95, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x26, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x08, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_task_service_proto_rawDescOnce sync.Once
	file_task_service_proto_rawDescData = file_task_service_proto_rawDesc
)

func file_task_service_proto_rawDescGZIP() []byte {
	file_task_service_proto_rawDescOnce.Do(func() {
		file_task_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_service_proto_rawDescData)
	})
	return file_task_service_proto_rawDescData
}

var file_task_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_task_service_proto_goTypes = []interface{}{
	(*TaskReqMessage)(nil),        // 0: task.service.v1.TaskReqMessage
	(*TaskStagesMessage)(nil),     // 1: task.service.v1.TaskStagesMessage
	(*TaskStagesResponse)(nil),    // 2: task.service.v1.TaskStagesResponse
	(*MemberProjectMessage)(nil),  // 3: task.service.v1.MemberProjectMessage
	(*MemberProjectResponse)(nil), // 4: task.service.v1.MemberProjectResponse
	(*TaskMessage)(nil),           // 5: task.service.v1.TaskMessage
	(*ExecutorMessage)(nil),       // 6: task.service.v1.ExecutorMessage
	(*TaskListResponse)(nil),      // 7: task.service.v1.TaskListResponse
}
var file_task_service_proto_depIdxs = []int32{
	1, // 0: task.service.v1.TaskStagesResponse.list:type_name -> task.service.v1.TaskStagesMessage
	3, // 1: task.service.v1.MemberProjectResponse.list:type_name -> task.service.v1.MemberProjectMessage
	6, // 2: task.service.v1.TaskMessage.executor:type_name -> task.service.v1.ExecutorMessage
	5, // 3: task.service.v1.TaskListResponse.list:type_name -> task.service.v1.TaskMessage
	0, // 4: task.service.v1.TaskService.TaskStages:input_type -> task.service.v1.TaskReqMessage
	0, // 5: task.service.v1.TaskService.MemberProjectList:input_type -> task.service.v1.TaskReqMessage
	0, // 6: task.service.v1.TaskService.TaskList:input_type -> task.service.v1.TaskReqMessage
	2, // 7: task.service.v1.TaskService.TaskStages:output_type -> task.service.v1.TaskStagesResponse
	4, // 8: task.service.v1.TaskService.MemberProjectList:output_type -> task.service.v1.MemberProjectResponse
	7, // 9: task.service.v1.TaskService.TaskList:output_type -> task.service.v1.TaskListResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_task_service_proto_init() }
func file_task_service_proto_init() {
	if File_task_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReqMessage); i {
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
		file_task_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStagesMessage); i {
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
		file_task_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStagesResponse); i {
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
		file_task_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberProjectMessage); i {
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
		file_task_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberProjectResponse); i {
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
		file_task_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskMessage); i {
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
		file_task_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorMessage); i {
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
		file_task_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse); i {
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
			RawDescriptor: file_task_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_service_proto_goTypes,
		DependencyIndexes: file_task_service_proto_depIdxs,
		MessageInfos:      file_task_service_proto_msgTypes,
	}.Build()
	File_task_service_proto = out.File
	file_task_service_proto_rawDesc = nil
	file_task_service_proto_goTypes = nil
	file_task_service_proto_depIdxs = nil
}
