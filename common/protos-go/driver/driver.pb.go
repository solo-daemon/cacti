// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: driver/driver.proto

package driver

import (
	common "github.com/hyperledger-labs/weaver-dlt-interoperability/common/protos-go/common"
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

// Data for a View processing by dest-driver
type WriteExternalStateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewPayload *common.ViewPayload         `protobuf:"bytes,1,opt,name=view_payload,json=viewPayload,proto3" json:"view_payload,omitempty"`
	Ctx         *common.ContractTransaction `protobuf:"bytes,2,opt,name=ctx,proto3" json:"ctx,omitempty"`
}

func (x *WriteExternalStateMessage) Reset() {
	*x = WriteExternalStateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteExternalStateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteExternalStateMessage) ProtoMessage() {}

func (x *WriteExternalStateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_driver_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteExternalStateMessage.ProtoReflect.Descriptor instead.
func (*WriteExternalStateMessage) Descriptor() ([]byte, []int) {
	return file_driver_driver_proto_rawDescGZIP(), []int{0}
}

func (x *WriteExternalStateMessage) GetViewPayload() *common.ViewPayload {
	if x != nil {
		return x.ViewPayload
	}
	return nil
}

func (x *WriteExternalStateMessage) GetCtx() *common.ContractTransaction {
	if x != nil {
		return x.Ctx
	}
	return nil
}

var File_driver_driver_proto protoreflect.FileDescriptor

var file_driver_driver_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x19, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x0b, 0x76, 0x69, 0x65, 0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x34, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x03, 0x63, 0x74, 0x78, 0x32, 0xcd, 0x02, 0x0a, 0x13, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a,
	0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x23, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x72, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x61,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2d, 0x64, 0x6c, 0x74,
	0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d,
	0x67, 0x6f, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_driver_driver_proto_rawDescOnce sync.Once
	file_driver_driver_proto_rawDescData = file_driver_driver_proto_rawDesc
)

func file_driver_driver_proto_rawDescGZIP() []byte {
	file_driver_driver_proto_rawDescOnce.Do(func() {
		file_driver_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_driver_proto_rawDescData)
	})
	return file_driver_driver_proto_rawDescData
}

var file_driver_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_driver_driver_proto_goTypes = []interface{}{
	(*WriteExternalStateMessage)(nil),  // 0: driver.driver.WriteExternalStateMessage
	(*common.ViewPayload)(nil),         // 1: common.state.ViewPayload
	(*common.ContractTransaction)(nil), // 2: common.events.ContractTransaction
	(*common.Query)(nil),               // 3: common.query.Query
	(*common.EventSubscription)(nil),   // 4: common.events.EventSubscription
	(*common.Ack)(nil),                 // 5: common.ack.Ack
}
var file_driver_driver_proto_depIdxs = []int32{
	1, // 0: driver.driver.WriteExternalStateMessage.view_payload:type_name -> common.state.ViewPayload
	2, // 1: driver.driver.WriteExternalStateMessage.ctx:type_name -> common.events.ContractTransaction
	3, // 2: driver.driver.DriverCommunication.RequestDriverState:input_type -> common.query.Query
	4, // 3: driver.driver.DriverCommunication.SubscribeEvent:input_type -> common.events.EventSubscription
	4, // 4: driver.driver.DriverCommunication.RequestSignedEventSubscriptionQuery:input_type -> common.events.EventSubscription
	0, // 5: driver.driver.DriverCommunication.WriteExternalState:input_type -> driver.driver.WriteExternalStateMessage
	5, // 6: driver.driver.DriverCommunication.RequestDriverState:output_type -> common.ack.Ack
	5, // 7: driver.driver.DriverCommunication.SubscribeEvent:output_type -> common.ack.Ack
	3, // 8: driver.driver.DriverCommunication.RequestSignedEventSubscriptionQuery:output_type -> common.query.Query
	5, // 9: driver.driver.DriverCommunication.WriteExternalState:output_type -> common.ack.Ack
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_driver_driver_proto_init() }
func file_driver_driver_proto_init() {
	if File_driver_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driver_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteExternalStateMessage); i {
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
			RawDescriptor: file_driver_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_driver_proto_goTypes,
		DependencyIndexes: file_driver_driver_proto_depIdxs,
		MessageInfos:      file_driver_driver_proto_msgTypes,
	}.Build()
	File_driver_driver_proto = out.File
	file_driver_driver_proto_rawDesc = nil
	file_driver_driver_proto_goTypes = nil
	file_driver_driver_proto_depIdxs = nil
}
