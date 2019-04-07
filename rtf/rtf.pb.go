// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rtf.proto

package rtf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type RTFStatement struct {
	// ID of the client
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// ID of the current statement
	NodeId int64 `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// ID of the parent of this statement, if any
	ParentId int64 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Ordered list of the contexts, every context has a ID
	// The context are created by the python statement `with`
	Contexts []int64 `protobuf:"varint,4,rep,packed,name=contexts,proto3" json:"contexts,omitempty"`
	// The statement sent
	Stmt                 string   `protobuf:"bytes,5,opt,name=stmt,proto3" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RTFStatement) Reset()         { *m = RTFStatement{} }
func (m *RTFStatement) String() string { return proto.CompactTextString(m) }
func (*RTFStatement) ProtoMessage()    {}
func (*RTFStatement) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f3d763eab9a67f5, []int{0}
}

func (m *RTFStatement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTFStatement.Unmarshal(m, b)
}
func (m *RTFStatement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTFStatement.Marshal(b, m, deterministic)
}
func (m *RTFStatement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTFStatement.Merge(m, src)
}
func (m *RTFStatement) XXX_Size() int {
	return xxx_messageInfo_RTFStatement.Size(m)
}
func (m *RTFStatement) XXX_DiscardUnknown() {
	xxx_messageInfo_RTFStatement.DiscardUnknown(m)
}

var xxx_messageInfo_RTFStatement proto.InternalMessageInfo

func (m *RTFStatement) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RTFStatement) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RTFStatement) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *RTFStatement) GetContexts() []int64 {
	if m != nil {
		return m.Contexts
	}
	return nil
}

func (m *RTFStatement) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

type RTFResponse struct {
	// The ID of the node that has been executed
	// to generate the response.
	NodeId int64 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Execxution status
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// Captured Standard Output
	Stdout string `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// The response value
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RTFResponse) Reset()         { *m = RTFResponse{} }
func (m *RTFResponse) String() string { return proto.CompactTextString(m) }
func (*RTFResponse) ProtoMessage()    {}
func (*RTFResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f3d763eab9a67f5, []int{1}
}

func (m *RTFResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTFResponse.Unmarshal(m, b)
}
func (m *RTFResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTFResponse.Marshal(b, m, deterministic)
}
func (m *RTFResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTFResponse.Merge(m, src)
}
func (m *RTFResponse) XXX_Size() int {
	return xxx_messageInfo_RTFResponse.Size(m)
}
func (m *RTFResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RTFResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RTFResponse proto.InternalMessageInfo

func (m *RTFResponse) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RTFResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *RTFResponse) GetStdout() string {
	if m != nil {
		return m.Stdout
	}
	return ""
}

func (m *RTFResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*RTFStatement)(nil), "RTFStatement")
	proto.RegisterType((*RTFResponse)(nil), "RTFResponse")
}

func init() { proto.RegisterFile("rtf.proto", fileDescriptor_4f3d763eab9a67f5) }

var fileDescriptor_4f3d763eab9a67f5 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x12, 0x42, 0x72, 0xa4, 0x8b, 0x07, 0xb0, 0xca, 0x12, 0x75, 0xca, 0x14, 0x55,
	0x30, 0x30, 0x23, 0x50, 0xa4, 0xae, 0x47, 0x76, 0x94, 0xe2, 0xab, 0x14, 0xd4, 0xda, 0x51, 0x7c,
	0x96, 0xe0, 0x1f, 0xf0, 0xb3, 0x91, 0x9d, 0x22, 0xa5, 0xdb, 0xfb, 0xee, 0x0d, 0xdf, 0xd3, 0x41,
	0x31, 0xf1, 0xa1, 0x19, 0x27, 0xcb, 0x76, 0xf3, 0x2b, 0xa0, 0xc4, 0xae, 0x7d, 0xe7, 0x9e, 0xe9,
	0x44, 0x86, 0xa5, 0x84, 0xd4, 0xfb, 0x41, 0x2b, 0x51, 0x89, 0xba, 0xc0, 0x98, 0xe5, 0x3d, 0xdc,
	0x18, 0xab, 0xe9, 0x63, 0xd0, 0xea, 0xaa, 0x12, 0x75, 0x82, 0x59, 0xc0, 0x9d, 0x96, 0x0f, 0x50,
	0x8c, 0xfd, 0x44, 0x86, 0x43, 0x95, 0xc4, 0x2a, 0x9f, 0x0f, 0x3b, 0x2d, 0xd7, 0x90, 0x7f, 0x5a,
	0xc3, 0xf4, 0xcd, 0x4e, 0xa5, 0x55, 0x12, 0xba, 0x7f, 0x0e, 0x16, 0xc7, 0x27, 0x56, 0xd7, 0xb3,
	0x25, 0xe4, 0xcd, 0x17, 0xdc, 0x62, 0xd7, 0x22, 0xb9, 0xd1, 0x1a, 0x47, 0x4b, 0xa9, 0xb8, 0x90,
	0xde, 0x41, 0xe6, 0xb8, 0x67, 0xef, 0xe2, 0x98, 0x1c, 0xcf, 0x34, 0xdf, 0xb5, 0xf5, 0x1c, 0x97,
	0x14, 0x78, 0xa6, 0xe0, 0xda, 0x5b, 0xfd, 0xa3, 0xd2, 0x4a, 0xd4, 0x25, 0xc6, 0xfc, 0xf8, 0x0c,
	0x09, 0x76, 0xad, 0xdc, 0xc2, 0xea, 0x8d, 0x0e, 0x83, 0xa1, 0x17, 0xa3, 0x5f, 0xfb, 0xe3, 0x51,
	0xae, 0x9a, 0xe5, 0x33, 0xd6, 0x65, 0xb3, 0x58, 0x54, 0x8b, 0xad, 0xd8, 0x67, 0xf1, 0x6d, 0x4f,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xc2, 0x0c, 0x78, 0x43, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RTFClient is the client API for RTF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RTFClient interface {
	// accept a stream of RTFStatement that define the Python function body.
	// Returns a stream of RTFResponse.
	DefineAndCall(ctx context.Context, opts ...grpc.CallOption) (RTF_DefineAndCallClient, error)
}

type rTFClient struct {
	cc *grpc.ClientConn
}

func NewRTFClient(cc *grpc.ClientConn) RTFClient {
	return &rTFClient{cc}
}

func (c *rTFClient) DefineAndCall(ctx context.Context, opts ...grpc.CallOption) (RTF_DefineAndCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RTF_serviceDesc.Streams[0], "/RTF/DefineAndCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &rTFDefineAndCallClient{stream}
	return x, nil
}

type RTF_DefineAndCallClient interface {
	Send(*RTFStatement) error
	Recv() (*RTFResponse, error)
	grpc.ClientStream
}

type rTFDefineAndCallClient struct {
	grpc.ClientStream
}

func (x *rTFDefineAndCallClient) Send(m *RTFStatement) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rTFDefineAndCallClient) Recv() (*RTFResponse, error) {
	m := new(RTFResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RTFServer is the server API for RTF service.
type RTFServer interface {
	// accept a stream of RTFStatement that define the Python function body.
	// Returns a stream of RTFResponse.
	DefineAndCall(RTF_DefineAndCallServer) error
}

// UnimplementedRTFServer can be embedded to have forward compatible implementations.
type UnimplementedRTFServer struct {
}

func (*UnimplementedRTFServer) DefineAndCall(srv RTF_DefineAndCallServer) error {
	return status.Errorf(codes.Unimplemented, "method DefineAndCall not implemented")
}

func RegisterRTFServer(s *grpc.Server, srv RTFServer) {
	s.RegisterService(&_RTF_serviceDesc, srv)
}

func _RTF_DefineAndCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RTFServer).DefineAndCall(&rTFDefineAndCallServer{stream})
}

type RTF_DefineAndCallServer interface {
	Send(*RTFResponse) error
	Recv() (*RTFStatement, error)
	grpc.ServerStream
}

type rTFDefineAndCallServer struct {
	grpc.ServerStream
}

func (x *rTFDefineAndCallServer) Send(m *RTFResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rTFDefineAndCallServer) Recv() (*RTFStatement, error) {
	m := new(RTFStatement)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RTF_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RTF",
	HandlerType: (*RTFServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DefineAndCall",
			Handler:       _RTF_DefineAndCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rtf.proto",
}
