// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/routes/v1/route_service.proto

package routes

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("google/maps/routes/v1/route_service.proto", fileDescriptor_3616bef07f2cf21d)
}

var fileDescriptor_3616bef07f2cf21d = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4d, 0x2c, 0x28, 0xd6, 0x2f, 0xca, 0x2f, 0x2d, 0x49, 0x2d, 0xd6, 0x2f,
	0x33, 0x84, 0xb0, 0xe2, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x44, 0x21, 0x4a, 0xf5, 0x40, 0x4a, 0xf5, 0x20, 0x4a, 0xf5, 0xca, 0x0c, 0xa5, 0x64,
	0xa0, 0x26, 0x24, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7,
	0x15, 0x43, 0x34, 0x49, 0x19, 0x61, 0x37, 0x3f, 0x39, 0x3f, 0xb7, 0x00, 0x64, 0x03, 0x44, 0x24,
	0xbe, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x04, 0xaa, 0xc7, 0x98, 0x48, 0x3d, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0x50, 0xd7, 0x49, 0x89, 0x23, 0x39, 0x23, 0x39, 0x27, 0x33, 0x35, 0x0f, 0x6a, 0x9a,
	0xd1, 0x7e, 0x46, 0x2e, 0xfe, 0x20, 0xb0, 0x96, 0x80, 0xa2, 0xd4, 0xb4, 0xd4, 0xa2, 0xa2, 0xd4,
	0x14, 0xa1, 0x0e, 0x46, 0x2e, 0x5e, 0x67, 0x88, 0x71, 0x10, 0x29, 0x21, 0x6d, 0x3d, 0xac, 0xbe,
	0xd3, 0x43, 0x51, 0x15, 0x04, 0x71, 0xa6, 0x94, 0x0e, 0x71, 0x8a, 0x21, 0xee, 0x53, 0x92, 0x69,
	0xba, 0xfc, 0x64, 0x32, 0x93, 0x98, 0x92, 0xa0, 0x7e, 0x99, 0xa1, 0x55, 0x32, 0xb2, 0x12, 0x2b,
	0x46, 0x2d, 0x29, 0xc5, 0x53, 0x8e, 0x72, 0x10, 0x13, 0x0a, 0x60, 0xce, 0x83, 0x9a, 0x9e, 0x58,
	0x90, 0x59, 0xac, 0x97, 0x9c, 0x9f, 0xeb, 0xb4, 0x82, 0x91, 0x4b, 0x32, 0x39, 0x3f, 0x17, 0xbb,
	0xa5, 0x4e, 0x42, 0x10, 0xb3, 0x82, 0x21, 0x71, 0x15, 0x00, 0xf2, 0x73, 0x00, 0x63, 0x94, 0x35,
	0x54, 0x71, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e,
	0x38, 0x44, 0xf4, 0x11, 0xc6, 0xa3, 0x05, 0xb1, 0x35, 0x84, 0xf5, 0x83, 0x91, 0x71, 0x11, 0x13,
	0x8b, 0xbb, 0x6f, 0x50, 0xf0, 0x2a, 0x26, 0x51, 0x77, 0x88, 0x39, 0xbe, 0x20, 0x4b, 0x21, 0x36,
	0xe9, 0x85, 0x19, 0x9e, 0x82, 0x89, 0xc7, 0x80, 0xc4, 0x63, 0x20, 0xe2, 0x31, 0x61, 0x86, 0x49,
	0x6c, 0x60, 0x1b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0x85, 0x3f, 0x66, 0x57, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoutesPreferredClient is the client API for RoutesPreferred service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutesPreferredClient interface {
	// Returns the primary route along with optional alternate routes, given a set
	// of terminal and intermediate waypoints.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using URL parameter
	// `$fields` or `fields`, or by using an HTTP/gRPC header `X-Goog-FieldMask`
	// (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See detailed documentation about
	// [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of Route-level duration, distance, and polyline (an example
	// production setup):
	//   `X-Goog-FieldMask: routes.duration,routes.distanceMeters,routes.polyline`
	//
	// Google discourage the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need
	// in your production job ensures stable latency performance. We might add
	// more response fields in the future, and those new fields might require
	// extra computation time. If you select all fields, or if you select all
	// fields at the top level, then you might experience performance degradation
	// because any new field we add will be automatically included in the
	// response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRoutes(ctx context.Context, in *ComputeRoutesRequest, opts ...grpc.CallOption) (*ComputeRoutesResponse, error)
}

type routesPreferredClient struct {
	cc *grpc.ClientConn
}

func NewRoutesPreferredClient(cc *grpc.ClientConn) RoutesPreferredClient {
	return &routesPreferredClient{cc}
}

func (c *routesPreferredClient) ComputeRoutes(ctx context.Context, in *ComputeRoutesRequest, opts ...grpc.CallOption) (*ComputeRoutesResponse, error) {
	out := new(ComputeRoutesResponse)
	err := c.cc.Invoke(ctx, "/google.maps.routes.v1.RoutesPreferred/ComputeRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutesPreferredServer is the server API for RoutesPreferred service.
type RoutesPreferredServer interface {
	// Returns the primary route along with optional alternate routes, given a set
	// of terminal and intermediate waypoints.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using URL parameter
	// `$fields` or `fields`, or by using an HTTP/gRPC header `X-Goog-FieldMask`
	// (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See detailed documentation about
	// [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of Route-level duration, distance, and polyline (an example
	// production setup):
	//   `X-Goog-FieldMask: routes.duration,routes.distanceMeters,routes.polyline`
	//
	// Google discourage the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need
	// in your production job ensures stable latency performance. We might add
	// more response fields in the future, and those new fields might require
	// extra computation time. If you select all fields, or if you select all
	// fields at the top level, then you might experience performance degradation
	// because any new field we add will be automatically included in the
	// response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRoutes(context.Context, *ComputeRoutesRequest) (*ComputeRoutesResponse, error)
}

// UnimplementedRoutesPreferredServer can be embedded to have forward compatible implementations.
type UnimplementedRoutesPreferredServer struct {
}

func (*UnimplementedRoutesPreferredServer) ComputeRoutes(ctx context.Context, req *ComputeRoutesRequest) (*ComputeRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeRoutes not implemented")
}

func RegisterRoutesPreferredServer(s *grpc.Server, srv RoutesPreferredServer) {
	s.RegisterService(&_RoutesPreferred_serviceDesc, srv)
}

func _RoutesPreferred_ComputeRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesPreferredServer).ComputeRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.routes.v1.RoutesPreferred/ComputeRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesPreferredServer).ComputeRoutes(ctx, req.(*ComputeRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutesPreferred_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.routes.v1.RoutesPreferred",
	HandlerType: (*RoutesPreferredServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeRoutes",
			Handler:    _RoutesPreferred_ComputeRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/maps/routes/v1/route_service.proto",
}
