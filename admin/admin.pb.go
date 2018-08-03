// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.proto

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReqClassCreate struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Standard             string   `protobuf:"bytes,2,opt,name=standard,proto3" json:"standard,omitempty"`
	Division             string   `protobuf:"bytes,3,opt,name=division,proto3" json:"division,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClassCreate) Reset()         { *m = ReqClassCreate{} }
func (m *ReqClassCreate) String() string { return proto.CompactTextString(m) }
func (*ReqClassCreate) ProtoMessage()    {}
func (*ReqClassCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{0}
}
func (m *ReqClassCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClassCreate.Unmarshal(m, b)
}
func (m *ReqClassCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClassCreate.Marshal(b, m, deterministic)
}
func (dst *ReqClassCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClassCreate.Merge(dst, src)
}
func (m *ReqClassCreate) XXX_Size() int {
	return xxx_messageInfo_ReqClassCreate.Size(m)
}
func (m *ReqClassCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClassCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClassCreate proto.InternalMessageInfo

func (m *ReqClassCreate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReqClassCreate) GetStandard() string {
	if m != nil {
		return m.Standard
	}
	return ""
}

func (m *ReqClassCreate) GetDivision() string {
	if m != nil {
		return m.Division
	}
	return ""
}

type ResSuccess struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResSuccess) Reset()         { *m = ResSuccess{} }
func (m *ResSuccess) String() string { return proto.CompactTextString(m) }
func (*ResSuccess) ProtoMessage()    {}
func (*ResSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{1}
}
func (m *ResSuccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResSuccess.Unmarshal(m, b)
}
func (m *ResSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResSuccess.Marshal(b, m, deterministic)
}
func (dst *ResSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResSuccess.Merge(dst, src)
}
func (m *ResSuccess) XXX_Size() int {
	return xxx_messageInfo_ResSuccess.Size(m)
}
func (m *ResSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_ResSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_ResSuccess proto.InternalMessageInfo

func (m *ResSuccess) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResSuccess) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ClassInfo struct {
	Standard             string   `protobuf:"bytes,1,opt,name=standard,proto3" json:"standard,omitempty"`
	Division             string   `protobuf:"bytes,2,opt,name=division,proto3" json:"division,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassInfo) Reset()         { *m = ClassInfo{} }
func (m *ClassInfo) String() string { return proto.CompactTextString(m) }
func (*ClassInfo) ProtoMessage()    {}
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{2}
}
func (m *ClassInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassInfo.Unmarshal(m, b)
}
func (m *ClassInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassInfo.Marshal(b, m, deterministic)
}
func (dst *ClassInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassInfo.Merge(dst, src)
}
func (m *ClassInfo) XXX_Size() int {
	return xxx_messageInfo_ClassInfo.Size(m)
}
func (m *ClassInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClassInfo proto.InternalMessageInfo

func (m *ClassInfo) GetStandard() string {
	if m != nil {
		return m.Standard
	}
	return ""
}

func (m *ClassInfo) GetDivision() string {
	if m != nil {
		return m.Division
	}
	return ""
}

type ReqClass struct {
	ClassId              string   `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClass) Reset()         { *m = ReqClass{} }
func (m *ReqClass) String() string { return proto.CompactTextString(m) }
func (*ReqClass) ProtoMessage()    {}
func (*ReqClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{3}
}
func (m *ReqClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClass.Unmarshal(m, b)
}
func (m *ReqClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClass.Marshal(b, m, deterministic)
}
func (dst *ReqClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClass.Merge(dst, src)
}
func (m *ReqClass) XXX_Size() int {
	return xxx_messageInfo_ReqClass.Size(m)
}
func (m *ReqClass) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClass.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClass proto.InternalMessageInfo

func (m *ReqClass) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

type ResClass struct {
	Class                *ClassInfo     `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Students             []*StudentInfo `protobuf:"bytes,2,rep,name=students,proto3" json:"students,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResClass) Reset()         { *m = ResClass{} }
func (m *ResClass) String() string { return proto.CompactTextString(m) }
func (*ResClass) ProtoMessage()    {}
func (*ResClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{4}
}
func (m *ResClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResClass.Unmarshal(m, b)
}
func (m *ResClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResClass.Marshal(b, m, deterministic)
}
func (dst *ResClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResClass.Merge(dst, src)
}
func (m *ResClass) XXX_Size() int {
	return xxx_messageInfo_ResClass.Size(m)
}
func (m *ResClass) XXX_DiscardUnknown() {
	xxx_messageInfo_ResClass.DiscardUnknown(m)
}

var xxx_messageInfo_ResClass proto.InternalMessageInfo

func (m *ResClass) GetClass() *ClassInfo {
	if m != nil {
		return m.Class
	}
	return nil
}

func (m *ResClass) GetStudents() []*StudentInfo {
	if m != nil {
		return m.Students
	}
	return nil
}

type StudentInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentInfo) Reset()         { *m = StudentInfo{} }
func (m *StudentInfo) String() string { return proto.CompactTextString(m) }
func (*StudentInfo) ProtoMessage()    {}
func (*StudentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{5}
}
func (m *StudentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudentInfo.Unmarshal(m, b)
}
func (m *StudentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudentInfo.Marshal(b, m, deterministic)
}
func (dst *StudentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentInfo.Merge(dst, src)
}
func (m *StudentInfo) XXX_Size() int {
	return xxx_messageInfo_StudentInfo.Size(m)
}
func (m *StudentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StudentInfo proto.InternalMessageInfo

func (m *StudentInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type ReqStudent struct {
	StudentId            string   `protobuf:"bytes,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqStudent) Reset()         { *m = ReqStudent{} }
func (m *ReqStudent) String() string { return proto.CompactTextString(m) }
func (*ReqStudent) ProtoMessage()    {}
func (*ReqStudent) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{6}
}
func (m *ReqStudent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStudent.Unmarshal(m, b)
}
func (m *ReqStudent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStudent.Marshal(b, m, deterministic)
}
func (dst *ReqStudent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStudent.Merge(dst, src)
}
func (m *ReqStudent) XXX_Size() int {
	return xxx_messageInfo_ReqStudent.Size(m)
}
func (m *ReqStudent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStudent.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStudent proto.InternalMessageInfo

func (m *ReqStudent) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

type ResStudent struct {
	Student              *StudentInfo `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	Class                *ClassInfo   `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResStudent) Reset()         { *m = ResStudent{} }
func (m *ResStudent) String() string { return proto.CompactTextString(m) }
func (*ResStudent) ProtoMessage()    {}
func (*ResStudent) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_e2eff7549ffc3416, []int{7}
}
func (m *ResStudent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResStudent.Unmarshal(m, b)
}
func (m *ResStudent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResStudent.Marshal(b, m, deterministic)
}
func (dst *ResStudent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResStudent.Merge(dst, src)
}
func (m *ResStudent) XXX_Size() int {
	return xxx_messageInfo_ResStudent.Size(m)
}
func (m *ResStudent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResStudent.DiscardUnknown(m)
}

var xxx_messageInfo_ResStudent proto.InternalMessageInfo

func (m *ResStudent) GetStudent() *StudentInfo {
	if m != nil {
		return m.Student
	}
	return nil
}

func (m *ResStudent) GetClass() *ClassInfo {
	if m != nil {
		return m.Class
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqClassCreate)(nil), "admin.ReqClassCreate")
	proto.RegisterType((*ResSuccess)(nil), "admin.ResSuccess")
	proto.RegisterType((*ClassInfo)(nil), "admin.ClassInfo")
	proto.RegisterType((*ReqClass)(nil), "admin.ReqClass")
	proto.RegisterType((*ResClass)(nil), "admin.ResClass")
	proto.RegisterType((*StudentInfo)(nil), "admin.StudentInfo")
	proto.RegisterType((*ReqStudent)(nil), "admin.ReqStudent")
	proto.RegisterType((*ResStudent)(nil), "admin.ResStudent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClassClient is the client API for Class service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClassClient interface {
	Class(ctx context.Context, in *ReqClass, opts ...grpc.CallOption) (*ResClass, error)
	Create(ctx context.Context, in *ReqClassCreate, opts ...grpc.CallOption) (*ResSuccess, error)
}

type classClient struct {
	cc *grpc.ClientConn
}

func NewClassClient(cc *grpc.ClientConn) ClassClient {
	return &classClient{cc}
}

func (c *classClient) Class(ctx context.Context, in *ReqClass, opts ...grpc.CallOption) (*ResClass, error) {
	out := new(ResClass)
	err := c.cc.Invoke(ctx, "/admin.Class/Class", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classClient) Create(ctx context.Context, in *ReqClassCreate, opts ...grpc.CallOption) (*ResSuccess, error) {
	out := new(ResSuccess)
	err := c.cc.Invoke(ctx, "/admin.Class/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServer is the server API for Class service.
type ClassServer interface {
	Class(context.Context, *ReqClass) (*ResClass, error)
	Create(context.Context, *ReqClassCreate) (*ResSuccess, error)
}

func RegisterClassServer(s *grpc.Server, srv ClassServer) {
	s.RegisterService(&_Class_serviceDesc, srv)
}

func _Class_Class_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqClass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServer).Class(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Class/Class",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServer).Class(ctx, req.(*ReqClass))
	}
	return interceptor(ctx, in, info, handler)
}

func _Class_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqClassCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Class/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServer).Create(ctx, req.(*ReqClassCreate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Class_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Class",
	HandlerType: (*ClassServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Class",
			Handler:    _Class_Class_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Class_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentClient interface {
	Student(ctx context.Context, in *ReqStudent, opts ...grpc.CallOption) (*ResStudent, error)
}

type studentClient struct {
	cc *grpc.ClientConn
}

func NewStudentClient(cc *grpc.ClientConn) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) Student(ctx context.Context, in *ReqStudent, opts ...grpc.CallOption) (*ResStudent, error) {
	out := new(ResStudent)
	err := c.cc.Invoke(ctx, "/admin.Student/Student", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServer is the server API for Student service.
type StudentServer interface {
	Student(context.Context, *ReqStudent) (*ResStudent, error)
}

func RegisterStudentServer(s *grpc.Server, srv StudentServer) {
	s.RegisterService(&_Student_serviceDesc, srv)
}

func _Student_Student_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqStudent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).Student(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Student/Student",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).Student(ctx, req.(*ReqStudent))
	}
	return interceptor(ctx, in, info, handler)
}

var _Student_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Student",
			Handler:    _Student_Student_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor_admin_e2eff7549ffc3416) }

var fileDescriptor_admin_e2eff7549ffc3416 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4b, 0x6a, 0xe3, 0x40,
	0x10, 0x45, 0xf2, 0xc8, 0x9f, 0xd2, 0xa0, 0xb1, 0x9b, 0x99, 0x89, 0x10, 0x5e, 0x98, 0x26, 0x04,
	0x63, 0x82, 0x05, 0xf2, 0x26, 0x64, 0x93, 0x85, 0x17, 0x21, 0x64, 0x27, 0x6f, 0xb2, 0x6d, 0x5b,
	0x1d, 0xd1, 0x60, 0x77, 0xdb, 0x6e, 0x39, 0x07, 0xc8, 0x15, 0x42, 0x4e, 0x96, 0x2b, 0xe4, 0x20,
	0xa1, 0x7f, 0x52, 0x6c, 0x82, 0x77, 0x55, 0xaf, 0x9e, 0x5e, 0xd5, 0x7b, 0xb4, 0x20, 0x24, 0xc5,
	0x86, 0xf1, 0xe9, 0x76, 0x2f, 0x2a, 0x81, 0x02, 0xdd, 0x24, 0xc3, 0x52, 0x88, 0x72, 0x4d, 0x53,
	0xb2, 0x65, 0x29, 0xe1, 0x5c, 0x54, 0xa4, 0x62, 0x82, 0x4b, 0x43, 0xc2, 0x4f, 0x10, 0xe5, 0x74,
	0x37, 0x5f, 0x13, 0x29, 0xe7, 0x7b, 0x4a, 0x2a, 0x8a, 0x22, 0xf0, 0x59, 0x11, 0x7b, 0x23, 0x6f,
	0xdc, 0xcb, 0x7d, 0x56, 0xa0, 0x04, 0xba, 0xb2, 0x22, 0xbc, 0x20, 0xfb, 0x22, 0xf6, 0x35, 0x5a,
	0xf7, 0x6a, 0x56, 0xb0, 0x17, 0x26, 0x99, 0xe0, 0x71, 0xcb, 0xcc, 0x5c, 0x8f, 0x6f, 0x00, 0x72,
	0x2a, 0x17, 0x87, 0xd5, 0x8a, 0x4a, 0x89, 0x62, 0xe8, 0x48, 0x53, 0x6a, 0xe9, 0x6e, 0xee, 0x5a,
	0xd4, 0x87, 0xd6, 0x46, 0x96, 0x56, 0x5a, 0x95, 0x78, 0x0e, 0x3d, 0x7d, 0xd0, 0x03, 0x7f, 0x16,
	0x47, 0xeb, 0xbd, 0x33, 0xeb, 0xfd, 0x93, 0xf5, 0x97, 0xd0, 0x75, 0xc6, 0xd4, 0xf2, 0x95, 0x16,
	0x74, 0x12, 0xae, 0xc5, 0x4b, 0xc5, 0x92, 0x86, 0x75, 0x05, 0x81, 0x86, 0x35, 0x27, 0xcc, 0xfa,
	0x53, 0x13, 0x66, 0x7d, 0x4a, 0x6e, 0xc6, 0x68, 0xaa, 0x2e, 0x3a, 0x14, 0x94, 0x57, 0x32, 0xf6,
	0x47, 0xad, 0x71, 0x98, 0x21, 0x4b, 0x5d, 0x18, 0x58, 0x93, 0x6b, 0x0e, 0x9e, 0x41, 0xf8, 0x6d,
	0x80, 0x10, 0xfc, 0xe2, 0x64, 0x43, 0xed, 0x25, 0xba, 0x56, 0x19, 0x90, 0x92, 0x6a, 0x0f, 0x41,
	0xae, 0x4a, 0x3c, 0x51, 0xe9, 0xed, 0xec, 0x77, 0x68, 0x08, 0x3d, 0x2b, 0x57, 0x5b, 0x68, 0x00,
	0xbc, 0x34, 0x49, 0x5b, 0xee, 0x35, 0x74, 0xec, 0xc8, 0x1a, 0xf9, 0xe9, 0x3a, 0x47, 0x69, 0x4c,
	0xfb, 0x67, 0x4d, 0x67, 0xef, 0x1e, 0x04, 0x26, 0xa6, 0x3b, 0x57, 0xfc, 0xb1, 0x5c, 0x17, 0x73,
	0xd2, 0x00, 0x26, 0x51, 0xfc, 0xf7, 0xf5, 0xe3, 0xf3, 0xcd, 0x8f, 0xd0, 0xef, 0x54, 0x0f, 0x52,
	0x93, 0xdf, 0x23, 0xb4, 0xed, 0x53, 0xfb, 0x77, 0xa2, 0x60, 0xe0, 0x64, 0xd0, 0xe8, 0xd8, 0xe7,
	0x83, 0x2f, 0xb4, 0xd2, 0x00, 0x1f, 0x29, 0xdd, 0x7a, 0x93, 0x2c, 0x87, 0x8e, 0x33, 0x7e, 0xdf,
	0x94, 0x8d, 0x82, 0x8b, 0xf0, 0x48, 0xd4, 0x40, 0xf8, 0xbf, 0x16, 0xed, 0xa3, 0xc8, 0x8a, 0xda,
	0x4c, 0x96, 0x6d, 0xfd, 0x6b, 0xcc, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xab, 0xd1, 0x91,
	0x4e, 0x03, 0x00, 0x00,
}