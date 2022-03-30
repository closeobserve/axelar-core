// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nexus/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("nexus/v1beta1/service.proto", fileDescriptor_e8a22d972057ace6) }
func init() {
	golang_proto.RegisterFile("nexus/v1beta1/service.proto", fileDescriptor_e8a22d972057ace6)
}

var fileDescriptor_e8a22d972057ace6 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xc7, 0x99, 0xdf, 0x2f, 0x69, 0xcc, 0x28, 0x51, 0x27, 0x44, 0xd3, 0x52, 0x37, 0x76, 0x29,
	0x20, 0x85, 0x76, 0xa0, 0xc6, 0x0b, 0x37, 0x90, 0x10, 0x63, 0x24, 0x41, 0xf0, 0xe4, 0x65, 0x33,
	0x6d, 0x9f, 0x2e, 0x1b, 0x70, 0xa7, 0xec, 0x4c, 0xb1, 0x0d, 0x21, 0x31, 0xea, 0xcd, 0x8b, 0xff,
	0x2e, 0xbe, 0x05, 0x8f, 0xbe, 0x02, 0xbd, 0x71, 0x24, 0xf1, 0xe2, 0xd1, 0x50, 0x5f, 0x88, 0xd9,
	0xd9, 0x19, 0x64, 0xb7, 0xbb, 0x85, 0xd3, 0xb6, 0xf3, 0x7c, 0x9f, 0xf9, 0x7e, 0x9e, 0x67, 0x9f,
	0x99, 0xc5, 0x93, 0x3e, 0xf4, 0xba, 0x82, 0x1e, 0x2c, 0x35, 0x40, 0xb2, 0x25, 0x2a, 0x20, 0x38,
	0xf0, 0x9a, 0x50, 0xeb, 0x04, 0x5c, 0x72, 0x32, 0xae, 0x82, 0x35, 0x1d, 0x2c, 0x4c, 0xb8, 0xdc,
	0xe5, 0x2a, 0x42, 0xc3, 0x5f, 0x91, 0xa8, 0x50, 0x74, 0x39, 0x77, 0xf7, 0x80, 0xb2, 0x8e, 0x47,
	0x99, 0xef, 0x73, 0xc9, 0xa4, 0xc7, 0x7d, 0xa1, 0xa3, 0xb7, 0xe2, 0xfb, 0xcb, 0x9e, 0x5e, 0xcf,
	0xc7, 0xd7, 0xf7, 0xbb, 0x10, 0xf4, 0xa3, 0x50, 0xfd, 0x4b, 0x0e, 0xe3, 0x0d, 0xe1, 0x6e, 0x47,
	0x28, 0xe4, 0x2b, 0xc2, 0xb7, 0xb7, 0xc0, 0xf5, 0x84, 0x84, 0xe0, 0xe1, 0x0e, 0xf3, 0xfc, 0x0d,
	0xe6, 0xf9, 0x92, 0x79, 0x3e, 0x04, 0xa4, 0x5a, 0x8b, 0x11, 0xd6, 0x32, 0x74, 0x5b, 0xb0, 0xdf,
	0x05, 0x21, 0x0b, 0xb5, 0xcb, 0xca, 0x45, 0x87, 0xfb, 0x02, 0xec, 0xc5, 0xd7, 0x3f, 0xff, 0x7c,
	0xfa, 0xaf, 0x62, 0x4f, 0x53, 0xd6, 0x83, 0x3d, 0x16, 0xd0, 0x08, 0x3a, 0x48, 0x4f, 0x5b, 0x46,
	0x15, 0xf2, 0x0d, 0xe1, 0xfc, 0x1a, 0x64, 0x08, 0x08, 0x4d, 0xf8, 0x67, 0x2a, 0x0d, 0xf0, 0xe2,
	0xe5, 0x13, 0x34, 0x72, 0x5d, 0x21, 0x2f, 0xd8, 0xb3, 0x71, 0xe4, 0x16, 0x8c, 0x80, 0x7e, 0x83,
	0xf0, 0xf8, 0x4a, 0x53, 0x7a, 0x07, 0x4c, 0x82, 0x0a, 0x93, 0xa9, 0x84, 0x6f, 0x2c, 0x6a, 0xe0,
	0xca, 0xa3, 0x45, 0x1a, 0x68, 0x56, 0x01, 0x95, 0xec, 0x62, 0x1c, 0x88, 0x69, 0x71, 0xb5, 0x19,
	0xaa, 0x43, 0x8a, 0x77, 0x08, 0x5f, 0x5f, 0x03, 0x16, 0xe3, 0x98, 0x1e, 0xaa, 0x9f, 0xa5, 0x91,
	0xcc, 0x5c, 0x24, 0xd3, 0x2c, 0x73, 0x8a, 0x65, 0xca, 0xb6, 0x92, 0xcd, 0x19, 0xa6, 0xf9, 0x8c,
	0xf0, 0x0d, 0x33, 0x1e, 0x2b, 0x42, 0x80, 0x5c, 0x07, 0x20, 0x33, 0x19, 0xf3, 0x63, 0x04, 0x86,
	0x67, 0xf6, 0x42, 0x9d, 0x06, 0xa2, 0x0a, 0x68, 0xce, 0x2e, 0x1b, 0xa0, 0xe8, 0xe1, 0x83, 0x3c,
	0x1b, 0xb2, 0x2a, 0x0b, 0x93, 0xaa, 0x6d, 0x80, 0x65, 0x54, 0xa9, 0x1f, 0x5f, 0xc1, 0xd7, 0x9e,
	0x86, 0x67, 0xc5, 0x9c, 0x8e, 0x1f, 0x08, 0x4f, 0x3c, 0x61, 0x12, 0x84, 0x5c, 0x83, 0x0e, 0x17,
	0x9e, 0x5c, 0x69, 0xb5, 0x02, 0x10, 0x82, 0x54, 0x12, 0x0c, 0x69, 0x22, 0xc3, 0x3b, 0x7f, 0x29,
	0xad, 0x66, 0xde, 0x54, 0xcc, 0x8f, 0xc9, 0x23, 0x1a, 0x3f, 0xc2, 0x7b, 0x2a, 0xc9, 0x69, 0x45,
	0x59, 0x0e, 0x8b, 0xd2, 0xe8, 0x61, 0x00, 0x4d, 0xaf, 0xe3, 0x81, 0x2f, 0x1d, 0xd5, 0xdc, 0xa3,
	0xf3, 0x2b, 0xa1, 0xe8, 0x88, 0x7c, 0x44, 0xf8, 0xe6, 0xb3, 0x80, 0xf9, 0xa2, 0x0d, 0x81, 0x58,
	0xe7, 0xd1, 0x88, 0x92, 0x64, 0x13, 0x87, 0x14, 0x86, 0xfe, 0xde, 0xc5, 0x42, 0x8d, 0x5e, 0x51,
	0xe8, 0x65, 0x62, 0x27, 0xd0, 0xa5, 0xc9, 0x70, 0xda, 0x3c, 0x88, 0x30, 0x89, 0x83, 0xff, 0x0f,
	0x5f, 0x79, 0x3e, 0xb1, 0xf9, 0xb9, 0xb7, 0x5c, 0x48, 0x0b, 0x69, 0xa7, 0x92, 0x72, 0x9a, 0x24,
	0xf9, 0xf8, 0xa4, 0x19, 0xc3, 0x36, 0x00, 0x79, 0x85, 0xf0, 0x55, 0x83, 0x1a, 0x3a, 0x95, 0x32,
	0xca, 0x38, 0xe7, 0x68, 0x8f, 0x92, 0x0c, 0xd5, 0x98, 0xea, 0x6c, 0x4a, 0x75, 0x42, 0x84, 0x5d,
	0x9c, 0x53, 0x0d, 0x12, 0xa4, 0x98, 0xd8, 0x39, 0x5a, 0x36, 0xbe, 0x77, 0x32, 0xa2, 0xda, 0xb2,
	0xac, 0x2c, 0x2d, 0x52, 0x4c, 0xb7, 0x6c, 0x46, 0x16, 0x02, 0xe7, 0xd4, 0xfc, 0x0f, 0x9b, 0x45,
	0xcb, 0x59, 0x66, 0x26, 0xaa, 0xcd, 0x16, 0x94, 0xd9, 0x0c, 0x29, 0xa7, 0x9b, 0xa9, 0xd3, 0x22,
	0xe8, 0x61, 0x34, 0x6b, 0xe4, 0x2d, 0xc2, 0x58, 0xd1, 0x6e, 0x4b, 0x26, 0x81, 0xdc, 0x4d, 0x2b,
	0x44, 0x85, 0x8c, 0x7b, 0x69, 0x84, 0x42, 0x13, 0x2c, 0x29, 0x82, 0x79, 0x32, 0x37, 0xa2, 0x5c,
	0x47, 0x84, 0x29, 0x67, 0x18, 0x1f, 0x10, 0x1e, 0x8f, 0x9a, 0xb6, 0xda, 0x57, 0xf5, 0x0c, 0xdd,
	0xb0, 0xb1, 0x68, 0xd6, 0x0d, 0x9b, 0x10, 0x69, 0x9e, 0x07, 0x8a, 0x87, 0x92, 0xea, 0xa8, 0xf6,
	0x3b, 0x8d, 0xbe, 0xa3, 0x7a, 0x43, 0x0f, 0xd5, 0xe3, 0x68, 0x75, 0xf3, 0xf8, 0xd4, 0x42, 0x27,
	0xa7, 0x16, 0xfa, 0x7d, 0x6a, 0xa1, 0xf7, 0x03, 0x6b, 0xec, 0xfb, 0xc0, 0x42, 0x27, 0x03, 0x6b,
	0xec, 0xd7, 0xc0, 0x1a, 0x7b, 0x5e, 0x77, 0x3d, 0xb9, 0xd3, 0x6d, 0xd4, 0x9a, 0xfc, 0xc5, 0xbf,
	0x4b, 0xe9, 0x25, 0x0f, 0x76, 0xf5, 0xbf, 0x6a, 0x93, 0x07, 0x40, 0x7b, 0xda, 0x4b, 0xf6, 0x3b,
	0x20, 0x1a, 0x39, 0xf5, 0xfd, 0xbe, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x65, 0x2d, 0x05,
	0x54, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	RegisterChainMaintainer(ctx context.Context, in *RegisterChainMaintainerRequest, opts ...grpc.CallOption) (*RegisterChainMaintainerResponse, error)
	DeregisterChainMaintainer(ctx context.Context, in *DeregisterChainMaintainerRequest, opts ...grpc.CallOption) (*DeregisterChainMaintainerResponse, error)
	ActivateChain(ctx context.Context, in *ActivateChainRequest, opts ...grpc.CallOption) (*ActivateChainResponse, error)
	DeactivateChain(ctx context.Context, in *DeactivateChainRequest, opts ...grpc.CallOption) (*DeactivateChainResponse, error)
	RegisterAssetFee(ctx context.Context, in *RegisterAssetFeeRequest, opts ...grpc.CallOption) (*RegisterAssetFeeResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) RegisterChainMaintainer(ctx context.Context, in *RegisterChainMaintainerRequest, opts ...grpc.CallOption) (*RegisterChainMaintainerResponse, error) {
	out := new(RegisterChainMaintainerResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/RegisterChainMaintainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeregisterChainMaintainer(ctx context.Context, in *DeregisterChainMaintainerRequest, opts ...grpc.CallOption) (*DeregisterChainMaintainerResponse, error) {
	out := new(DeregisterChainMaintainerResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/DeregisterChainMaintainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ActivateChain(ctx context.Context, in *ActivateChainRequest, opts ...grpc.CallOption) (*ActivateChainResponse, error) {
	out := new(ActivateChainResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/ActivateChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeactivateChain(ctx context.Context, in *DeactivateChainRequest, opts ...grpc.CallOption) (*DeactivateChainResponse, error) {
	out := new(DeactivateChainResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/DeactivateChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RegisterAssetFee(ctx context.Context, in *RegisterAssetFeeRequest, opts ...grpc.CallOption) (*RegisterAssetFeeResponse, error) {
	out := new(RegisterAssetFeeResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/RegisterAssetFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	RegisterChainMaintainer(context.Context, *RegisterChainMaintainerRequest) (*RegisterChainMaintainerResponse, error)
	DeregisterChainMaintainer(context.Context, *DeregisterChainMaintainerRequest) (*DeregisterChainMaintainerResponse, error)
	ActivateChain(context.Context, *ActivateChainRequest) (*ActivateChainResponse, error)
	DeactivateChain(context.Context, *DeactivateChainRequest) (*DeactivateChainResponse, error)
	RegisterAssetFee(context.Context, *RegisterAssetFeeRequest) (*RegisterAssetFeeResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) RegisterChainMaintainer(ctx context.Context, req *RegisterChainMaintainerRequest) (*RegisterChainMaintainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterChainMaintainer not implemented")
}
func (*UnimplementedMsgServiceServer) DeregisterChainMaintainer(ctx context.Context, req *DeregisterChainMaintainerRequest) (*DeregisterChainMaintainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterChainMaintainer not implemented")
}
func (*UnimplementedMsgServiceServer) ActivateChain(ctx context.Context, req *ActivateChainRequest) (*ActivateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateChain not implemented")
}
func (*UnimplementedMsgServiceServer) DeactivateChain(ctx context.Context, req *DeactivateChainRequest) (*DeactivateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateChain not implemented")
}
func (*UnimplementedMsgServiceServer) RegisterAssetFee(ctx context.Context, req *RegisterAssetFeeRequest) (*RegisterAssetFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAssetFee not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_RegisterChainMaintainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterChainMaintainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterChainMaintainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/RegisterChainMaintainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterChainMaintainer(ctx, req.(*RegisterChainMaintainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeregisterChainMaintainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterChainMaintainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeregisterChainMaintainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/DeregisterChainMaintainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeregisterChainMaintainer(ctx, req.(*DeregisterChainMaintainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ActivateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ActivateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/ActivateChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ActivateChain(ctx, req.(*ActivateChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeactivateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeactivateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/DeactivateChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeactivateChain(ctx, req.(*DeactivateChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RegisterAssetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAssetFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterAssetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/RegisterAssetFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterAssetFee(ctx, req.(*RegisterAssetFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nexus.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterChainMaintainer",
			Handler:    _MsgService_RegisterChainMaintainer_Handler,
		},
		{
			MethodName: "DeregisterChainMaintainer",
			Handler:    _MsgService_DeregisterChainMaintainer_Handler,
		},
		{
			MethodName: "ActivateChain",
			Handler:    _MsgService_ActivateChain_Handler,
		},
		{
			MethodName: "DeactivateChain",
			Handler:    _MsgService_DeactivateChain_Handler,
		},
		{
			MethodName: "RegisterAssetFee",
			Handler:    _MsgService_RegisterAssetFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nexus/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// LatestDepositAddress queries the a deposit address by recipient
	LatestDepositAddress(ctx context.Context, in *LatestDepositAddressRequest, opts ...grpc.CallOption) (*LatestDepositAddressResponse, error)
	// TransfersForChain queries transfers by chain
	TransfersForChain(ctx context.Context, in *TransfersForChainRequest, opts ...grpc.CallOption) (*TransfersForChainResponse, error)
	// Fee queries the fee info by chain and asset
	Fee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error)
	// TransferFee queries the transfer fee by the source, destination chain,
	// asset and amount
	TransferFee(ctx context.Context, in *TransferFeeRequest, opts ...grpc.CallOption) (*TransferFeeResponse, error)
	// Chains queries the chains registered on the network
	Chains(ctx context.Context, in *ChainsRequest, opts ...grpc.CallOption) (*ChainsResponse, error)
	// Assets queries the assets registered for a chain
	Assets(ctx context.Context, in *AssetsRequest, opts ...grpc.CallOption) (*AssetsResponse, error)
	// ChainState queries the state of a registered chain on the network
	ChainState(ctx context.Context, in *ChainStateRequest, opts ...grpc.CallOption) (*ChainStateResponse, error)
	// ChainsByAsset queries the chains that support an asset on the network
	ChainsByAsset(ctx context.Context, in *ChainsByAssetRequest, opts ...grpc.CallOption) (*ChainsByAssetResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) LatestDepositAddress(ctx context.Context, in *LatestDepositAddressRequest, opts ...grpc.CallOption) (*LatestDepositAddressResponse, error) {
	out := new(LatestDepositAddressResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/LatestDepositAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) TransfersForChain(ctx context.Context, in *TransfersForChainRequest, opts ...grpc.CallOption) (*TransfersForChainResponse, error) {
	out := new(TransfersForChainResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/TransfersForChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Fee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/Fee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) TransferFee(ctx context.Context, in *TransferFeeRequest, opts ...grpc.CallOption) (*TransferFeeResponse, error) {
	out := new(TransferFeeResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/TransferFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Chains(ctx context.Context, in *ChainsRequest, opts ...grpc.CallOption) (*ChainsResponse, error) {
	out := new(ChainsResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/Chains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Assets(ctx context.Context, in *AssetsRequest, opts ...grpc.CallOption) (*AssetsResponse, error) {
	out := new(AssetsResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/Assets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) ChainState(ctx context.Context, in *ChainStateRequest, opts ...grpc.CallOption) (*ChainStateResponse, error) {
	out := new(ChainStateResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/ChainState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) ChainsByAsset(ctx context.Context, in *ChainsByAssetRequest, opts ...grpc.CallOption) (*ChainsByAssetResponse, error) {
	out := new(ChainsByAssetResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.QueryService/ChainsByAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// LatestDepositAddress queries the a deposit address by recipient
	LatestDepositAddress(context.Context, *LatestDepositAddressRequest) (*LatestDepositAddressResponse, error)
	// TransfersForChain queries transfers by chain
	TransfersForChain(context.Context, *TransfersForChainRequest) (*TransfersForChainResponse, error)
	// Fee queries the fee info by chain and asset
	Fee(context.Context, *FeeRequest) (*FeeResponse, error)
	// TransferFee queries the transfer fee by the source, destination chain,
	// asset and amount
	TransferFee(context.Context, *TransferFeeRequest) (*TransferFeeResponse, error)
	// Chains queries the chains registered on the network
	Chains(context.Context, *ChainsRequest) (*ChainsResponse, error)
	// Assets queries the assets registered for a chain
	Assets(context.Context, *AssetsRequest) (*AssetsResponse, error)
	// ChainState queries the state of a registered chain on the network
	ChainState(context.Context, *ChainStateRequest) (*ChainStateResponse, error)
	// ChainsByAsset queries the chains that support an asset on the network
	ChainsByAsset(context.Context, *ChainsByAssetRequest) (*ChainsByAssetResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) LatestDepositAddress(ctx context.Context, req *LatestDepositAddressRequest) (*LatestDepositAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestDepositAddress not implemented")
}
func (*UnimplementedQueryServiceServer) TransfersForChain(ctx context.Context, req *TransfersForChainRequest) (*TransfersForChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransfersForChain not implemented")
}
func (*UnimplementedQueryServiceServer) Fee(ctx context.Context, req *FeeRequest) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fee not implemented")
}
func (*UnimplementedQueryServiceServer) TransferFee(ctx context.Context, req *TransferFeeRequest) (*TransferFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferFee not implemented")
}
func (*UnimplementedQueryServiceServer) Chains(ctx context.Context, req *ChainsRequest) (*ChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chains not implemented")
}
func (*UnimplementedQueryServiceServer) Assets(ctx context.Context, req *AssetsRequest) (*AssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assets not implemented")
}
func (*UnimplementedQueryServiceServer) ChainState(ctx context.Context, req *ChainStateRequest) (*ChainStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainState not implemented")
}
func (*UnimplementedQueryServiceServer) ChainsByAsset(ctx context.Context, req *ChainsByAssetRequest) (*ChainsByAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainsByAsset not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_LatestDepositAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestDepositAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).LatestDepositAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/LatestDepositAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).LatestDepositAddress(ctx, req.(*LatestDepositAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_TransfersForChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransfersForChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).TransfersForChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/TransfersForChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).TransfersForChain(ctx, req.(*TransfersForChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Fee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Fee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/Fee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Fee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_TransferFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).TransferFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/TransferFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).TransferFee(ctx, req.(*TransferFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Chains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Chains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/Chains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Chains(ctx, req.(*ChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Assets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Assets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/Assets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Assets(ctx, req.(*AssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_ChainState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).ChainState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/ChainState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).ChainState(ctx, req.(*ChainStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_ChainsByAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainsByAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).ChainsByAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.QueryService/ChainsByAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).ChainsByAsset(ctx, req.(*ChainsByAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nexus.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LatestDepositAddress",
			Handler:    _QueryService_LatestDepositAddress_Handler,
		},
		{
			MethodName: "TransfersForChain",
			Handler:    _QueryService_TransfersForChain_Handler,
		},
		{
			MethodName: "Fee",
			Handler:    _QueryService_Fee_Handler,
		},
		{
			MethodName: "TransferFee",
			Handler:    _QueryService_TransferFee_Handler,
		},
		{
			MethodName: "Chains",
			Handler:    _QueryService_Chains_Handler,
		},
		{
			MethodName: "Assets",
			Handler:    _QueryService_Assets_Handler,
		},
		{
			MethodName: "ChainState",
			Handler:    _QueryService_ChainState_Handler,
		},
		{
			MethodName: "ChainsByAsset",
			Handler:    _QueryService_ChainsByAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nexus/v1beta1/service.proto",
}
