// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: messageService.proto

package service

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

type DouyinMessageChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" form:"token,omitempty"`                                // 用户鉴权token
	FromUserId int64  `protobuf:"varint,2,opt,name=from_user_id,json=fromUserId,proto3" form:"from_user_id,omitempty"` //
	ToUserId   int64  `protobuf:"varint,3,opt,name=to_user_id,json=toUserId,proto3" form:"to_user_id,omitempty"`       // 对方用户id
	PreMsgTime int64  `protobuf:"varint,4,opt,name=pre_msg_time,json=preMsgTime,proto3" form:"pre_msg_time,omitempty"` // 上次最新消息的时间
}

func (x *DouyinMessageChatRequest) Reset() {
	*x = DouyinMessageChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinMessageChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinMessageChatRequest) ProtoMessage() {}

func (x *DouyinMessageChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messageService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinMessageChatRequest.ProtoReflect.Descriptor instead.
func (*DouyinMessageChatRequest) Descriptor() ([]byte, []int) {
	return file_messageService_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinMessageChatRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinMessageChatRequest) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *DouyinMessageChatRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *DouyinMessageChatRequest) GetPreMsgTime() int64 {
	if x != nil {
		return x.PreMsgTime
	}
	return 0
}

type DouyinMessageChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code,omitempty"`   // 状态码, 0-成功, 其他值-失败
	StatusMsg   string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg,omitempty"`       // 返回状态描述
	MessageList []*Message `protobuf:"bytes,3,rep,name=message_list,json=messageList,proto3" form:"message_list,omitempty"` // 消息列表
}

func (x *DouyinMessageChatResponse) Reset() {
	*x = DouyinMessageChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinMessageChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinMessageChatResponse) ProtoMessage() {}

func (x *DouyinMessageChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messageService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinMessageChatResponse.ProtoReflect.Descriptor instead.
func (*DouyinMessageChatResponse) Descriptor() ([]byte, []int) {
	return file_messageService_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinMessageChatResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinMessageChatResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinMessageChatResponse) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

type DouyinMessageActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" form:"token,omitempty"`                                // 用户鉴权token
	FromUserId int64  `protobuf:"varint,2,opt,name=from_user_id,json=fromUserId,proto3" form:"from_user_id,omitempty"` //
	ToUserId   int64  `protobuf:"varint,3,opt,name=to_user_id,json=toUserId,proto3" form:"to_user_id,omitempty"`       // 对方用户id
	ActionType int32  `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" form:"action_type,omitempty"`   // 1-发送消息
	Content    string `protobuf:"bytes,5,opt,name=content,proto3" form:"content,omitempty"`                            // 消息内容
}

func (x *DouyinMessageActionRequest) Reset() {
	*x = DouyinMessageActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinMessageActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinMessageActionRequest) ProtoMessage() {}

func (x *DouyinMessageActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messageService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinMessageActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinMessageActionRequest) Descriptor() ([]byte, []int) {
	return file_messageService_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinMessageActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinMessageActionRequest) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *DouyinMessageActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *DouyinMessageActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *DouyinMessageActionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DouyinMessageActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code,omitempty"` // 状态码, 0-成功, 其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg,omitempty"`     // 返回状态描述
}

func (x *DouyinMessageActionResponse) Reset() {
	*x = DouyinMessageActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinMessageActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinMessageActionResponse) ProtoMessage() {}

func (x *DouyinMessageActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messageService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinMessageActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinMessageActionResponse) Descriptor() ([]byte, []int) {
	return file_messageService_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinMessageActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinMessageActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_messageService_proto protoreflect.FileDescriptor

var file_messageService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x95, 0x01, 0x0a, 0x1b, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x5f, 0x6d, 0x73, 0x67,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x65,
	0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x1c, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x1d, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x1e, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x32, 0xba, 0x01,
	0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x56, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messageService_proto_rawDescOnce sync.Once
	file_messageService_proto_rawDescData = file_messageService_proto_rawDesc
)

func file_messageService_proto_rawDescGZIP() []byte {
	file_messageService_proto_rawDescOnce.Do(func() {
		file_messageService_proto_rawDescData = protoimpl.X.CompressGZIP(file_messageService_proto_rawDescData)
	})
	return file_messageService_proto_rawDescData
}

var file_messageService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messageService_proto_goTypes = []interface{}{
	(*DouyinMessageChatRequest)(nil),    // 0: pb.douyin_message_chat_request
	(*DouyinMessageChatResponse)(nil),   // 1: pb.douyin_message_chat_response
	(*DouyinMessageActionRequest)(nil),  // 2: pb.douyin_message_action_request
	(*DouyinMessageActionResponse)(nil), // 3: pb.douyin_message_action_response
	(*Message)(nil),                     // 4: pb.Message
}
var file_messageService_proto_depIdxs = []int32{
	4, // 0: pb.douyin_message_chat_response.message_list:type_name -> pb.Message
	2, // 1: pb.MessageService.MessageAction:input_type -> pb.douyin_message_action_request
	0, // 2: pb.MessageService.MessageChat:input_type -> pb.douyin_message_chat_request
	3, // 3: pb.MessageService.MessageAction:output_type -> pb.douyin_message_action_response
	1, // 4: pb.MessageService.MessageChat:output_type -> pb.douyin_message_chat_response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messageService_proto_init() }
func file_messageService_proto_init() {
	if File_messageService_proto != nil {
		return
	}
	file_messageModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messageService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinMessageChatRequest); i {
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
		file_messageService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinMessageChatResponse); i {
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
		file_messageService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinMessageActionRequest); i {
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
		file_messageService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinMessageActionResponse); i {
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
			RawDescriptor: file_messageService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messageService_proto_goTypes,
		DependencyIndexes: file_messageService_proto_depIdxs,
		MessageInfos:      file_messageService_proto_msgTypes,
	}.Build()
	File_messageService_proto = out.File
	file_messageService_proto_rawDesc = nil
	file_messageService_proto_goTypes = nil
	file_messageService_proto_depIdxs = nil
}