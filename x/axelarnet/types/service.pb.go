// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelarnetwork/axelarnet/v1beta1/service.proto

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

func init() {
	proto.RegisterFile("axelarnetwork/axelarnet/v1beta1/service.proto", fileDescriptor_ab91bd2713aa659f)
}
func init() {
	golang_proto.RegisterFile("axelarnetwork/axelarnet/v1beta1/service.proto", fileDescriptor_ab91bd2713aa659f)
}

var fileDescriptor_ab91bd2713aa659f = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0xc7, 0x19, 0x63, 0x3c, 0xd4, 0x5f, 0xb1, 0x92, 0x90, 0x10, 0x53, 0xb5, 0x06, 0x43, 0x80,
	0x69, 0xf9, 0xa1, 0xab, 0x22, 0x46, 0xd9, 0xaa, 0x09, 0x89, 0x26, 0x88, 0x9c, 0xbc, 0x90, 0xd9,
	0xe9, 0xa3, 0x3b, 0x61, 0x99, 0x29, 0x33, 0xb3, 0xb8, 0xc4, 0x78, 0xf1, 0x2f, 0x30, 0xf1, 0xe2,
	0xff, 0xe0, 0xc1, 0x8b, 0x47, 0x0f, 0x5e, 0x4c, 0x38, 0x12, 0xf5, 0xe0, 0xd1, 0xb0, 0x26, 0xfe,
	0x1b, 0x66, 0xda, 0x0e, 0x2e, 0x2c, 0xd8, 0x65, 0x6f, 0xdb, 0xec, 0xfb, 0xb4, 0x9f, 0xef, 0xeb,
	0xeb, 0x1b, 0x07, 0x93, 0x16, 0x34, 0x88, 0xe4, 0xa0, 0x5f, 0x0a, 0xb9, 0x16, 0xee, 0x5d, 0x85,
	0x9b, 0x53, 0x35, 0xd0, 0x64, 0x2a, 0x54, 0x20, 0x37, 0x19, 0x85, 0x20, 0x95, 0x42, 0x0b, 0xf7,
	0xf2, 0xbe, 0xf2, 0x60, 0xef, 0x2a, 0x28, 0xca, 0x87, 0x07, 0x13, 0x91, 0x88, 0xac, 0x36, 0x34,
	0xbf, 0x72, 0x6c, 0xf8, 0x52, 0x22, 0x44, 0xd2, 0x80, 0x90, 0xa4, 0x2c, 0x24, 0x9c, 0x0b, 0x4d,
	0x34, 0x13, 0x5c, 0x15, 0xff, 0x8e, 0x96, 0x39, 0xe8, 0x56, 0x51, 0x39, 0x5e, 0x56, 0xb9, 0xd1,
	0x04, 0xb9, 0x95, 0x17, 0x4f, 0xff, 0x39, 0xed, 0x38, 0x4f, 0x55, 0xf2, 0x3c, 0x0f, 0xe0, 0xbe,
	0x47, 0xce, 0xc9, 0x27, 0x8c, 0xaf, 0xb9, 0x13, 0x41, 0x49, 0x88, 0xc0, 0x94, 0x2d, 0xc1, 0x46,
	0x13, 0x94, 0x1e, 0xc6, 0x3d, 0x56, 0xab, 0x54, 0x70, 0x05, 0xfe, 0xcc, 0x9b, 0xef, 0xbf, 0xdf,
	0x9d, 0xc0, 0xfe, 0x68, 0x21, 0xd7, 0xe1, 0xd8, 0x60, 0x7c, 0x2d, 0x7c, 0x25, 0x81, 0xb2, 0x94,
	0x01, 0xd7, 0x2b, 0xb4, 0x4e, 0x18, 0x7f, 0x3d, 0x8b, 0xc6, 0xdc, 0x8f, 0xc8, 0x39, 0x17, 0x09,
	0xbe, 0xca, 0xe4, 0xfa, 0x43, 0x48, 0x85, 0x62, 0xda, 0xad, 0x94, 0x3e, 0x76, 0x3f, 0x60, 0x75,
	0x6f, 0x1d, 0x9b, 0x2b, 0xc4, 0x27, 0x32, 0xf1, 0xeb, 0xfe, 0xd5, 0x6e, 0x71, 0x9a, 0x13, 0x38,
	0xce, 0x11, 0x63, 0xfc, 0x03, 0x39, 0x43, 0x8f, 0x5a, 0x40, 0x9b, 0x1a, 0x16, 0x81, 0xc7, 0x8c,
	0x27, 0xcb, 0x92, 0x70, 0xb5, 0x0a, 0x52, 0xb9, 0xf7, 0x4b, 0x15, 0x8e, 0x20, 0x6d, 0x86, 0x07,
	0xfd, 0xdf, 0xa0, 0x08, 0x53, 0xc9, 0xc2, 0x4c, 0xfa, 0xe3, 0xdd, 0x61, 0x20, 0x47, 0x71, 0x9a,
	0xb3, 0x58, 0x5b, 0xd8, 0xc4, 0xfa, 0x84, 0x9c, 0xf3, 0x4b, 0x90, 0x30, 0xa5, 0x41, 0x2e, 0x54,
	0xa3, 0x45, 0xa2, 0xeb, 0x6e, 0x79, 0x47, 0x0f, 0x10, 0x36, 0xc6, 0xed, 0xe3, 0x83, 0x85, 0x7e,
	0x90, 0xe9, 0x8f, 0xfa, 0xd7, 0xba, 0xf5, 0x65, 0x81, 0x60, 0x56, 0xa3, 0x38, 0x25, 0xba, 0x6e,
	0xb4, 0xbf, 0x22, 0xe7, 0xe2, 0x7c, 0x1c, 0x47, 0x42, 0xad, 0x0b, 0x55, 0x25, 0x0a, 0xe2, 0xc8,
	0xcc, 0x96, 0x7b, 0xb7, 0xd4, 0xe0, 0x10, 0xca, 0xea, 0xcf, 0xf5, 0x07, 0x97, 0x7f, 0x07, 0x24,
	0x8e, 0x31, 0xcd, 0x38, 0x5c, 0x33, 0x20, 0xce, 0x3e, 0x04, 0x93, 0xe3, 0x03, 0x72, 0xce, 0xda,
	0x9e, 0xcc, 0x2b, 0x05, 0xda, 0xbd, 0xd9, 0x73, 0x0f, 0xb3, 0x7a, 0xeb, 0x5e, 0x39, 0x2e, 0x56,
	0x58, 0x8f, 0x67, 0xd6, 0x23, 0xfe, 0x95, 0xff, 0x34, 0x9e, 0x18, 0xc2, 0xd8, 0x7e, 0x46, 0xce,
	0x85, 0x25, 0xd1, 0xd4, 0xb0, 0x50, 0x8d, 0xfe, 0x4d, 0xff, 0x9d, 0xf2, 0x47, 0x1f, 0x64, 0xac,
	0xf5, 0x6c, 0x3f, 0x68, 0x61, 0x3e, 0x99, 0x99, 0x8f, 0xf9, 0x23, 0x87, 0x98, 0x1b, 0x28, 0x9b,
	0x97, 0x7d, 0xb3, 0xbe, 0x8d, 0x9c, 0x41, 0xdb, 0x85, 0xc7, 0x00, 0x91, 0x68, 0x34, 0x80, 0x6a,
	0x21, 0xdd, 0xb9, 0x9e, 0x9b, 0xd7, 0x89, 0xd9, 0x10, 0xf7, 0xfa, 0xa4, 0xcb, 0xe7, 0x66, 0xef,
	0x0d, 0xac, 0x02, 0x60, 0x6a, 0xc9, 0x59, 0x34, 0x36, 0xdd, 0x46, 0xce, 0x99, 0x67, 0x66, 0xf3,
	0xdb, 0x5d, 0xff, 0x0d, 0x39, 0x43, 0xc5, 0x72, 0xe8, 0xe8, 0x56, 0x24, 0x9a, 0x5c, 0xf7, 0xb0,
	0x9e, 0x8e, 0x20, 0x7b, 0x5f, 0x4f, 0x47, 0xde, 0xa0, 0x08, 0x79, 0x23, 0x0b, 0x19, 0xb8, 0x13,
	0xdd, 0x21, 0xed, 0x41, 0xc6, 0x6a, 0x74, 0xc5, 0xbe, 0xae, 0x15, 0x6a, 0xe8, 0xea, 0xf2, 0xf6,
	0xae, 0x87, 0x76, 0x76, 0x3d, 0xf4, 0x6b, 0xd7, 0x43, 0x6f, 0xdb, 0xde, 0xc0, 0x97, 0xb6, 0x87,
	0x76, 0xda, 0xde, 0xc0, 0xcf, 0xb6, 0x37, 0xf0, 0xa2, 0x92, 0x30, 0x5d, 0x6f, 0xd6, 0x02, 0x2a,
	0xd6, 0xc3, 0xc3, 0x4e, 0x49, 0x4c, 0x85, 0x84, 0xb0, 0xd5, 0xf1, 0x28, 0xbd, 0x95, 0x82, 0xaa,
	0x9d, 0xca, 0x0e, 0xcb, 0x99, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x07, 0xb7, 0x99, 0x09,
	0x08, 0x00, 0x00,
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
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error)
	ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error)
	ExecutePendingTransfers(ctx context.Context, in *ExecutePendingTransfersRequest, opts ...grpc.CallOption) (*ExecutePendingTransfersResponse, error)
	RegisterIBCPath(ctx context.Context, in *RegisterIBCPathRequest, opts ...grpc.CallOption) (*RegisterIBCPathResponse, error)
	AddCosmosBasedChain(ctx context.Context, in *AddCosmosBasedChainRequest, opts ...grpc.CallOption) (*AddCosmosBasedChainResponse, error)
	RegisterAsset(ctx context.Context, in *RegisterAssetRequest, opts ...grpc.CallOption) (*RegisterAssetResponse, error)
	RouteIBCTransfers(ctx context.Context, in *RouteIBCTransfersRequest, opts ...grpc.CallOption) (*RouteIBCTransfersResponse, error)
	RegisterFeeCollector(ctx context.Context, in *RegisterFeeCollectorRequest, opts ...grpc.CallOption) (*RegisterFeeCollectorResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error) {
	out := new(LinkResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error) {
	out := new(ConfirmDepositResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/ConfirmDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ExecutePendingTransfers(ctx context.Context, in *ExecutePendingTransfersRequest, opts ...grpc.CallOption) (*ExecutePendingTransfersResponse, error) {
	out := new(ExecutePendingTransfersResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/ExecutePendingTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RegisterIBCPath(ctx context.Context, in *RegisterIBCPathRequest, opts ...grpc.CallOption) (*RegisterIBCPathResponse, error) {
	out := new(RegisterIBCPathResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/RegisterIBCPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) AddCosmosBasedChain(ctx context.Context, in *AddCosmosBasedChainRequest, opts ...grpc.CallOption) (*AddCosmosBasedChainResponse, error) {
	out := new(AddCosmosBasedChainResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/AddCosmosBasedChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RegisterAsset(ctx context.Context, in *RegisterAssetRequest, opts ...grpc.CallOption) (*RegisterAssetResponse, error) {
	out := new(RegisterAssetResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/RegisterAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RouteIBCTransfers(ctx context.Context, in *RouteIBCTransfersRequest, opts ...grpc.CallOption) (*RouteIBCTransfersResponse, error) {
	out := new(RouteIBCTransfersResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/RouteIBCTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RegisterFeeCollector(ctx context.Context, in *RegisterFeeCollectorRequest, opts ...grpc.CallOption) (*RegisterFeeCollectorResponse, error) {
	out := new(RegisterFeeCollectorResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.MsgService/RegisterFeeCollector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	Link(context.Context, *LinkRequest) (*LinkResponse, error)
	ConfirmDeposit(context.Context, *ConfirmDepositRequest) (*ConfirmDepositResponse, error)
	ExecutePendingTransfers(context.Context, *ExecutePendingTransfersRequest) (*ExecutePendingTransfersResponse, error)
	RegisterIBCPath(context.Context, *RegisterIBCPathRequest) (*RegisterIBCPathResponse, error)
	AddCosmosBasedChain(context.Context, *AddCosmosBasedChainRequest) (*AddCosmosBasedChainResponse, error)
	RegisterAsset(context.Context, *RegisterAssetRequest) (*RegisterAssetResponse, error)
	RouteIBCTransfers(context.Context, *RouteIBCTransfersRequest) (*RouteIBCTransfersResponse, error)
	RegisterFeeCollector(context.Context, *RegisterFeeCollectorRequest) (*RegisterFeeCollectorResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) Link(ctx context.Context, req *LinkRequest) (*LinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmDeposit(ctx context.Context, req *ConfirmDepositRequest) (*ConfirmDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeposit not implemented")
}
func (*UnimplementedMsgServiceServer) ExecutePendingTransfers(ctx context.Context, req *ExecutePendingTransfersRequest) (*ExecutePendingTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutePendingTransfers not implemented")
}
func (*UnimplementedMsgServiceServer) RegisterIBCPath(ctx context.Context, req *RegisterIBCPathRequest) (*RegisterIBCPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterIBCPath not implemented")
}
func (*UnimplementedMsgServiceServer) AddCosmosBasedChain(ctx context.Context, req *AddCosmosBasedChainRequest) (*AddCosmosBasedChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCosmosBasedChain not implemented")
}
func (*UnimplementedMsgServiceServer) RegisterAsset(ctx context.Context, req *RegisterAssetRequest) (*RegisterAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAsset not implemented")
}
func (*UnimplementedMsgServiceServer) RouteIBCTransfers(ctx context.Context, req *RouteIBCTransfersRequest) (*RouteIBCTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteIBCTransfers not implemented")
}
func (*UnimplementedMsgServiceServer) RegisterFeeCollector(ctx context.Context, req *RegisterFeeCollectorRequest) (*RegisterFeeCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFeeCollector not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/ConfirmDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmDeposit(ctx, req.(*ConfirmDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ExecutePendingTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutePendingTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ExecutePendingTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/ExecutePendingTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ExecutePendingTransfers(ctx, req.(*ExecutePendingTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RegisterIBCPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterIBCPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterIBCPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/RegisterIBCPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterIBCPath(ctx, req.(*RegisterIBCPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_AddCosmosBasedChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCosmosBasedChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).AddCosmosBasedChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/AddCosmosBasedChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).AddCosmosBasedChain(ctx, req.(*AddCosmosBasedChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RegisterAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/RegisterAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterAsset(ctx, req.(*RegisterAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RouteIBCTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteIBCTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RouteIBCTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/RouteIBCTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RouteIBCTransfers(ctx, req.(*RouteIBCTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RegisterFeeCollector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterFeeCollectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterFeeCollector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.MsgService/RegisterFeeCollector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterFeeCollector(ctx, req.(*RegisterFeeCollectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelarnetwork.axelarnet.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Link",
			Handler:    _MsgService_Link_Handler,
		},
		{
			MethodName: "ConfirmDeposit",
			Handler:    _MsgService_ConfirmDeposit_Handler,
		},
		{
			MethodName: "ExecutePendingTransfers",
			Handler:    _MsgService_ExecutePendingTransfers_Handler,
		},
		{
			MethodName: "RegisterIBCPath",
			Handler:    _MsgService_RegisterIBCPath_Handler,
		},
		{
			MethodName: "AddCosmosBasedChain",
			Handler:    _MsgService_AddCosmosBasedChain_Handler,
		},
		{
			MethodName: "RegisterAsset",
			Handler:    _MsgService_RegisterAsset_Handler,
		},
		{
			MethodName: "RouteIBCTransfers",
			Handler:    _MsgService_RouteIBCTransfers_Handler,
		},
		{
			MethodName: "RegisterFeeCollector",
			Handler:    _MsgService_RegisterFeeCollector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelarnetwork/axelarnet/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	PendingIBCTransferCount(ctx context.Context, in *PendingIBCTransferCountRequest, opts ...grpc.CallOption) (*PendingIBCTransferCountResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) PendingIBCTransferCount(ctx context.Context, in *PendingIBCTransferCountRequest, opts ...grpc.CallOption) (*PendingIBCTransferCountResponse, error) {
	out := new(PendingIBCTransferCountResponse)
	err := c.cc.Invoke(ctx, "/axelarnetwork.axelarnet.v1beta1.QueryService/PendingIBCTransferCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	PendingIBCTransferCount(context.Context, *PendingIBCTransferCountRequest) (*PendingIBCTransferCountResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) PendingIBCTransferCount(ctx context.Context, req *PendingIBCTransferCountRequest) (*PendingIBCTransferCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingIBCTransferCount not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_PendingIBCTransferCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PendingIBCTransferCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PendingIBCTransferCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelarnetwork.axelarnet.v1beta1.QueryService/PendingIBCTransferCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PendingIBCTransferCount(ctx, req.(*PendingIBCTransferCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelarnetwork.axelarnet.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PendingIBCTransferCount",
			Handler:    _QueryService_PendingIBCTransferCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelarnetwork/axelarnet/v1beta1/service.proto",
}
