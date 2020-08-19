// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: proto/command.proto

package auctioncmd

import (
	proto "github.com/golang/protobuf/proto"
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

// ********************* Player profile commands *********************
type PlayerType int32

const (
	PlayerType_NONE                PlayerType = 0
	PlayerType_RIGHT_HANDED_BAT    PlayerType = 1
	PlayerType_LEFT_HANDED_BAT     PlayerType = 2
	PlayerType_RIGHT_ARM_BOWL      PlayerType = 3
	PlayerType_LEFT_ARM_BOWL       PlayerType = 4
	PlayerType_BATTING_ALL_ROUNDER PlayerType = 5
	PlayerType_BOWLING_ALL_ROUNDER PlayerType = 6
)

// Enum value maps for PlayerType.
var (
	PlayerType_name = map[int32]string{
		0: "NONE",
		1: "RIGHT_HANDED_BAT",
		2: "LEFT_HANDED_BAT",
		3: "RIGHT_ARM_BOWL",
		4: "LEFT_ARM_BOWL",
		5: "BATTING_ALL_ROUNDER",
		6: "BOWLING_ALL_ROUNDER",
	}
	PlayerType_value = map[string]int32{
		"NONE":                0,
		"RIGHT_HANDED_BAT":    1,
		"LEFT_HANDED_BAT":     2,
		"RIGHT_ARM_BOWL":      3,
		"LEFT_ARM_BOWL":       4,
		"BATTING_ALL_ROUNDER": 5,
		"BOWLING_ALL_ROUNDER": 6,
	}
)

func (x PlayerType) Enum() *PlayerType {
	p := new(PlayerType)
	*p = x
	return p
}

func (x PlayerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_command_proto_enumTypes[0].Descriptor()
}

func (PlayerType) Type() protoreflect.EnumType {
	return &file_proto_command_proto_enumTypes[0]
}

func (x PlayerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerType.Descriptor instead.
func (PlayerType) EnumDescriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{0}
}

type AuctionCommand_CmdType int32

const (
	AuctionCommand_GET_PLAYER_INFO    AuctionCommand_CmdType = 0
	AuctionCommand_UPDATE_PLAYER_INFO AuctionCommand_CmdType = 1
)

// Enum value maps for AuctionCommand_CmdType.
var (
	AuctionCommand_CmdType_name = map[int32]string{
		0: "GET_PLAYER_INFO",
		1: "UPDATE_PLAYER_INFO",
	}
	AuctionCommand_CmdType_value = map[string]int32{
		"GET_PLAYER_INFO":    0,
		"UPDATE_PLAYER_INFO": 1,
	}
)

func (x AuctionCommand_CmdType) Enum() *AuctionCommand_CmdType {
	p := new(AuctionCommand_CmdType)
	*p = x
	return p
}

func (x AuctionCommand_CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuctionCommand_CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_command_proto_enumTypes[1].Descriptor()
}

func (AuctionCommand_CmdType) Type() protoreflect.EnumType {
	return &file_proto_command_proto_enumTypes[1]
}

func (x AuctionCommand_CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuctionCommand_CmdType.Descriptor instead.
func (AuctionCommand_CmdType) EnumDescriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{0, 0}
}

type AuctionResponse_ResponseType int32

const (
	AuctionResponse_GET_PLAYER_INFO    AuctionResponse_ResponseType = 0
	AuctionResponse_UPDATE_PLAYER_INFO AuctionResponse_ResponseType = 1
)

// Enum value maps for AuctionResponse_ResponseType.
var (
	AuctionResponse_ResponseType_name = map[int32]string{
		0: "GET_PLAYER_INFO",
		1: "UPDATE_PLAYER_INFO",
	}
	AuctionResponse_ResponseType_value = map[string]int32{
		"GET_PLAYER_INFO":    0,
		"UPDATE_PLAYER_INFO": 1,
	}
)

func (x AuctionResponse_ResponseType) Enum() *AuctionResponse_ResponseType {
	p := new(AuctionResponse_ResponseType)
	*p = x
	return p
}

func (x AuctionResponse_ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuctionResponse_ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_command_proto_enumTypes[2].Descriptor()
}

func (AuctionResponse_ResponseType) Type() protoreflect.EnumType {
	return &file_proto_command_proto_enumTypes[2]
}

func (x AuctionResponse_ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuctionResponse_ResponseType.Descriptor instead.
func (AuctionResponse_ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{1, 0}
}

type AuctionCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdType AuctionCommand_CmdType `protobuf:"varint,1,opt,name=cmd_type,json=cmdType,proto3,enum=auctioncmd.AuctionCommand_CmdType" json:"cmd_type,omitempty"`
	// Types that are assignable to Command:
	//	*AuctionCommand_GetPlayerInfoCmd
	//	*AuctionCommand_UpdatePlayerInfoCmd
	Command isAuctionCommand_Command `protobuf_oneof:"command"`
}

func (x *AuctionCommand) Reset() {
	*x = AuctionCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionCommand) ProtoMessage() {}

func (x *AuctionCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionCommand.ProtoReflect.Descriptor instead.
func (*AuctionCommand) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{0}
}

func (x *AuctionCommand) GetCmdType() AuctionCommand_CmdType {
	if x != nil {
		return x.CmdType
	}
	return AuctionCommand_GET_PLAYER_INFO
}

func (m *AuctionCommand) GetCommand() isAuctionCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *AuctionCommand) GetGetPlayerInfoCmd() *GetPlayerInfoCommand {
	if x, ok := x.GetCommand().(*AuctionCommand_GetPlayerInfoCmd); ok {
		return x.GetPlayerInfoCmd
	}
	return nil
}

func (x *AuctionCommand) GetUpdatePlayerInfoCmd() *UpdatePlayerInfoCommand {
	if x, ok := x.GetCommand().(*AuctionCommand_UpdatePlayerInfoCmd); ok {
		return x.UpdatePlayerInfoCmd
	}
	return nil
}

type isAuctionCommand_Command interface {
	isAuctionCommand_Command()
}

type AuctionCommand_GetPlayerInfoCmd struct {
	GetPlayerInfoCmd *GetPlayerInfoCommand `protobuf:"bytes,2,opt,name=get_player_info_cmd,json=getPlayerInfoCmd,proto3,oneof"`
}

type AuctionCommand_UpdatePlayerInfoCmd struct {
	UpdatePlayerInfoCmd *UpdatePlayerInfoCommand `protobuf:"bytes,3,opt,name=update_player_info_cmd,json=updatePlayerInfoCmd,proto3,oneof"`
}

func (*AuctionCommand_GetPlayerInfoCmd) isAuctionCommand_Command() {}

func (*AuctionCommand_UpdatePlayerInfoCmd) isAuctionCommand_Command() {}

type AuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errormsg     string                       `protobuf:"bytes,1,opt,name=errormsg,proto3" json:"errormsg,omitempty"`
	ResponseType AuctionResponse_ResponseType `protobuf:"varint,2,opt,name=response_type,json=responseType,proto3,enum=auctioncmd.AuctionResponse_ResponseType" json:"response_type,omitempty"`
	// Types that are assignable to Response:
	//	*AuctionResponse_GetPlayerInfoResponse
	//	*AuctionResponse_UpdatePlayerInfoResponse
	Response isAuctionResponse_Response `protobuf_oneof:"response"`
}

func (x *AuctionResponse) Reset() {
	*x = AuctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionResponse) ProtoMessage() {}

func (x *AuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionResponse.ProtoReflect.Descriptor instead.
func (*AuctionResponse) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{1}
}

func (x *AuctionResponse) GetErrormsg() string {
	if x != nil {
		return x.Errormsg
	}
	return ""
}

func (x *AuctionResponse) GetResponseType() AuctionResponse_ResponseType {
	if x != nil {
		return x.ResponseType
	}
	return AuctionResponse_GET_PLAYER_INFO
}

func (m *AuctionResponse) GetResponse() isAuctionResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *AuctionResponse) GetGetPlayerInfoResponse() *GetPlayerInfoResponse {
	if x, ok := x.GetResponse().(*AuctionResponse_GetPlayerInfoResponse); ok {
		return x.GetPlayerInfoResponse
	}
	return nil
}

func (x *AuctionResponse) GetUpdatePlayerInfoResponse() *UpdatePlayerInfoResponse {
	if x, ok := x.GetResponse().(*AuctionResponse_UpdatePlayerInfoResponse); ok {
		return x.UpdatePlayerInfoResponse
	}
	return nil
}

type isAuctionResponse_Response interface {
	isAuctionResponse_Response()
}

type AuctionResponse_GetPlayerInfoResponse struct {
	GetPlayerInfoResponse *GetPlayerInfoResponse `protobuf:"bytes,3,opt,name=get_player_info_response,json=getPlayerInfoResponse,proto3,oneof"`
}

type AuctionResponse_UpdatePlayerInfoResponse struct {
	UpdatePlayerInfoResponse *UpdatePlayerInfoResponse `protobuf:"bytes,4,opt,name=update_player_info_response,json=updatePlayerInfoResponse,proto3,oneof"`
}

func (*AuctionResponse_GetPlayerInfoResponse) isAuctionResponse_Response() {}

func (*AuctionResponse_UpdatePlayerInfoResponse) isAuctionResponse_Response() {}

type GetPlayerInfoCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *GetPlayerInfoCommand) Reset() {
	*x = GetPlayerInfoCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfoCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfoCommand) ProtoMessage() {}

func (x *GetPlayerInfoCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerInfoCommand.ProtoReflect.Descriptor instead.
func (*GetPlayerInfoCommand) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlayerInfoCommand) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type GetPlayerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName        string     `protobuf:"bytes,1,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerBio         string     `protobuf:"bytes,2,opt,name=player_bio,json=playerBio,proto3" json:"player_bio,omitempty"`
	PlayerProfileLink string     `protobuf:"bytes,3,opt,name=player_profile_link,json=playerProfileLink,proto3" json:"player_profile_link,omitempty"`
	PlayerType        PlayerType `protobuf:"varint,4,opt,name=player_type,json=playerType,proto3,enum=auctioncmd.PlayerType" json:"player_type,omitempty"`
	PlayerPicture     []byte     `protobuf:"bytes,5,opt,name=player_picture,json=playerPicture,proto3" json:"player_picture,omitempty"`
	IsPlayerActive    bool       `protobuf:"varint,6,opt,name=is_player_active,json=isPlayerActive,proto3" json:"is_player_active,omitempty"`
}

func (x *GetPlayerInfoResponse) Reset() {
	*x = GetPlayerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfoResponse) ProtoMessage() {}

func (x *GetPlayerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerInfoResponse) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *GetPlayerInfoResponse) GetPlayerBio() string {
	if x != nil {
		return x.PlayerBio
	}
	return ""
}

func (x *GetPlayerInfoResponse) GetPlayerProfileLink() string {
	if x != nil {
		return x.PlayerProfileLink
	}
	return ""
}

func (x *GetPlayerInfoResponse) GetPlayerType() PlayerType {
	if x != nil {
		return x.PlayerType
	}
	return PlayerType_NONE
}

func (x *GetPlayerInfoResponse) GetPlayerPicture() []byte {
	if x != nil {
		return x.PlayerPicture
	}
	return nil
}

func (x *GetPlayerInfoResponse) GetIsPlayerActive() bool {
	if x != nil {
		return x.IsPlayerActive
	}
	return false
}

type UpdatePlayerInfoCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid          string     `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	PlayerName        string     `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerBio         string     `protobuf:"bytes,3,opt,name=player_bio,json=playerBio,proto3" json:"player_bio,omitempty"`
	PlayerProfileLink string     `protobuf:"bytes,4,opt,name=player_profile_link,json=playerProfileLink,proto3" json:"player_profile_link,omitempty"`
	PlayerType        PlayerType `protobuf:"varint,5,opt,name=player_type,json=playerType,proto3,enum=auctioncmd.PlayerType" json:"player_type,omitempty"`
	PlayerPicture     []byte     `protobuf:"bytes,6,opt,name=player_picture,json=playerPicture,proto3" json:"player_picture,omitempty"`
	IsPlayerActive    bool       `protobuf:"varint,7,opt,name=is_player_active,json=isPlayerActive,proto3" json:"is_player_active,omitempty"`
}

func (x *UpdatePlayerInfoCommand) Reset() {
	*x = UpdatePlayerInfoCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerInfoCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerInfoCommand) ProtoMessage() {}

func (x *UpdatePlayerInfoCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerInfoCommand.ProtoReflect.Descriptor instead.
func (*UpdatePlayerInfoCommand) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePlayerInfoCommand) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UpdatePlayerInfoCommand) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *UpdatePlayerInfoCommand) GetPlayerBio() string {
	if x != nil {
		return x.PlayerBio
	}
	return ""
}

func (x *UpdatePlayerInfoCommand) GetPlayerProfileLink() string {
	if x != nil {
		return x.PlayerProfileLink
	}
	return ""
}

func (x *UpdatePlayerInfoCommand) GetPlayerType() PlayerType {
	if x != nil {
		return x.PlayerType
	}
	return PlayerType_NONE
}

func (x *UpdatePlayerInfoCommand) GetPlayerPicture() []byte {
	if x != nil {
		return x.PlayerPicture
	}
	return nil
}

func (x *UpdatePlayerInfoCommand) GetIsPlayerActive() bool {
	if x != nil {
		return x.IsPlayerActive
	}
	return false
}

type UpdatePlayerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateSuccess bool `protobuf:"varint,1,opt,name=update_success,json=updateSuccess,proto3" json:"update_success,omitempty"`
}

func (x *UpdatePlayerInfoResponse) Reset() {
	*x = UpdatePlayerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerInfoResponse) ProtoMessage() {}

func (x *UpdatePlayerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdatePlayerInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePlayerInfoResponse) GetUpdateSuccess() bool {
	if x != nil {
		return x.UpdateSuccess
	}
	return false
}

var File_proto_command_proto protoreflect.FileDescriptor

var file_proto_command_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6d,
	0x64, 0x22, 0xc1, 0x02, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6d, 0x64, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x51, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6d, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x10, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6d, 0x64, 0x12, 0x5a, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x6d, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6d, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x13, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43,
	0x6d, 0x64, 0x22, 0x36, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4c, 0x41,
	0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x8a, 0x03, 0x0a, 0x0f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x4d, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6d, 0x64, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x5c, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6d, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x15, 0x67, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x65, 0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x63, 0x6d, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x18, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x45, 0x54,
	0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x16,
	0x0a, 0x12, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x69,
	0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6d, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0xb0, 0x02, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x42, 0x69, 0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x63, 0x6d, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x69, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x41,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2a, 0x9a, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x54, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x45, 0x46, 0x54, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x5f,
	0x42, 0x41, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x41,
	0x52, 0x4d, 0x5f, 0x42, 0x4f, 0x57, 0x4c, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x46,
	0x54, 0x5f, 0x41, 0x52, 0x4d, 0x5f, 0x42, 0x4f, 0x57, 0x4c, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13,
	0x42, 0x41, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x4f, 0x55, 0x4e,
	0x44, 0x45, 0x52, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x4f, 0x57, 0x4c, 0x49, 0x4e, 0x47,
	0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x06, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_command_proto_rawDescOnce sync.Once
	file_proto_command_proto_rawDescData = file_proto_command_proto_rawDesc
)

func file_proto_command_proto_rawDescGZIP() []byte {
	file_proto_command_proto_rawDescOnce.Do(func() {
		file_proto_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_command_proto_rawDescData)
	})
	return file_proto_command_proto_rawDescData
}

var file_proto_command_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_command_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_command_proto_goTypes = []interface{}{
	(PlayerType)(0),                   // 0: auctioncmd.PlayerType
	(AuctionCommand_CmdType)(0),       // 1: auctioncmd.AuctionCommand.CmdType
	(AuctionResponse_ResponseType)(0), // 2: auctioncmd.AuctionResponse.ResponseType
	(*AuctionCommand)(nil),            // 3: auctioncmd.AuctionCommand
	(*AuctionResponse)(nil),           // 4: auctioncmd.AuctionResponse
	(*GetPlayerInfoCommand)(nil),      // 5: auctioncmd.GetPlayerInfoCommand
	(*GetPlayerInfoResponse)(nil),     // 6: auctioncmd.GetPlayerInfoResponse
	(*UpdatePlayerInfoCommand)(nil),   // 7: auctioncmd.UpdatePlayerInfoCommand
	(*UpdatePlayerInfoResponse)(nil),  // 8: auctioncmd.UpdatePlayerInfoResponse
}
var file_proto_command_proto_depIdxs = []int32{
	1, // 0: auctioncmd.AuctionCommand.cmd_type:type_name -> auctioncmd.AuctionCommand.CmdType
	5, // 1: auctioncmd.AuctionCommand.get_player_info_cmd:type_name -> auctioncmd.GetPlayerInfoCommand
	7, // 2: auctioncmd.AuctionCommand.update_player_info_cmd:type_name -> auctioncmd.UpdatePlayerInfoCommand
	2, // 3: auctioncmd.AuctionResponse.response_type:type_name -> auctioncmd.AuctionResponse.ResponseType
	6, // 4: auctioncmd.AuctionResponse.get_player_info_response:type_name -> auctioncmd.GetPlayerInfoResponse
	8, // 5: auctioncmd.AuctionResponse.update_player_info_response:type_name -> auctioncmd.UpdatePlayerInfoResponse
	0, // 6: auctioncmd.GetPlayerInfoResponse.player_type:type_name -> auctioncmd.PlayerType
	0, // 7: auctioncmd.UpdatePlayerInfoCommand.player_type:type_name -> auctioncmd.PlayerType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_command_proto_init() }
func file_proto_command_proto_init() {
	if File_proto_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionCommand); i {
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
		file_proto_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionResponse); i {
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
		file_proto_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerInfoCommand); i {
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
		file_proto_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerInfoResponse); i {
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
		file_proto_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerInfoCommand); i {
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
		file_proto_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerInfoResponse); i {
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
	file_proto_command_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AuctionCommand_GetPlayerInfoCmd)(nil),
		(*AuctionCommand_UpdatePlayerInfoCmd)(nil),
	}
	file_proto_command_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AuctionResponse_GetPlayerInfoResponse)(nil),
		(*AuctionResponse_UpdatePlayerInfoResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_command_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_command_proto_goTypes,
		DependencyIndexes: file_proto_command_proto_depIdxs,
		EnumInfos:         file_proto_command_proto_enumTypes,
		MessageInfos:      file_proto_command_proto_msgTypes,
	}.Build()
	File_proto_command_proto = out.File
	file_proto_command_proto_rawDesc = nil
	file_proto_command_proto_goTypes = nil
	file_proto_command_proto_depIdxs = nil
}
