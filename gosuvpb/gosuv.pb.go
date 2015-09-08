// Code generated by protoc-gen-go.
// source: gosuv.proto
// DO NOT EDIT!

/*
Package gosuvpb is a generated protocol buffer package.

It is generated from these files:
	gosuv.proto

It has these top-level messages:
	NopRequest
	Response
	Request
	ProgramStatus
	StatusResponse
*/
package gosuvpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NopRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NopRequest) Reset()         { *m = NopRequest{} }
func (m *NopRequest) String() string { return proto.CompactTextString(m) }
func (*NopRequest) ProtoMessage()    {}

type Response struct {
	Code             *int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type Request struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ProgramStatus struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Status           *string `protobuf:"bytes,2,req,name=status" json:"status,omitempty"`
	Extra            *string `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProgramStatus) Reset()         { *m = ProgramStatus{} }
func (m *ProgramStatus) String() string { return proto.CompactTextString(m) }
func (*ProgramStatus) ProtoMessage()    {}

func (m *ProgramStatus) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProgramStatus) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *ProgramStatus) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

type StatusResponse struct {
	Programs         []*ProgramStatus `protobuf:"bytes,1,rep,name=programs" json:"programs,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}

func (m *StatusResponse) GetPrograms() []*ProgramStatus {
	if m != nil {
		return m.Programs
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GoSuv service

type GoSuvClient interface {
	Shutdown(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error)
	Version(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error)
	Status(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type goSuvClient struct {
	cc *grpc.ClientConn
}

func NewGoSuvClient(cc *grpc.ClientConn) GoSuvClient {
	return &goSuvClient{cc}
}

func (c *goSuvClient) Shutdown(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Version(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Status(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoSuv service

type GoSuvServer interface {
	Shutdown(context.Context, *NopRequest) (*Response, error)
	Version(context.Context, *NopRequest) (*Response, error)
	Status(context.Context, *NopRequest) (*StatusResponse, error)
}

func RegisterGoSuvServer(s *grpc.Server, srv GoSuvServer) {
	s.RegisterService(&_GoSuv_serviceDesc, srv)
}

func _GoSuv_Shutdown_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NopRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Shutdown(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Version_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NopRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Version(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Status_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NopRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GoSuv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosuvpb.GoSuv",
	HandlerType: (*GoSuvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shutdown",
			Handler:    _GoSuv_Shutdown_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _GoSuv_Version_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _GoSuv_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Program service

type ProgramClient interface {
	Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type programClient struct {
	cc *grpc.ClientConn
}

func NewProgramClient(cc *grpc.ClientConn) ProgramClient {
	return &programClient{cc}
}

func (c *programClient) Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Program service

type ProgramServer interface {
	Start(context.Context, *Request) (*Response, error)
	Stop(context.Context, *Request) (*Response, error)
}

func RegisterProgramServer(s *grpc.Server, srv ProgramServer) {
	s.RegisterService(&_Program_serviceDesc, srv)
}

func _Program_Start_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Start(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Program_Stop_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Stop(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Program_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosuvpb.Program",
	HandlerType: (*ProgramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Program_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Program_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
