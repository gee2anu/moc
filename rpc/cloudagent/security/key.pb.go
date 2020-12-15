// Code generated by protoc-gen-go. DO NOT EDIT.
// source: key.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type KeyRequest struct {
	Keys                 []*Key           `protobuf:"bytes,1,rep,name=Keys,proto3" json:"Keys,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2accf5df6005b7d4, []int{0}
}

func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRequest.Unmarshal(m, b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
}
func (m *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(m, src)
}
func (m *KeyRequest) XXX_Size() int {
	return xxx_messageInfo_KeyRequest.Size(m)
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *KeyRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type KeyResponse struct {
	Keys                 []*Key              `protobuf:"bytes,1,rep,name=Keys,proto3" json:"Keys,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KeyResponse) Reset()         { *m = KeyResponse{} }
func (m *KeyResponse) String() string { return proto.CompactTextString(m) }
func (*KeyResponse) ProtoMessage()    {}
func (*KeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2accf5df6005b7d4, []int{1}
}

func (m *KeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyResponse.Unmarshal(m, b)
}
func (m *KeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyResponse.Marshal(b, m, deterministic)
}
func (m *KeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyResponse.Merge(m, src)
}
func (m *KeyResponse) XXX_Size() int {
	return xxx_messageInfo_KeyResponse.Size(m)
}
func (m *KeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyResponse proto.InternalMessageInfo

func (m *KeyResponse) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *KeyResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *KeyResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type KeyOperationRequest struct {
	Key                  *Key                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 string              `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Algorithm            common.Algorithm    `protobuf:"varint,3,opt,name=algorithm,proto3,enum=moc.Algorithm" json:"algorithm,omitempty"`
	OperationType        common.KeyOperation `protobuf:"varint,4,opt,name=OperationType,proto3,enum=moc.KeyOperation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KeyOperationRequest) Reset()         { *m = KeyOperationRequest{} }
func (m *KeyOperationRequest) String() string { return proto.CompactTextString(m) }
func (*KeyOperationRequest) ProtoMessage()    {}
func (*KeyOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2accf5df6005b7d4, []int{2}
}

func (m *KeyOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyOperationRequest.Unmarshal(m, b)
}
func (m *KeyOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyOperationRequest.Marshal(b, m, deterministic)
}
func (m *KeyOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyOperationRequest.Merge(m, src)
}
func (m *KeyOperationRequest) XXX_Size() int {
	return xxx_messageInfo_KeyOperationRequest.Size(m)
}
func (m *KeyOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyOperationRequest proto.InternalMessageInfo

func (m *KeyOperationRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyOperationRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *KeyOperationRequest) GetAlgorithm() common.Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return common.Algorithm_A_UNKNOWN
}

func (m *KeyOperationRequest) GetOperationType() common.KeyOperation {
	if m != nil {
		return m.OperationType
	}
	return common.KeyOperation_ENCRYPT
}

type KeyOperationResponse struct {
	Data                 string              `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KeyOperationResponse) Reset()         { *m = KeyOperationResponse{} }
func (m *KeyOperationResponse) String() string { return proto.CompactTextString(m) }
func (*KeyOperationResponse) ProtoMessage()    {}
func (*KeyOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2accf5df6005b7d4, []int{3}
}

func (m *KeyOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyOperationResponse.Unmarshal(m, b)
}
func (m *KeyOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyOperationResponse.Marshal(b, m, deterministic)
}
func (m *KeyOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyOperationResponse.Merge(m, src)
}
func (m *KeyOperationResponse) XXX_Size() int {
	return xxx_messageInfo_KeyOperationResponse.Size(m)
}
func (m *KeyOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyOperationResponse proto.InternalMessageInfo

func (m *KeyOperationResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *KeyOperationResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *KeyOperationResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Key struct {
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	LocationName string `protobuf:"bytes,3,opt,name=locationName,proto3" json:"locationName,omitempty"`
	// Public Key Value
	PublicKey            []byte                     `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Type                 common.JsonWebKeyType      `protobuf:"varint,5,opt,name=type,proto3,enum=moc.JsonWebKeyType" json:"type,omitempty"`
	VaultName            string                     `protobuf:"bytes,6,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	GroupName            string                     `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Status               *common.Status             `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Size                 common.KeySize             `protobuf:"varint,9,opt,name=size,proto3,enum=moc.KeySize" json:"size,omitempty"`
	Curve                common.JsonWebKeyCurveName `protobuf:"varint,10,opt,name=curve,proto3,enum=moc.JsonWebKeyCurveName" json:"curve,omitempty"`
	KeyOps               []common.KeyOperation      `protobuf:"varint,11,rep,packed,name=keyOps,proto3,enum=moc.KeyOperation" json:"keyOps,omitempty"`
	Tags                 *common.Tags               `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_2accf5df6005b7d4, []int{4}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Key) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Key) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *Key) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Key) GetType() common.JsonWebKeyType {
	if m != nil {
		return m.Type
	}
	return common.JsonWebKeyType_EC
}

func (m *Key) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *Key) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Key) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Key) GetSize() common.KeySize {
	if m != nil {
		return m.Size
	}
	return common.KeySize_K_UNKNOWN
}

func (m *Key) GetCurve() common.JsonWebKeyCurveName {
	if m != nil {
		return m.Curve
	}
	return common.JsonWebKeyCurveName_P_256
}

func (m *Key) GetKeyOps() []common.KeyOperation {
	if m != nil {
		return m.KeyOps
	}
	return nil
}

func (m *Key) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRequest)(nil), "moc.cloudagent.security.KeyRequest")
	proto.RegisterType((*KeyResponse)(nil), "moc.cloudagent.security.KeyResponse")
	proto.RegisterType((*KeyOperationRequest)(nil), "moc.cloudagent.security.KeyOperationRequest")
	proto.RegisterType((*KeyOperationResponse)(nil), "moc.cloudagent.security.KeyOperationResponse")
	proto.RegisterType((*Key)(nil), "moc.cloudagent.security.Key")
}

func init() { proto.RegisterFile("key.proto", fileDescriptor_2accf5df6005b7d4) }

var fileDescriptor_2accf5df6005b7d4 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xb9, 0x49, 0xd3, 0x7a, 0x92, 0x5f, 0x24, 0xb6, 0x95, 0xb0, 0xa2, 0x82, 0xa2, 0x14,
	0x89, 0x22, 0x15, 0xbb, 0x0a, 0x48, 0x9c, 0x5b, 0xe0, 0x00, 0x96, 0x40, 0xda, 0x16, 0x90, 0xb8,
	0x39, 0xee, 0xd4, 0xb5, 0x62, 0x7b, 0xcd, 0xfe, 0x09, 0x72, 0xbf, 0x01, 0x47, 0xbe, 0x11, 0x77,
	0xbe, 0x14, 0xda, 0xb1, 0xdd, 0x90, 0x08, 0x22, 0x84, 0xb8, 0x79, 0x67, 0xde, 0x9b, 0xf7, 0xf6,
	0xed, 0x18, 0xdc, 0x39, 0x56, 0x7e, 0x29, 0x85, 0x16, 0xec, 0x6e, 0x2e, 0x62, 0x3f, 0xce, 0x84,
	0xb9, 0x8c, 0x12, 0x2c, 0xb4, 0xaf, 0x30, 0x36, 0x32, 0xd5, 0xd5, 0xe8, 0x7e, 0x22, 0x44, 0x92,
	0x61, 0x40, 0xb0, 0x99, 0xb9, 0x0a, 0x3e, 0xcb, 0xa8, 0x2c, 0x51, 0xaa, 0x9a, 0x38, 0x1a, 0xc4,
	0x22, 0xcf, 0x45, 0xd1, 0x9c, 0x86, 0x2d, 0xaf, 0x3e, 0x4f, 0x34, 0x40, 0x88, 0x15, 0xc7, 0x4f,
	0x06, 0x95, 0x66, 0x27, 0xd0, 0x0d, 0xb1, 0x52, 0x9e, 0x33, 0xee, 0x1c, 0xf5, 0xa7, 0x07, 0xfe,
	0x6f, 0x34, 0x7d, 0x4b, 0x21, 0x24, 0x7b, 0x0a, 0xff, 0xbf, 0x2d, 0x51, 0x46, 0x3a, 0x15, 0xc5,
	0x45, 0x55, 0xa2, 0xb7, 0x35, 0x76, 0x8e, 0x86, 0xd3, 0x21, 0x51, 0x6f, 0x3b, 0x7c, 0x15, 0x34,
	0xf9, 0xe2, 0x40, 0x9f, 0x64, 0x55, 0x29, 0x0a, 0x85, 0x7f, 0xa1, 0x3b, 0x85, 0x1e, 0x47, 0x65,
	0x32, 0x4d, 0x82, 0xfd, 0xe9, 0xc8, 0xaf, 0x63, 0xf0, 0xdb, 0x18, 0xfc, 0x33, 0x21, 0xb2, 0xf7,
	0x51, 0x66, 0x90, 0x37, 0x48, 0xb6, 0x0f, 0xdb, 0x2f, 0xa5, 0x14, 0xd2, 0xeb, 0x8c, 0x9d, 0x23,
	0x97, 0xd7, 0x87, 0xc9, 0x37, 0x07, 0xf6, 0x42, 0xac, 0x96, 0x5e, 0x9b, 0x2c, 0x7c, 0xe8, 0xcc,
	0xb1, 0xf2, 0x1c, 0x1a, 0xbf, 0xd9, 0x92, 0x05, 0x32, 0x06, 0xdd, 0x17, 0x91, 0x8e, 0xc8, 0x8f,
	0xcb, 0xe9, 0x9b, 0x1d, 0x83, 0x1b, 0x65, 0x89, 0x90, 0xa9, 0xbe, 0xce, 0x49, 0xb5, 0x4d, 0xe6,
	0xb4, 0xad, 0xf2, 0x25, 0x80, 0x3d, 0x5b, 0xcf, 0xb2, 0x4b, 0x8c, 0x3b, 0xc4, 0x58, 0xb1, 0xb8,
	0x16, 0xa7, 0x86, 0xfd, 0xd5, 0x1b, 0x34, 0xb1, 0xb6, 0x96, 0x9c, 0x9f, 0x2c, 0xfd, 0xbb, 0xe0,
	0xbe, 0x76, 0xa0, 0x13, 0xd6, 0x17, 0x2f, 0xa2, 0x1c, 0x5b, 0x15, 0xfb, 0xcd, 0x86, 0xb0, 0x95,
	0x5e, 0x36, 0x51, 0x6c, 0xa5, 0x97, 0x6c, 0x02, 0x83, 0x4c, 0xc4, 0xe4, 0xee, 0x8d, 0xc5, 0xd6,
	0x83, 0x56, 0x6a, 0xec, 0x00, 0xdc, 0xd2, 0xcc, 0xb2, 0x34, 0x0e, 0xb1, 0xa2, 0xab, 0x0f, 0xf8,
	0xb2, 0xc0, 0x1e, 0x42, 0x57, 0xdb, 0x4c, 0xb6, 0x29, 0x93, 0x3d, 0xca, 0xe4, 0xb5, 0x12, 0xc5,
	0x07, 0x9c, 0x85, 0x58, 0xd9, 0x18, 0x38, 0x01, 0xec, 0x98, 0x45, 0x64, 0x32, 0x4d, 0x3a, 0x3d,
	0xd2, 0x59, 0x16, 0x6c, 0x37, 0x91, 0xc2, 0x94, 0xd4, 0xdd, 0xa9, 0xbb, 0xb7, 0x05, 0x76, 0x08,
	0x3d, 0xa5, 0x23, 0x6d, 0x94, 0xb7, 0x4b, 0xe1, 0xf4, 0x49, 0xe6, 0x9c, 0x4a, 0xbc, 0x69, 0xb1,
	0x31, 0x74, 0x55, 0x7a, 0x83, 0x9e, 0x4b, 0x4e, 0x06, 0xed, 0xeb, 0x9c, 0xa7, 0x37, 0xc8, 0xa9,
	0xc3, 0x7c, 0xd8, 0x8e, 0x8d, 0x5c, 0xa0, 0x07, 0x04, 0xf1, 0xd6, 0xcc, 0x3e, 0xb7, 0x3d, 0xab,
	0xc7, 0x6b, 0x18, 0x7b, 0x04, 0xbd, 0xb9, 0x7d, 0x3f, 0xe5, 0xf5, 0xc7, 0x9d, 0x5f, 0xbf, 0x78,
	0x03, 0x60, 0xf7, 0xa0, 0xab, 0xa3, 0x44, 0x79, 0x03, 0xf2, 0xe7, 0x12, 0xf0, 0x22, 0x4a, 0x14,
	0xa7, 0xf2, 0xf4, 0xbb, 0x03, 0xbb, 0x21, 0x56, 0xa7, 0x76, 0x47, 0xd9, 0x3b, 0xe8, 0xbd, 0x2a,
	0x16, 0x62, 0x8e, 0xec, 0x70, 0xe3, 0xfa, 0xd6, 0x0b, 0x3f, 0x7a, 0xb0, 0x19, 0x54, 0xef, 0xd4,
	0xe4, 0x3f, 0x76, 0x0d, 0x3b, 0xb5, 0x2f, 0x64, 0xc7, 0x9b, 0x28, 0xeb, 0x7f, 0xd4, 0xe8, 0xf1,
	0x1f, 0xa2, 0x5b, 0xa5, 0xb3, 0xe9, 0xc7, 0x93, 0x24, 0xd5, 0xd7, 0x66, 0xe6, 0xc7, 0x22, 0x0f,
	0xf2, 0x34, 0x96, 0x42, 0x89, 0x2b, 0x1d, 0xe4, 0x22, 0x0e, 0x64, 0x19, 0x07, 0xcb, 0x51, 0x41,
	0x3b, 0x6a, 0xd6, 0xa3, 0x3d, 0x7e, 0xf2, 0x23, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xff, 0xa0, 0xbe,
	0x3b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyAgentClient is the client API for KeyAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyAgentClient interface {
	Invoke(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	Operate(ctx context.Context, in *KeyOperationRequest, opts ...grpc.CallOption) (*KeyOperationResponse, error)
}

type keyAgentClient struct {
	cc *grpc.ClientConn
}

func NewKeyAgentClient(cc *grpc.ClientConn) KeyAgentClient {
	return &keyAgentClient{cc}
}

func (c *keyAgentClient) Invoke(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.KeyAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyAgentClient) Operate(ctx context.Context, in *KeyOperationRequest, opts ...grpc.CallOption) (*KeyOperationResponse, error) {
	out := new(KeyOperationResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.KeyAgent/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyAgentServer is the server API for KeyAgent service.
type KeyAgentServer interface {
	Invoke(context.Context, *KeyRequest) (*KeyResponse, error)
	Operate(context.Context, *KeyOperationRequest) (*KeyOperationResponse, error)
}

// UnimplementedKeyAgentServer can be embedded to have forward compatible implementations.
type UnimplementedKeyAgentServer struct {
}

func (*UnimplementedKeyAgentServer) Invoke(ctx context.Context, req *KeyRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedKeyAgentServer) Operate(ctx context.Context, req *KeyOperationRequest) (*KeyOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
}

func RegisterKeyAgentServer(s *grpc.Server, srv KeyAgentServer) {
	s.RegisterService(&_KeyAgent_serviceDesc, srv)
}

func _KeyAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.KeyAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyAgentServer).Invoke(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyAgent_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyAgentServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.KeyAgent/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyAgentServer).Operate(ctx, req.(*KeyOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.KeyAgent",
	HandlerType: (*KeyAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _KeyAgent_Invoke_Handler,
		},
		{
			MethodName: "Operate",
			Handler:    _KeyAgent_Operate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "key.proto",
}