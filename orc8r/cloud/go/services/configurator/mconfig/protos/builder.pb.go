//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: orc8r/cloud/go/services/configurator/mconfig/protos/builder.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	storage "magma/orc8r/cloud/go/services/configurator/storage"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// network containing the gateway
	Network *storage.Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// graph of entities associated with the gateway
	Graph *storage.EntityGraph `protobuf:"bytes,2,opt,name=graph,proto3" json:"graph,omitempty"`
	// gateway_id of the gateway
	GatewayId string `protobuf:"bytes,3,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescGZIP(), []int{0}
}

func (x *BuildRequest) GetNetwork() *storage.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *BuildRequest) GetGraph() *storage.EntityGraph {
	if x != nil {
		return x.Graph
	}
	return nil
}

func (x *BuildRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

type BuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// configs_by_key contains the partial mconfig from this mconfig builder
	// Each config value contains a proto which is
	//   - first serialized to an any.Any proto
	//   - then serialized to JSON
	//
	// TODO(#2310): remove the need to serialize to JSON by sending proto descriptors
	ConfigsByKey map[string][]byte `protobuf:"bytes,1,rep,name=configs_by_key,json=configsByKey,proto3" json:"configs_by_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescGZIP(), []int{1}
}

func (x *BuildResponse) GetConfigsByKey() map[string][]byte {
	if x != nil {
		return x.ConfigsByKey
	}
	return nil
}

var File_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto protoreflect.FileDescriptor

var file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDesc = []byte{
	0x0a, 0x41, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x3a, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x43, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x0d,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42, 0x79,
	0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x7c, 0x0a, 0x0e, 0x4d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x05, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x2e, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescOnce sync.Once
	file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescData = file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDesc
)

func file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescGZIP() []byte {
	file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescOnce.Do(func() {
		file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescData)
	})
	return file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDescData
}

var file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_goTypes = []interface{}{
	(*BuildRequest)(nil),        // 0: magma.orc8r.configurator.mconfig.BuildRequest
	(*BuildResponse)(nil),       // 1: magma.orc8r.configurator.mconfig.BuildResponse
	nil,                         // 2: magma.orc8r.configurator.mconfig.BuildResponse.ConfigsByKeyEntry
	(*storage.Network)(nil),     // 3: magma.orc8r.configurator.storage.Network
	(*storage.EntityGraph)(nil), // 4: magma.orc8r.configurator.storage.EntityGraph
}
var file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_depIdxs = []int32{
	3, // 0: magma.orc8r.configurator.mconfig.BuildRequest.network:type_name -> magma.orc8r.configurator.storage.Network
	4, // 1: magma.orc8r.configurator.mconfig.BuildRequest.graph:type_name -> magma.orc8r.configurator.storage.EntityGraph
	2, // 2: magma.orc8r.configurator.mconfig.BuildResponse.configs_by_key:type_name -> magma.orc8r.configurator.mconfig.BuildResponse.ConfigsByKeyEntry
	0, // 3: magma.orc8r.configurator.mconfig.MconfigBuilder.Build:input_type -> magma.orc8r.configurator.mconfig.BuildRequest
	1, // 4: magma.orc8r.configurator.mconfig.MconfigBuilder.Build:output_type -> magma.orc8r.configurator.mconfig.BuildResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_init() }
func file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_init() {
	if File_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRequest); i {
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
		file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResponse); i {
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
			RawDescriptor: file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_goTypes,
		DependencyIndexes: file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_depIdxs,
		MessageInfos:      file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_msgTypes,
	}.Build()
	File_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto = out.File
	file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_rawDesc = nil
	file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_goTypes = nil
	file_orc8r_cloud_go_services_configurator_mconfig_protos_builder_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MconfigBuilderClient is the client API for MconfigBuilder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MconfigBuilderClient interface {
	// Build returns a partial mconfig containing the gateway configs for which
	// this builder is responsible.
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
}

type mconfigBuilderClient struct {
	cc grpc.ClientConnInterface
}

func NewMconfigBuilderClient(cc grpc.ClientConnInterface) MconfigBuilderClient {
	return &mconfigBuilderClient{cc}
}

func (c *mconfigBuilderClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.mconfig.MconfigBuilder/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MconfigBuilderServer is the server API for MconfigBuilder service.
type MconfigBuilderServer interface {
	// Build returns a partial mconfig containing the gateway configs for which
	// this builder is responsible.
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
}

// UnimplementedMconfigBuilderServer can be embedded to have forward compatible implementations.
type UnimplementedMconfigBuilderServer struct {
}

func (*UnimplementedMconfigBuilderServer) Build(context.Context, *BuildRequest) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}

func RegisterMconfigBuilderServer(s *grpc.Server, srv MconfigBuilderServer) {
	s.RegisterService(&_MconfigBuilder_serviceDesc, srv)
}

func _MconfigBuilder_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MconfigBuilderServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.mconfig.MconfigBuilder/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MconfigBuilderServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MconfigBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.configurator.mconfig.MconfigBuilder",
	HandlerType: (*MconfigBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _MconfigBuilder_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/configurator/mconfig/protos/builder.proto",
}
