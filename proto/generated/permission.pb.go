// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: permission.proto

package generated

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

type CreatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId       uint32  `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceParentId *uint32 `protobuf:"varint,3,opt,name=resource_parent_id,json=resourceParentId,proto3,oneof" json:"resource_parent_id,omitempty"`
	Relation         uint32  `protobuf:"varint,4,opt,name=relation,proto3" json:"relation,omitempty"`
	UserId           uint32  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId          *uint32 `protobuf:"varint,6,opt,name=group_id,json=groupId,proto3,oneof" json:"group_id,omitempty"`
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionRequest.ProtoReflect.Descriptor instead.
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePermissionRequest) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *CreatePermissionRequest) GetResourceParentId() uint32 {
	if x != nil && x.ResourceParentId != nil {
		return *x.ResourceParentId
	}
	return 0
}

func (x *CreatePermissionRequest) GetRelation() uint32 {
	if x != nil {
		return x.Relation
	}
	return 0
}

func (x *CreatePermissionRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreatePermissionRequest) GetGroupId() uint32 {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return 0
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResourceId       uint32 `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceParentId uint32 `protobuf:"varint,3,opt,name=resource_parent_id,json=resourceParentId,proto3" json:"resource_parent_id,omitempty"`
	Relation         uint32 `protobuf:"varint,4,opt,name=relation,proto3" json:"relation,omitempty"`
	UserId           uint32 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId          uint32 `protobuf:"varint,6,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	IsDeleted        bool   `protobuf:"varint,7,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Permission) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *Permission) GetResourceParentId() uint32 {
	if x != nil {
		return x.ResourceParentId
	}
	return 0
}

func (x *Permission) GetRelation() uint32 {
	if x != nil {
		return x.Relation
	}
	return 0
}

func (x *Permission) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Permission) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Permission) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type CheckPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *CheckPermissionResponse) Reset() {
	*x = CheckPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionResponse) ProtoMessage() {}

func (x *CheckPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckPermissionResponse) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{2}
}

func (x *CheckPermissionResponse) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

var File_permission_proto protoreflect.FileDescriptor

var file_permission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x22, 0xda, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x2b, 0x0a,
	0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x99, 0x02, 0x0a, 0x11, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x59, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x64, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_proto_rawDescOnce sync.Once
	file_permission_proto_rawDescData = file_permission_proto_rawDesc
)

func file_permission_proto_rawDescGZIP() []byte {
	file_permission_proto_rawDescOnce.Do(func() {
		file_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_proto_rawDescData)
	})
	return file_permission_proto_rawDescData
}

var file_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_permission_proto_goTypes = []interface{}{
	(*CreatePermissionRequest)(nil), // 0: CreatePermissionRequest
	(*Permission)(nil),              // 1: Permission
	(*CheckPermissionResponse)(nil), // 2: CheckPermissionResponse
	(*ID)(nil),                      // 3: ID
	(*Message)(nil),                 // 4: Message
}
var file_permission_proto_depIdxs = []int32{
	3, // 0: PermissionService.DeletePermission:input_type -> ID
	0, // 1: PermissionService.CreatePermission:input_type -> CreatePermissionRequest
	0, // 2: PermissionService.CheckPermission:input_type -> CreatePermissionRequest
	4, // 3: PermissionService.DeletePermission:output_type -> Message
	1, // 4: PermissionService.CreatePermission:output_type -> Permission
	2, // 5: PermissionService.CheckPermission:output_type -> CheckPermissionResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_permission_proto_init() }
func file_permission_proto_init() {
	if File_permission_proto != nil {
		return
	}
	file_pagination_proto_init()
	file_annotations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePermissionRequest); i {
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
		file_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionResponse); i {
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
	file_permission_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_permission_proto_goTypes,
		DependencyIndexes: file_permission_proto_depIdxs,
		MessageInfos:      file_permission_proto_msgTypes,
	}.Build()
	File_permission_proto = out.File
	file_permission_proto_rawDesc = nil
	file_permission_proto_goTypes = nil
	file_permission_proto_depIdxs = nil
}
