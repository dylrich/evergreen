// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_results.proto

package internal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TestResultsInfo struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Variant              string   `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	TaskName             string   `protobuf:"bytes,4,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskId               string   `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Execution            int32    `protobuf:"varint,6,opt,name=execution,proto3" json:"execution,omitempty"`
	RequestType          string   `protobuf:"bytes,7,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	Mainline             bool     `protobuf:"varint,8,opt,name=mainline,proto3" json:"mainline,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResultsInfo) Reset()         { *m = TestResultsInfo{} }
func (m *TestResultsInfo) String() string { return proto.CompactTextString(m) }
func (*TestResultsInfo) ProtoMessage()    {}
func (*TestResultsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdf98c574c2f8ed6, []int{0}
}

func (m *TestResultsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultsInfo.Unmarshal(m, b)
}
func (m *TestResultsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultsInfo.Marshal(b, m, deterministic)
}
func (m *TestResultsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultsInfo.Merge(m, src)
}
func (m *TestResultsInfo) XXX_Size() int {
	return xxx_messageInfo_TestResultsInfo.Size(m)
}
func (m *TestResultsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultsInfo proto.InternalMessageInfo

func (m *TestResultsInfo) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *TestResultsInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TestResultsInfo) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *TestResultsInfo) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *TestResultsInfo) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TestResultsInfo) GetExecution() int32 {
	if m != nil {
		return m.Execution
	}
	return 0
}

func (m *TestResultsInfo) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *TestResultsInfo) GetMainline() bool {
	if m != nil {
		return m.Mainline
	}
	return false
}

type TestResults struct {
	TestResultsRecordId  string        `protobuf:"bytes,1,opt,name=test_results_record_id,json=testResultsRecordId,proto3" json:"test_results_record_id,omitempty"`
	Results              []*TestResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TestResults) Reset()         { *m = TestResults{} }
func (m *TestResults) String() string { return proto.CompactTextString(m) }
func (*TestResults) ProtoMessage()    {}
func (*TestResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdf98c574c2f8ed6, []int{1}
}

func (m *TestResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResults.Unmarshal(m, b)
}
func (m *TestResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResults.Marshal(b, m, deterministic)
}
func (m *TestResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResults.Merge(m, src)
}
func (m *TestResults) XXX_Size() int {
	return xxx_messageInfo_TestResults.Size(m)
}
func (m *TestResults) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResults.DiscardUnknown(m)
}

var xxx_messageInfo_TestResults proto.InternalMessageInfo

func (m *TestResults) GetTestResultsRecordId() string {
	if m != nil {
		return m.TestResultsRecordId
	}
	return ""
}

func (m *TestResults) GetResults() []*TestResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type TestResult struct {
	TestName             string               `protobuf:"bytes,1,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	Trial                int32                `protobuf:"varint,2,opt,name=trial,proto3" json:"trial,omitempty"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	LogUrl               string               `protobuf:"bytes,4,opt,name=log_url,json=logUrl,proto3" json:"log_url,omitempty"`
	LineNum              int32                `protobuf:"varint,5,opt,name=line_num,json=lineNum,proto3" json:"line_num,omitempty"`
	TaskCreateTime       *timestamp.Timestamp `protobuf:"bytes,6,opt,name=task_create_time,json=taskCreateTime,proto3" json:"task_create_time,omitempty"`
	TestStartTime        *timestamp.Timestamp `protobuf:"bytes,7,opt,name=test_start_time,json=testStartTime,proto3" json:"test_start_time,omitempty"`
	TestEndTime          *timestamp.Timestamp `protobuf:"bytes,8,opt,name=test_end_time,json=testEndTime,proto3" json:"test_end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestResult) Reset()         { *m = TestResult{} }
func (m *TestResult) String() string { return proto.CompactTextString(m) }
func (*TestResult) ProtoMessage()    {}
func (*TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdf98c574c2f8ed6, []int{2}
}

func (m *TestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResult.Unmarshal(m, b)
}
func (m *TestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResult.Marshal(b, m, deterministic)
}
func (m *TestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResult.Merge(m, src)
}
func (m *TestResult) XXX_Size() int {
	return xxx_messageInfo_TestResult.Size(m)
}
func (m *TestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResult.DiscardUnknown(m)
}

var xxx_messageInfo_TestResult proto.InternalMessageInfo

func (m *TestResult) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *TestResult) GetTrial() int32 {
	if m != nil {
		return m.Trial
	}
	return 0
}

func (m *TestResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TestResult) GetLogUrl() string {
	if m != nil {
		return m.LogUrl
	}
	return ""
}

func (m *TestResult) GetLineNum() int32 {
	if m != nil {
		return m.LineNum
	}
	return 0
}

func (m *TestResult) GetTaskCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.TaskCreateTime
	}
	return nil
}

func (m *TestResult) GetTestStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.TestStartTime
	}
	return nil
}

func (m *TestResult) GetTestEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.TestEndTime
	}
	return nil
}

type TestResultsEndInfo struct {
	TestResultsRecordId  string   `protobuf:"bytes,1,opt,name=test_results_record_id,json=testResultsRecordId,proto3" json:"test_results_record_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResultsEndInfo) Reset()         { *m = TestResultsEndInfo{} }
func (m *TestResultsEndInfo) String() string { return proto.CompactTextString(m) }
func (*TestResultsEndInfo) ProtoMessage()    {}
func (*TestResultsEndInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdf98c574c2f8ed6, []int{3}
}

func (m *TestResultsEndInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultsEndInfo.Unmarshal(m, b)
}
func (m *TestResultsEndInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultsEndInfo.Marshal(b, m, deterministic)
}
func (m *TestResultsEndInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultsEndInfo.Merge(m, src)
}
func (m *TestResultsEndInfo) XXX_Size() int {
	return xxx_messageInfo_TestResultsEndInfo.Size(m)
}
func (m *TestResultsEndInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultsEndInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultsEndInfo proto.InternalMessageInfo

func (m *TestResultsEndInfo) GetTestResultsRecordId() string {
	if m != nil {
		return m.TestResultsRecordId
	}
	return ""
}

type TestResultsResponse struct {
	TestResultsRecordId  string   `protobuf:"bytes,1,opt,name=test_results_record_id,json=testResultsRecordId,proto3" json:"test_results_record_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResultsResponse) Reset()         { *m = TestResultsResponse{} }
func (m *TestResultsResponse) String() string { return proto.CompactTextString(m) }
func (*TestResultsResponse) ProtoMessage()    {}
func (*TestResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdf98c574c2f8ed6, []int{4}
}

func (m *TestResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultsResponse.Unmarshal(m, b)
}
func (m *TestResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultsResponse.Marshal(b, m, deterministic)
}
func (m *TestResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultsResponse.Merge(m, src)
}
func (m *TestResultsResponse) XXX_Size() int {
	return xxx_messageInfo_TestResultsResponse.Size(m)
}
func (m *TestResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultsResponse proto.InternalMessageInfo

func (m *TestResultsResponse) GetTestResultsRecordId() string {
	if m != nil {
		return m.TestResultsRecordId
	}
	return ""
}

func init() {
	proto.RegisterType((*TestResultsInfo)(nil), "cedar.TestResultsInfo")
	proto.RegisterType((*TestResults)(nil), "cedar.TestResults")
	proto.RegisterType((*TestResult)(nil), "cedar.TestResult")
	proto.RegisterType((*TestResultsEndInfo)(nil), "cedar.TestResultsEndInfo")
	proto.RegisterType((*TestResultsResponse)(nil), "cedar.TestResultsResponse")
}

func init() { proto.RegisterFile("test_results.proto", fileDescriptor_cdf98c574c2f8ed6) }

var fileDescriptor_cdf98c574c2f8ed6 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0x12, 0xbb, 0x13, 0xe8, 0xc7, 0x16, 0xa5, 0x6e, 0x40, 0x22, 0xe4, 0x14, 0x09,
	0xc9, 0x95, 0xda, 0x3b, 0x82, 0x96, 0x1c, 0x82, 0x44, 0x91, 0xdc, 0x70, 0xe1, 0x62, 0x6d, 0xe3,
	0x69, 0x64, 0xb0, 0x77, 0xcd, 0xee, 0x18, 0xd1, 0xdf, 0xc6, 0xef, 0xe0, 0x8f, 0xf0, 0x0b, 0xd0,
	0x8e, 0x9d, 0xc6, 0x22, 0xa8, 0x95, 0xca, 0x71, 0xde, 0xdb, 0xf7, 0xd6, 0xf3, 0xf6, 0x19, 0x04,
	0xa1, 0xa5, 0xc4, 0xa0, 0xad, 0x72, 0xb2, 0x51, 0x69, 0x34, 0x69, 0xd1, 0x5d, 0x60, 0x2a, 0xcd,
	0xf0, 0xc5, 0x52, 0xeb, 0x65, 0x8e, 0xc7, 0x0c, 0x5e, 0x55, 0xd7, 0xc7, 0x94, 0x15, 0x68, 0x49,
	0x16, 0x65, 0x7d, 0x6e, 0xfc, 0xdb, 0x83, 0xdd, 0x39, 0x5a, 0x8a, 0x6b, 0xf5, 0x4c, 0x5d, 0x6b,
	0x11, 0x82, 0x5f, 0x1a, 0xfd, 0x05, 0x17, 0x14, 0x7a, 0x23, 0x6f, 0xb2, 0x1d, 0xaf, 0x46, 0xc7,
	0x7c, 0x47, 0x63, 0x33, 0xad, 0xc2, 0x4e, 0xcd, 0x34, 0x23, 0x33, 0xd2, 0x64, 0x52, 0x51, 0xb8,
	0xd5, 0x30, 0xf5, 0x28, 0x9e, 0xc1, 0x36, 0x49, 0xfb, 0x35, 0x51, 0xb2, 0xc0, 0xf0, 0x11, 0x73,
	0x81, 0x03, 0x2e, 0x64, 0x81, 0xe2, 0x10, 0x7c, 0x26, 0xb3, 0x34, 0xec, 0x32, 0xd5, 0x73, 0xe3,
	0x2c, 0x15, 0xcf, 0x61, 0x1b, 0x7f, 0xe0, 0xa2, 0x22, 0x77, 0x57, 0x6f, 0xe4, 0x4d, 0xba, 0xf1,
	0x1a, 0x10, 0x2f, 0xe1, 0xb1, 0xc1, 0x6f, 0x95, 0x5b, 0x9b, 0x6e, 0x4a, 0x0c, 0x7d, 0xd6, 0xf6,
	0x1b, 0x6c, 0x7e, 0x53, 0xa2, 0x18, 0x42, 0x50, 0xc8, 0x4c, 0xe5, 0x99, 0xc2, 0x30, 0x18, 0x79,
	0x93, 0x20, 0xbe, 0x9d, 0xc7, 0x1a, 0xfa, 0xad, 0x9d, 0xc5, 0x29, 0x0c, 0xda, 0x09, 0x26, 0x06,
	0x17, 0xda, 0xa4, 0xee, 0x9b, 0xea, 0xf5, 0x0f, 0x68, 0x7d, 0x38, 0x66, 0x6e, 0x96, 0x8a, 0x57,
	0xe0, 0x37, 0xe7, 0xc3, 0xce, 0x68, 0x6b, 0xd2, 0x3f, 0xd9, 0x8f, 0x38, 0xf2, 0x68, 0xed, 0x1c,
	0xaf, 0x4e, 0x8c, 0x7f, 0x75, 0x00, 0xd6, 0x38, 0x47, 0xe2, 0x2e, 0xe4, 0x48, 0xbc, 0x26, 0x12,
	0xb4, 0xc4, 0x91, 0x3c, 0x85, 0x2e, 0x99, 0x4c, 0xe6, 0x9c, 0x70, 0x37, 0xae, 0x07, 0x31, 0x80,
	0x9e, 0x25, 0x49, 0x95, 0x6d, 0xe2, 0x6d, 0x26, 0x17, 0x60, 0xae, 0x97, 0x49, 0x65, 0xf2, 0x26,
	0xdb, 0x5e, 0xae, 0x97, 0x9f, 0x4c, 0x2e, 0x8e, 0x20, 0x70, 0xbb, 0x26, 0xaa, 0x2a, 0x38, 0xda,
	0x6e, 0xec, 0xbb, 0xf9, 0xa2, 0x2a, 0xc4, 0x3b, 0xd8, 0xe3, 0xd0, 0x17, 0x06, 0x25, 0x61, 0xe2,
	0x2a, 0xc1, 0x11, 0xf7, 0x4f, 0x86, 0x51, 0xdd, 0x97, 0x68, 0xd5, 0x97, 0x68, 0xbe, 0xea, 0x4b,
	0xbc, 0xe3, 0x34, 0xe7, 0x2c, 0x71, 0xa0, 0x38, 0x83, 0x5d, 0x5e, 0xc2, 0x92, 0x34, 0x54, 0x9b,
	0xf8, 0xf7, 0x9a, 0x3c, 0x71, 0x92, 0x4b, 0xa7, 0x60, 0x8f, 0xd7, 0xc0, 0x40, 0x82, 0x2a, 0xad,
	0x1d, 0x82, 0x7b, 0x1d, 0xfa, 0x4e, 0x30, 0x55, 0xa9, 0x43, 0xc6, 0x33, 0x10, 0xad, 0x87, 0x9c,
	0xaa, 0x94, 0xfb, 0xfb, 0x90, 0xf7, 0x1c, 0xbf, 0x87, 0x83, 0x79, 0x1b, 0xb6, 0xa5, 0x56, 0x16,
	0x1f, 0xe4, 0x75, 0xf2, 0xb3, 0x03, 0x7b, 0xe7, 0xae, 0x0c, 0xed, 0x96, 0x7d, 0x80, 0xc3, 0x26,
	0xbd, 0xbf, 0x15, 0x62, 0xb0, 0x51, 0x1d, 0xfe, 0x11, 0x87, 0xc3, 0x4d, 0xfc, 0xf6, 0xc3, 0xde,
	0xc0, 0xce, 0xdb, 0x34, 0x6d, 0x5f, 0x20, 0x36, 0x4f, 0xdf, 0xe9, 0x30, 0x85, 0xfd, 0x4b, 0x32,
	0x28, 0x8b, 0xff, 0x30, 0x99, 0x78, 0xe2, 0x23, 0x0c, 0xce, 0x73, 0x6d, 0xff, 0xb1, 0xd6, 0xd1,
	0xa6, 0xae, 0x79, 0xa2, 0xbb, 0x2c, 0xcf, 0xe0, 0x73, 0x90, 0x29, 0x42, 0xa3, 0x64, 0x7e, 0xd5,
	0xe3, 0x06, 0x9c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x2f, 0x93, 0x35, 0xe3, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CedarTestResultsClient is the client API for CedarTestResults service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CedarTestResultsClient interface {
	CreateTestResultsRecord(ctx context.Context, in *TestResultsInfo, opts ...grpc.CallOption) (*TestResultsResponse, error)
	AddTestResults(ctx context.Context, in *TestResults, opts ...grpc.CallOption) (*TestResultsResponse, error)
	StreamTestResults(ctx context.Context, opts ...grpc.CallOption) (CedarTestResults_StreamTestResultsClient, error)
	CloseTestResultsRecord(ctx context.Context, in *TestResultsEndInfo, opts ...grpc.CallOption) (*TestResultsResponse, error)
}

type cedarTestResultsClient struct {
	cc grpc.ClientConnInterface
}

func NewCedarTestResultsClient(cc grpc.ClientConnInterface) CedarTestResultsClient {
	return &cedarTestResultsClient{cc}
}

func (c *cedarTestResultsClient) CreateTestResultsRecord(ctx context.Context, in *TestResultsInfo, opts ...grpc.CallOption) (*TestResultsResponse, error) {
	out := new(TestResultsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarTestResults/CreateTestResultsRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarTestResultsClient) AddTestResults(ctx context.Context, in *TestResults, opts ...grpc.CallOption) (*TestResultsResponse, error) {
	out := new(TestResultsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarTestResults/AddTestResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarTestResultsClient) StreamTestResults(ctx context.Context, opts ...grpc.CallOption) (CedarTestResults_StreamTestResultsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CedarTestResults_serviceDesc.Streams[0], "/cedar.CedarTestResults/StreamTestResults", opts...)
	if err != nil {
		return nil, err
	}
	x := &cedarTestResultsStreamTestResultsClient{stream}
	return x, nil
}

type CedarTestResults_StreamTestResultsClient interface {
	Send(*TestResults) error
	CloseAndRecv() (*TestResultsResponse, error)
	grpc.ClientStream
}

type cedarTestResultsStreamTestResultsClient struct {
	grpc.ClientStream
}

func (x *cedarTestResultsStreamTestResultsClient) Send(m *TestResults) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cedarTestResultsStreamTestResultsClient) CloseAndRecv() (*TestResultsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TestResultsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cedarTestResultsClient) CloseTestResultsRecord(ctx context.Context, in *TestResultsEndInfo, opts ...grpc.CallOption) (*TestResultsResponse, error) {
	out := new(TestResultsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarTestResults/CloseTestResultsRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CedarTestResultsServer is the server API for CedarTestResults service.
type CedarTestResultsServer interface {
	CreateTestResultsRecord(context.Context, *TestResultsInfo) (*TestResultsResponse, error)
	AddTestResults(context.Context, *TestResults) (*TestResultsResponse, error)
	StreamTestResults(CedarTestResults_StreamTestResultsServer) error
	CloseTestResultsRecord(context.Context, *TestResultsEndInfo) (*TestResultsResponse, error)
}

// UnimplementedCedarTestResultsServer can be embedded to have forward compatible implementations.
type UnimplementedCedarTestResultsServer struct {
}

func (*UnimplementedCedarTestResultsServer) CreateTestResultsRecord(ctx context.Context, req *TestResultsInfo) (*TestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestResultsRecord not implemented")
}
func (*UnimplementedCedarTestResultsServer) AddTestResults(ctx context.Context, req *TestResults) (*TestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTestResults not implemented")
}
func (*UnimplementedCedarTestResultsServer) StreamTestResults(srv CedarTestResults_StreamTestResultsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTestResults not implemented")
}
func (*UnimplementedCedarTestResultsServer) CloseTestResultsRecord(ctx context.Context, req *TestResultsEndInfo) (*TestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTestResultsRecord not implemented")
}

func RegisterCedarTestResultsServer(s *grpc.Server, srv CedarTestResultsServer) {
	s.RegisterService(&_CedarTestResults_serviceDesc, srv)
}

func _CedarTestResults_CreateTestResultsRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestResultsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarTestResultsServer).CreateTestResultsRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarTestResults/CreateTestResultsRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarTestResultsServer).CreateTestResultsRecord(ctx, req.(*TestResultsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarTestResults_AddTestResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestResults)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarTestResultsServer).AddTestResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarTestResults/AddTestResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarTestResultsServer).AddTestResults(ctx, req.(*TestResults))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarTestResults_StreamTestResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CedarTestResultsServer).StreamTestResults(&cedarTestResultsStreamTestResultsServer{stream})
}

type CedarTestResults_StreamTestResultsServer interface {
	SendAndClose(*TestResultsResponse) error
	Recv() (*TestResults, error)
	grpc.ServerStream
}

type cedarTestResultsStreamTestResultsServer struct {
	grpc.ServerStream
}

func (x *cedarTestResultsStreamTestResultsServer) SendAndClose(m *TestResultsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cedarTestResultsStreamTestResultsServer) Recv() (*TestResults, error) {
	m := new(TestResults)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CedarTestResults_CloseTestResultsRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestResultsEndInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarTestResultsServer).CloseTestResultsRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarTestResults/CloseTestResultsRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarTestResultsServer).CloseTestResultsRecord(ctx, req.(*TestResultsEndInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _CedarTestResults_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cedar.CedarTestResults",
	HandlerType: (*CedarTestResultsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTestResultsRecord",
			Handler:    _CedarTestResults_CreateTestResultsRecord_Handler,
		},
		{
			MethodName: "AddTestResults",
			Handler:    _CedarTestResults_AddTestResults_Handler,
		},
		{
			MethodName: "CloseTestResultsRecord",
			Handler:    _CedarTestResults_CloseTestResultsRecord_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTestResults",
			Handler:       _CedarTestResults_StreamTestResults_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test_results.proto",
}
