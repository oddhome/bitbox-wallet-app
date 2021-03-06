// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ETHCoin int32

const (
	ETHCoin_ETH        ETHCoin = 0
	ETHCoin_RopstenETH ETHCoin = 1
	ETHCoin_RinkebyETH ETHCoin = 2
)

var ETHCoin_name = map[int32]string{
	0: "ETH",
	1: "RopstenETH",
	2: "RinkebyETH",
}

var ETHCoin_value = map[string]int32{
	"ETH":        0,
	"RopstenETH": 1,
	"RinkebyETH": 2,
}

func (x ETHCoin) String() string {
	return proto.EnumName(ETHCoin_name, int32(x))
}

func (ETHCoin) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{0}
}

type ETHPubRequest_OutputType int32

const (
	ETHPubRequest_ADDRESS ETHPubRequest_OutputType = 0
	ETHPubRequest_XPUB    ETHPubRequest_OutputType = 1
)

var ETHPubRequest_OutputType_name = map[int32]string{
	0: "ADDRESS",
	1: "XPUB",
}

var ETHPubRequest_OutputType_value = map[string]int32{
	"ADDRESS": 0,
	"XPUB":    1,
}

func (x ETHPubRequest_OutputType) String() string {
	return proto.EnumName(ETHPubRequest_OutputType_name, int32(x))
}

func (ETHPubRequest_OutputType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{0, 0}
}

type ETHPubRequest struct {
	Keypath              []uint32                 `protobuf:"varint,1,rep,packed,name=keypath,proto3" json:"keypath,omitempty"`
	Coin                 ETHCoin                  `protobuf:"varint,2,opt,name=coin,proto3,enum=ETHCoin" json:"coin,omitempty"`
	OutputType           ETHPubRequest_OutputType `protobuf:"varint,3,opt,name=output_type,json=outputType,proto3,enum=ETHPubRequest_OutputType" json:"output_type,omitempty"`
	Display              bool                     `protobuf:"varint,4,opt,name=display,proto3" json:"display,omitempty"`
	ContractAddress      []byte                   `protobuf:"bytes,5,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ETHPubRequest) Reset()         { *m = ETHPubRequest{} }
func (m *ETHPubRequest) String() string { return proto.CompactTextString(m) }
func (*ETHPubRequest) ProtoMessage()    {}
func (*ETHPubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{0}
}

func (m *ETHPubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHPubRequest.Unmarshal(m, b)
}
func (m *ETHPubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHPubRequest.Marshal(b, m, deterministic)
}
func (m *ETHPubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHPubRequest.Merge(m, src)
}
func (m *ETHPubRequest) XXX_Size() int {
	return xxx_messageInfo_ETHPubRequest.Size(m)
}
func (m *ETHPubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHPubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ETHPubRequest proto.InternalMessageInfo

func (m *ETHPubRequest) GetKeypath() []uint32 {
	if m != nil {
		return m.Keypath
	}
	return nil
}

func (m *ETHPubRequest) GetCoin() ETHCoin {
	if m != nil {
		return m.Coin
	}
	return ETHCoin_ETH
}

func (m *ETHPubRequest) GetOutputType() ETHPubRequest_OutputType {
	if m != nil {
		return m.OutputType
	}
	return ETHPubRequest_ADDRESS
}

func (m *ETHPubRequest) GetDisplay() bool {
	if m != nil {
		return m.Display
	}
	return false
}

func (m *ETHPubRequest) GetContractAddress() []byte {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

type ETHSignRequest struct {
	Coin                 ETHCoin  `protobuf:"varint,1,opt,name=coin,proto3,enum=ETHCoin" json:"coin,omitempty"`
	Keypath              []uint32 `protobuf:"varint,2,rep,packed,name=keypath,proto3" json:"keypath,omitempty"`
	Nonce                []byte   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasPrice             []byte   `protobuf:"bytes,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit             []byte   `protobuf:"bytes,5,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Recipient            []byte   `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Value                []byte   `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Data                 []byte   `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ETHSignRequest) Reset()         { *m = ETHSignRequest{} }
func (m *ETHSignRequest) String() string { return proto.CompactTextString(m) }
func (*ETHSignRequest) ProtoMessage()    {}
func (*ETHSignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{1}
}

func (m *ETHSignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHSignRequest.Unmarshal(m, b)
}
func (m *ETHSignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHSignRequest.Marshal(b, m, deterministic)
}
func (m *ETHSignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHSignRequest.Merge(m, src)
}
func (m *ETHSignRequest) XXX_Size() int {
	return xxx_messageInfo_ETHSignRequest.Size(m)
}
func (m *ETHSignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHSignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ETHSignRequest proto.InternalMessageInfo

func (m *ETHSignRequest) GetCoin() ETHCoin {
	if m != nil {
		return m.Coin
	}
	return ETHCoin_ETH
}

func (m *ETHSignRequest) GetKeypath() []uint32 {
	if m != nil {
		return m.Keypath
	}
	return nil
}

func (m *ETHSignRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *ETHSignRequest) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *ETHSignRequest) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *ETHSignRequest) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *ETHSignRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ETHSignRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ETHSignMessageRequest struct {
	Coin                 ETHCoin  `protobuf:"varint,1,opt,name=coin,proto3,enum=ETHCoin" json:"coin,omitempty"`
	Keypath              []uint32 `protobuf:"varint,2,rep,packed,name=keypath,proto3" json:"keypath,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ETHSignMessageRequest) Reset()         { *m = ETHSignMessageRequest{} }
func (m *ETHSignMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ETHSignMessageRequest) ProtoMessage()    {}
func (*ETHSignMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{2}
}

func (m *ETHSignMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHSignMessageRequest.Unmarshal(m, b)
}
func (m *ETHSignMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHSignMessageRequest.Marshal(b, m, deterministic)
}
func (m *ETHSignMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHSignMessageRequest.Merge(m, src)
}
func (m *ETHSignMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ETHSignMessageRequest.Size(m)
}
func (m *ETHSignMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHSignMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ETHSignMessageRequest proto.InternalMessageInfo

func (m *ETHSignMessageRequest) GetCoin() ETHCoin {
	if m != nil {
		return m.Coin
	}
	return ETHCoin_ETH
}

func (m *ETHSignMessageRequest) GetKeypath() []uint32 {
	if m != nil {
		return m.Keypath
	}
	return nil
}

func (m *ETHSignMessageRequest) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type ETHSignResponse struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ETHSignResponse) Reset()         { *m = ETHSignResponse{} }
func (m *ETHSignResponse) String() string { return proto.CompactTextString(m) }
func (*ETHSignResponse) ProtoMessage()    {}
func (*ETHSignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{3}
}

func (m *ETHSignResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHSignResponse.Unmarshal(m, b)
}
func (m *ETHSignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHSignResponse.Marshal(b, m, deterministic)
}
func (m *ETHSignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHSignResponse.Merge(m, src)
}
func (m *ETHSignResponse) XXX_Size() int {
	return xxx_messageInfo_ETHSignResponse.Size(m)
}
func (m *ETHSignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHSignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ETHSignResponse proto.InternalMessageInfo

func (m *ETHSignResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type ETHRequest struct {
	// Types that are valid to be assigned to Request:
	//	*ETHRequest_Pub
	//	*ETHRequest_Sign
	//	*ETHRequest_SignMsg
	Request              isETHRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ETHRequest) Reset()         { *m = ETHRequest{} }
func (m *ETHRequest) String() string { return proto.CompactTextString(m) }
func (*ETHRequest) ProtoMessage()    {}
func (*ETHRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{4}
}

func (m *ETHRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHRequest.Unmarshal(m, b)
}
func (m *ETHRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHRequest.Marshal(b, m, deterministic)
}
func (m *ETHRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHRequest.Merge(m, src)
}
func (m *ETHRequest) XXX_Size() int {
	return xxx_messageInfo_ETHRequest.Size(m)
}
func (m *ETHRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ETHRequest proto.InternalMessageInfo

type isETHRequest_Request interface {
	isETHRequest_Request()
}

type ETHRequest_Pub struct {
	Pub *ETHPubRequest `protobuf:"bytes,1,opt,name=pub,proto3,oneof"`
}

type ETHRequest_Sign struct {
	Sign *ETHSignRequest `protobuf:"bytes,2,opt,name=sign,proto3,oneof"`
}

type ETHRequest_SignMsg struct {
	SignMsg *ETHSignMessageRequest `protobuf:"bytes,3,opt,name=sign_msg,json=signMsg,proto3,oneof"`
}

func (*ETHRequest_Pub) isETHRequest_Request() {}

func (*ETHRequest_Sign) isETHRequest_Request() {}

func (*ETHRequest_SignMsg) isETHRequest_Request() {}

func (m *ETHRequest) GetRequest() isETHRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ETHRequest) GetPub() *ETHPubRequest {
	if x, ok := m.GetRequest().(*ETHRequest_Pub); ok {
		return x.Pub
	}
	return nil
}

func (m *ETHRequest) GetSign() *ETHSignRequest {
	if x, ok := m.GetRequest().(*ETHRequest_Sign); ok {
		return x.Sign
	}
	return nil
}

func (m *ETHRequest) GetSignMsg() *ETHSignMessageRequest {
	if x, ok := m.GetRequest().(*ETHRequest_SignMsg); ok {
		return x.SignMsg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ETHRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ETHRequest_Pub)(nil),
		(*ETHRequest_Sign)(nil),
		(*ETHRequest_SignMsg)(nil),
	}
}

type ETHResponse struct {
	// Types that are valid to be assigned to Response:
	//	*ETHResponse_Pub
	//	*ETHResponse_Sign
	Response             isETHResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ETHResponse) Reset()         { *m = ETHResponse{} }
func (m *ETHResponse) String() string { return proto.CompactTextString(m) }
func (*ETHResponse) ProtoMessage()    {}
func (*ETHResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d980f775760351f6, []int{5}
}

func (m *ETHResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHResponse.Unmarshal(m, b)
}
func (m *ETHResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHResponse.Marshal(b, m, deterministic)
}
func (m *ETHResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHResponse.Merge(m, src)
}
func (m *ETHResponse) XXX_Size() int {
	return xxx_messageInfo_ETHResponse.Size(m)
}
func (m *ETHResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ETHResponse proto.InternalMessageInfo

type isETHResponse_Response interface {
	isETHResponse_Response()
}

type ETHResponse_Pub struct {
	Pub *PubResponse `protobuf:"bytes,1,opt,name=pub,proto3,oneof"`
}

type ETHResponse_Sign struct {
	Sign *ETHSignResponse `protobuf:"bytes,2,opt,name=sign,proto3,oneof"`
}

func (*ETHResponse_Pub) isETHResponse_Response() {}

func (*ETHResponse_Sign) isETHResponse_Response() {}

func (m *ETHResponse) GetResponse() isETHResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ETHResponse) GetPub() *PubResponse {
	if x, ok := m.GetResponse().(*ETHResponse_Pub); ok {
		return x.Pub
	}
	return nil
}

func (m *ETHResponse) GetSign() *ETHSignResponse {
	if x, ok := m.GetResponse().(*ETHResponse_Sign); ok {
		return x.Sign
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ETHResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ETHResponse_Pub)(nil),
		(*ETHResponse_Sign)(nil),
	}
}

func init() {
	proto.RegisterEnum("ETHCoin", ETHCoin_name, ETHCoin_value)
	proto.RegisterEnum("ETHPubRequest_OutputType", ETHPubRequest_OutputType_name, ETHPubRequest_OutputType_value)
	proto.RegisterType((*ETHPubRequest)(nil), "ETHPubRequest")
	proto.RegisterType((*ETHSignRequest)(nil), "ETHSignRequest")
	proto.RegisterType((*ETHSignMessageRequest)(nil), "ETHSignMessageRequest")
	proto.RegisterType((*ETHSignResponse)(nil), "ETHSignResponse")
	proto.RegisterType((*ETHRequest)(nil), "ETHRequest")
	proto.RegisterType((*ETHResponse)(nil), "ETHResponse")
}

func init() { proto.RegisterFile("eth.proto", fileDescriptor_d980f775760351f6) }

var fileDescriptor_d980f775760351f6 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x6d, 0xda, 0x6e, 0x49, 0x6f, 0xb2, 0x36, 0xb2, 0x00, 0x05, 0xd8, 0x43, 0x15, 0x04, 0x2a,
	0x3c, 0x04, 0xa9, 0x7b, 0xe3, 0x6d, 0x65, 0x91, 0xf2, 0xc0, 0x44, 0xe5, 0x06, 0x89, 0xb7, 0xca,
	0x4d, 0xad, 0xcc, 0x5a, 0x6b, 0x9b, 0xd8, 0x41, 0xea, 0x97, 0xf0, 0x7b, 0xfc, 0x00, 0xff, 0x80,
	0xec, 0x24, 0xeb, 0x8a, 0xf6, 0xc6, 0x53, 0xee, 0x39, 0xe7, 0xc6, 0xf7, 0x9c, 0x1b, 0x07, 0x46,
	0x54, 0xdf, 0x25, 0xb2, 0x12, 0x5a, 0xbc, 0x0a, 0x0a, 0xb1, 0xdf, 0x0b, 0xde, 0xa0, 0xf8, 0x8f,
	0x03, 0x17, 0x69, 0x9e, 0x2d, 0xeb, 0x0d, 0xa6, 0x3f, 0x6a, 0xaa, 0x34, 0x8a, 0xc0, 0xbd, 0xa7,
	0x07, 0x49, 0xf4, 0x5d, 0xe4, 0x4c, 0x07, 0xb3, 0x0b, 0xdc, 0x41, 0x74, 0x09, 0xc3, 0x42, 0x30,
	0x1e, 0xf5, 0xa7, 0xce, 0x6c, 0x3c, 0xf7, 0x92, 0x34, 0xcf, 0x3e, 0x0b, 0xc6, 0xb1, 0x65, 0xd1,
	0x27, 0xf0, 0x45, 0xad, 0x65, 0xad, 0xd7, 0xfa, 0x20, 0x69, 0x34, 0xb0, 0x4d, 0x2f, 0x93, 0x93,
	0xc3, 0x93, 0xaf, 0xb6, 0x23, 0x3f, 0x48, 0x8a, 0x41, 0x3c, 0xd4, 0x66, 0xe6, 0x96, 0x29, 0xb9,
	0x23, 0x87, 0x68, 0x38, 0x75, 0x66, 0x1e, 0xee, 0x20, 0x7a, 0x0f, 0x61, 0x21, 0xb8, 0xae, 0x48,
	0xa1, 0xd7, 0x64, 0xbb, 0xad, 0xa8, 0x52, 0xd1, 0xd9, 0xd4, 0x99, 0x05, 0x78, 0xd2, 0xf1, 0xd7,
	0x0d, 0x1d, 0xbf, 0x01, 0x38, 0x1e, 0x8f, 0x7c, 0x70, 0xaf, 0x6f, 0x6e, 0x70, 0xba, 0x5a, 0x85,
	0x3d, 0xe4, 0xc1, 0xf0, 0xfb, 0xf2, 0xdb, 0x22, 0x74, 0xe2, 0xdf, 0x0e, 0x8c, 0xd3, 0x3c, 0x5b,
	0xb1, 0x92, 0x77, 0x81, 0xbb, 0x58, 0xce, 0x93, 0xb1, 0x1e, 0xad, 0xa3, 0x7f, 0xba, 0x8e, 0x67,
	0x70, 0xc6, 0x05, 0x2f, 0x9a, 0xa8, 0x01, 0x6e, 0x00, 0x7a, 0x0d, 0xa3, 0x92, 0xa8, 0xb5, 0xac,
	0x58, 0x41, 0x6d, 0x98, 0x00, 0x7b, 0x25, 0x51, 0x4b, 0x83, 0x3b, 0x71, 0xc7, 0xf6, 0x4c, 0xb7,
	0x31, 0x8c, 0xf8, 0xc5, 0x60, 0x74, 0x09, 0xa3, 0x8a, 0x16, 0x4c, 0x32, 0xca, 0x75, 0x74, 0x6e,
	0xc5, 0x23, 0x61, 0xa6, 0xfd, 0x24, 0xbb, 0x9a, 0x46, 0x6e, 0x33, 0xcd, 0x02, 0x84, 0x60, 0xb8,
	0x25, 0x9a, 0x44, 0x9e, 0x25, 0x6d, 0x1d, 0x13, 0x78, 0xde, 0x26, 0xbc, 0xa5, 0x4a, 0x91, 0x92,
	0xfe, 0x6f, 0xd0, 0x10, 0x06, 0x7b, 0x55, 0xb6, 0x31, 0x4d, 0x19, 0x7f, 0x84, 0xc9, 0xc3, 0x12,
	0x95, 0x14, 0x5c, 0x51, 0xe3, 0x5e, 0xb1, 0x92, 0x13, 0x5d, 0x57, 0xd4, 0x4e, 0x08, 0xf0, 0x91,
	0x88, 0x7f, 0x39, 0x00, 0x69, 0x9e, 0x75, 0x4e, 0x62, 0x18, 0xc8, 0x7a, 0x63, 0xdb, 0xfc, 0xf9,
	0xf8, 0xf4, 0x8e, 0x64, 0x3d, 0x6c, 0x44, 0xf4, 0x16, 0x86, 0xe6, 0x7d, 0x7b, 0xdb, 0xfc, 0xf9,
	0x24, 0x39, 0xfd, 0x6a, 0x59, 0x0f, 0x5b, 0x19, 0x5d, 0x81, 0x67, 0x9e, 0xeb, 0xce, 0xa1, 0x3f,
	0x7f, 0x91, 0x3c, 0x19, 0x3f, 0xeb, 0x61, 0xd7, 0x74, 0xde, 0xaa, 0x72, 0x31, 0x02, 0xb7, 0x6a,
	0xd8, 0xb8, 0x00, 0xdf, 0x1a, 0x6b, 0x63, 0x4c, 0x1f, 0x3b, 0x0b, 0x12, 0x6b, 0xab, 0x91, 0x3a,
	0x5f, 0xef, 0x4e, 0x7c, 0x85, 0xc9, 0x3f, 0x8b, 0xe8, 0x8c, 0x2d, 0x00, 0xbc, 0xaa, 0xe5, 0x3e,
	0xcc, 0xc1, 0x6d, 0x97, 0x8d, 0x5c, 0x18, 0xa4, 0x79, 0x16, 0xf6, 0xd0, 0x18, 0x00, 0x0b, 0xa9,
	0x34, 0xe5, 0x06, 0x3b, 0x16, 0x33, 0x7e, 0x4f, 0x37, 0x07, 0x83, 0xfb, 0x9b, 0x73, 0xfb, 0x83,
	0x5e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xf7, 0xd4, 0xfd, 0xbb, 0x03, 0x00, 0x00,
}
