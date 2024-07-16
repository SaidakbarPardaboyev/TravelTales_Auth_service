// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: itineraries.proto

package itineraries

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ItinerariesClient is the client API for Itineraries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItinerariesClient interface {
	CreateItineraries(ctx context.Context, in *RequestCreateItineraries, opts ...grpc.CallOption) (*ResponseCreateItineraries, error)
	EditItineraries(ctx context.Context, in *RequestEditItineraries, opts ...grpc.CallOption) (*ResponseEditItineraries, error)
	DeleteItineraries(ctx context.Context, in *RequestDeleteItineraries, opts ...grpc.CallOption) (*ResponseDeleteItineraries, error)
	GetAllItineraries(ctx context.Context, in *RequestGetAllItineraries, opts ...grpc.CallOption) (*ResponseGetAllItineraries, error)
	GetItineraryFullInfo(ctx context.Context, in *RequestGetItineraryFullInfo, opts ...grpc.CallOption) (*ResponseGetItineraryFullInfo, error)
	WriteCommentToItinerary(ctx context.Context, in *RequestWriteCommentToItinerary, opts ...grpc.CallOption) (*ResponseWriteCommentToItinerary, error)
	GetDestinations(ctx context.Context, in *RequestGetDestinations, opts ...grpc.CallOption) (*ResponseGetDestinations, error)
	GetDestinationsAllInfo(ctx context.Context, in *RequestGetDestinationsAllInfo, opts ...grpc.CallOption) (*ResponseGetDestinationsAllInfo, error)
	WriteMessages(ctx context.Context, in *RequestWriteMessages, opts ...grpc.CallOption) (*ResponseWriteMessages, error)
	GetMessages(ctx context.Context, in *RequestGetMessages, opts ...grpc.CallOption) (*ResponseGetMessages, error)
	GetUserStatistic(ctx context.Context, in *RequestGetUserStatistic, opts ...grpc.CallOption) (*ResponseGetUserStatistic, error)
}

type itinerariesClient struct {
	cc grpc.ClientConnInterface
}

func NewItinerariesClient(cc grpc.ClientConnInterface) ItinerariesClient {
	return &itinerariesClient{cc}
}

func (c *itinerariesClient) CreateItineraries(ctx context.Context, in *RequestCreateItineraries, opts ...grpc.CallOption) (*ResponseCreateItineraries, error) {
	out := new(ResponseCreateItineraries)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/CreateItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) EditItineraries(ctx context.Context, in *RequestEditItineraries, opts ...grpc.CallOption) (*ResponseEditItineraries, error) {
	out := new(ResponseEditItineraries)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/EditItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) DeleteItineraries(ctx context.Context, in *RequestDeleteItineraries, opts ...grpc.CallOption) (*ResponseDeleteItineraries, error) {
	out := new(ResponseDeleteItineraries)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/DeleteItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetAllItineraries(ctx context.Context, in *RequestGetAllItineraries, opts ...grpc.CallOption) (*ResponseGetAllItineraries, error) {
	out := new(ResponseGetAllItineraries)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/GetAllItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetItineraryFullInfo(ctx context.Context, in *RequestGetItineraryFullInfo, opts ...grpc.CallOption) (*ResponseGetItineraryFullInfo, error) {
	out := new(ResponseGetItineraryFullInfo)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/GetItineraryFullInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) WriteCommentToItinerary(ctx context.Context, in *RequestWriteCommentToItinerary, opts ...grpc.CallOption) (*ResponseWriteCommentToItinerary, error) {
	out := new(ResponseWriteCommentToItinerary)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/WriteCommentToItinerary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetDestinations(ctx context.Context, in *RequestGetDestinations, opts ...grpc.CallOption) (*ResponseGetDestinations, error) {
	out := new(ResponseGetDestinations)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/GetDestinations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetDestinationsAllInfo(ctx context.Context, in *RequestGetDestinationsAllInfo, opts ...grpc.CallOption) (*ResponseGetDestinationsAllInfo, error) {
	out := new(ResponseGetDestinationsAllInfo)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/GetDestinationsAllInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) WriteMessages(ctx context.Context, in *RequestWriteMessages, opts ...grpc.CallOption) (*ResponseWriteMessages, error) {
	out := new(ResponseWriteMessages)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/WriteMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetMessages(ctx context.Context, in *RequestGetMessages, opts ...grpc.CallOption) (*ResponseGetMessages, error) {
	out := new(ResponseGetMessages)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetUserStatistic(ctx context.Context, in *RequestGetUserStatistic, opts ...grpc.CallOption) (*ResponseGetUserStatistic, error) {
	out := new(ResponseGetUserStatistic)
	err := c.cc.Invoke(ctx, "/itineraries.itineraries/GetUserStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItinerariesServer is the server API for Itineraries service.
// All implementations must embed UnimplementedItinerariesServer
// for forward compatibility
type ItinerariesServer interface {
	CreateItineraries(context.Context, *RequestCreateItineraries) (*ResponseCreateItineraries, error)
	EditItineraries(context.Context, *RequestEditItineraries) (*ResponseEditItineraries, error)
	DeleteItineraries(context.Context, *RequestDeleteItineraries) (*ResponseDeleteItineraries, error)
	GetAllItineraries(context.Context, *RequestGetAllItineraries) (*ResponseGetAllItineraries, error)
	GetItineraryFullInfo(context.Context, *RequestGetItineraryFullInfo) (*ResponseGetItineraryFullInfo, error)
	WriteCommentToItinerary(context.Context, *RequestWriteCommentToItinerary) (*ResponseWriteCommentToItinerary, error)
	GetDestinations(context.Context, *RequestGetDestinations) (*ResponseGetDestinations, error)
	GetDestinationsAllInfo(context.Context, *RequestGetDestinationsAllInfo) (*ResponseGetDestinationsAllInfo, error)
	WriteMessages(context.Context, *RequestWriteMessages) (*ResponseWriteMessages, error)
	GetMessages(context.Context, *RequestGetMessages) (*ResponseGetMessages, error)
	GetUserStatistic(context.Context, *RequestGetUserStatistic) (*ResponseGetUserStatistic, error)
	mustEmbedUnimplementedItinerariesServer()
}

// UnimplementedItinerariesServer must be embedded to have forward compatible implementations.
type UnimplementedItinerariesServer struct {
}

func (UnimplementedItinerariesServer) CreateItineraries(context.Context, *RequestCreateItineraries) (*ResponseCreateItineraries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItineraries not implemented")
}
func (UnimplementedItinerariesServer) EditItineraries(context.Context, *RequestEditItineraries) (*ResponseEditItineraries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditItineraries not implemented")
}
func (UnimplementedItinerariesServer) DeleteItineraries(context.Context, *RequestDeleteItineraries) (*ResponseDeleteItineraries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItineraries not implemented")
}
func (UnimplementedItinerariesServer) GetAllItineraries(context.Context, *RequestGetAllItineraries) (*ResponseGetAllItineraries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItineraries not implemented")
}
func (UnimplementedItinerariesServer) GetItineraryFullInfo(context.Context, *RequestGetItineraryFullInfo) (*ResponseGetItineraryFullInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItineraryFullInfo not implemented")
}
func (UnimplementedItinerariesServer) WriteCommentToItinerary(context.Context, *RequestWriteCommentToItinerary) (*ResponseWriteCommentToItinerary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteCommentToItinerary not implemented")
}
func (UnimplementedItinerariesServer) GetDestinations(context.Context, *RequestGetDestinations) (*ResponseGetDestinations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinations not implemented")
}
func (UnimplementedItinerariesServer) GetDestinationsAllInfo(context.Context, *RequestGetDestinationsAllInfo) (*ResponseGetDestinationsAllInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationsAllInfo not implemented")
}
func (UnimplementedItinerariesServer) WriteMessages(context.Context, *RequestWriteMessages) (*ResponseWriteMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMessages not implemented")
}
func (UnimplementedItinerariesServer) GetMessages(context.Context, *RequestGetMessages) (*ResponseGetMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedItinerariesServer) GetUserStatistic(context.Context, *RequestGetUserStatistic) (*ResponseGetUserStatistic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStatistic not implemented")
}
func (UnimplementedItinerariesServer) mustEmbedUnimplementedItinerariesServer() {}

// UnsafeItinerariesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItinerariesServer will
// result in compilation errors.
type UnsafeItinerariesServer interface {
	mustEmbedUnimplementedItinerariesServer()
}

func RegisterItinerariesServer(s grpc.ServiceRegistrar, srv ItinerariesServer) {
	s.RegisterService(&Itineraries_ServiceDesc, srv)
}

func _Itineraries_CreateItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateItineraries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).CreateItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/CreateItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).CreateItineraries(ctx, req.(*RequestCreateItineraries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_EditItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditItineraries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).EditItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/EditItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).EditItineraries(ctx, req.(*RequestEditItineraries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_DeleteItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteItineraries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).DeleteItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/DeleteItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).DeleteItineraries(ctx, req.(*RequestDeleteItineraries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetAllItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetAllItineraries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetAllItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/GetAllItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetAllItineraries(ctx, req.(*RequestGetAllItineraries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetItineraryFullInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetItineraryFullInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetItineraryFullInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/GetItineraryFullInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetItineraryFullInfo(ctx, req.(*RequestGetItineraryFullInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_WriteCommentToItinerary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWriteCommentToItinerary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).WriteCommentToItinerary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/WriteCommentToItinerary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).WriteCommentToItinerary(ctx, req.(*RequestWriteCommentToItinerary))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetDestinations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/GetDestinations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetDestinations(ctx, req.(*RequestGetDestinations))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetDestinationsAllInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetDestinationsAllInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetDestinationsAllInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/GetDestinationsAllInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetDestinationsAllInfo(ctx, req.(*RequestGetDestinationsAllInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_WriteMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWriteMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).WriteMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/WriteMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).WriteMessages(ctx, req.(*RequestWriteMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetMessages(ctx, req.(*RequestGetMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetUserStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetUserStatistic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetUserStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.itineraries/GetUserStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetUserStatistic(ctx, req.(*RequestGetUserStatistic))
	}
	return interceptor(ctx, in, info, handler)
}

// Itineraries_ServiceDesc is the grpc.ServiceDesc for Itineraries service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Itineraries_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "itineraries.itineraries",
	HandlerType: (*ItinerariesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItineraries",
			Handler:    _Itineraries_CreateItineraries_Handler,
		},
		{
			MethodName: "EditItineraries",
			Handler:    _Itineraries_EditItineraries_Handler,
		},
		{
			MethodName: "DeleteItineraries",
			Handler:    _Itineraries_DeleteItineraries_Handler,
		},
		{
			MethodName: "GetAllItineraries",
			Handler:    _Itineraries_GetAllItineraries_Handler,
		},
		{
			MethodName: "GetItineraryFullInfo",
			Handler:    _Itineraries_GetItineraryFullInfo_Handler,
		},
		{
			MethodName: "WriteCommentToItinerary",
			Handler:    _Itineraries_WriteCommentToItinerary_Handler,
		},
		{
			MethodName: "GetDestinations",
			Handler:    _Itineraries_GetDestinations_Handler,
		},
		{
			MethodName: "GetDestinationsAllInfo",
			Handler:    _Itineraries_GetDestinationsAllInfo_Handler,
		},
		{
			MethodName: "WriteMessages",
			Handler:    _Itineraries_WriteMessages_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Itineraries_GetMessages_Handler,
		},
		{
			MethodName: "GetUserStatistic",
			Handler:    _Itineraries_GetUserStatistic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "itineraries.proto",
}
