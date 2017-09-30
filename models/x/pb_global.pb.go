// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_global.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of PB_Chat from pb_table.proto

// Ignoring public import of PB_DirectMessage from pb_table.proto

// Ignoring public import of PB_MessageFile from pb_table.proto

// Ignoring public import of PB_User from pb_table.proto

type PB_CommandToServer struct {
	ClientCallId   int64  `protobuf:"varint,1,opt,name=ClientCallId" json:"ClientCallId,omitempty"`
	Command        string `protobuf:"bytes,2,opt,name=Command" json:"Command,omitempty"`
	RespondReached bool   `protobuf:"varint,3,opt,name=RespondReached" json:"RespondReached,omitempty"`
	Data           []byte `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *PB_CommandToServer) Reset()                    { *m = PB_CommandToServer{} }
func (m *PB_CommandToServer) String() string            { return proto.CompactTextString(m) }
func (*PB_CommandToServer) ProtoMessage()               {}
func (*PB_CommandToServer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PB_CommandToServer) GetClientCallId() int64 {
	if m != nil {
		return m.ClientCallId
	}
	return 0
}

func (m *PB_CommandToServer) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *PB_CommandToServer) GetRespondReached() bool {
	if m != nil {
		return m.RespondReached
	}
	return false
}

func (m *PB_CommandToServer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PB_CommandToClient struct {
	ServerCallId   int64  `protobuf:"varint,1,opt,name=ServerCallId" json:"ServerCallId,omitempty"`
	Command        string `protobuf:"bytes,2,opt,name=Command" json:"Command,omitempty"`
	RespondReached bool   `protobuf:"varint,3,opt,name=RespondReached" json:"RespondReached,omitempty"`
	Data           []byte `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *PB_CommandToClient) Reset()                    { *m = PB_CommandToClient{} }
func (m *PB_CommandToClient) String() string            { return proto.CompactTextString(m) }
func (*PB_CommandToClient) ProtoMessage()               {}
func (*PB_CommandToClient) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PB_CommandToClient) GetServerCallId() int64 {
	if m != nil {
		return m.ServerCallId
	}
	return 0
}

func (m *PB_CommandToClient) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *PB_CommandToClient) GetRespondReached() bool {
	if m != nil {
		return m.RespondReached
	}
	return false
}

func (m *PB_CommandToClient) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PB_CommandReachedToServer struct {
	ClientCallId int64 `protobuf:"varint,1,opt,name=ClientCallId" json:"ClientCallId,omitempty"`
}

func (m *PB_CommandReachedToServer) Reset()                    { *m = PB_CommandReachedToServer{} }
func (m *PB_CommandReachedToServer) String() string            { return proto.CompactTextString(m) }
func (*PB_CommandReachedToServer) ProtoMessage()               {}
func (*PB_CommandReachedToServer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *PB_CommandReachedToServer) GetClientCallId() int64 {
	if m != nil {
		return m.ClientCallId
	}
	return 0
}

type PB_CommandReachedToClient struct {
	ServerCallId int64 `protobuf:"varint,1,opt,name=ServerCallId" json:"ServerCallId,omitempty"`
}

func (m *PB_CommandReachedToClient) Reset()                    { *m = PB_CommandReachedToClient{} }
func (m *PB_CommandReachedToClient) String() string            { return proto.CompactTextString(m) }
func (*PB_CommandReachedToClient) ProtoMessage()               {}
func (*PB_CommandReachedToClient) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *PB_CommandReachedToClient) GetServerCallId() int64 {
	if m != nil {
		return m.ServerCallId
	}
	return 0
}

type PB_ResponseToClient struct {
	ClientCallId int64  `protobuf:"varint,1,opt,name=ClientCallId" json:"ClientCallId,omitempty"`
	PBClass      string `protobuf:"bytes,2,opt,name=PBClass" json:"PBClass,omitempty"`
	RpcFullName  string `protobuf:"bytes,3,opt,name=RpcFullName" json:"RpcFullName,omitempty"`
	Data         []byte `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *PB_ResponseToClient) Reset()                    { *m = PB_ResponseToClient{} }
func (m *PB_ResponseToClient) String() string            { return proto.CompactTextString(m) }
func (*PB_ResponseToClient) ProtoMessage()               {}
func (*PB_ResponseToClient) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *PB_ResponseToClient) GetClientCallId() int64 {
	if m != nil {
		return m.ClientCallId
	}
	return 0
}

func (m *PB_ResponseToClient) GetPBClass() string {
	if m != nil {
		return m.PBClass
	}
	return ""
}

func (m *PB_ResponseToClient) GetRpcFullName() string {
	if m != nil {
		return m.RpcFullName
	}
	return ""
}

func (m *PB_ResponseToClient) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_CommandToServer)(nil), "PB_CommandToServer")
	proto.RegisterType((*PB_CommandToClient)(nil), "PB_CommandToClient")
	proto.RegisterType((*PB_CommandReachedToServer)(nil), "PB_CommandReachedToServer")
	proto.RegisterType((*PB_CommandReachedToClient)(nil), "PB_CommandReachedToClient")
	proto.RegisterType((*PB_ResponseToClient)(nil), "PB_ResponseToClient")
}

func init() { proto.RegisterFile("pb_global.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0x8a, 0x4f,
	0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x2b, 0x48, 0x8a,
	0x2f, 0x49, 0x4c, 0xca, 0x49, 0x85, 0xf0, 0x95, 0xfa, 0x18, 0xb9, 0x84, 0x02, 0x9c, 0xe2, 0x9d,
	0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x42, 0xf2, 0x83, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x84, 0x94,
	0xb8, 0x78, 0x9c, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x9c, 0x13, 0x73, 0x72, 0x3c, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x98, 0x83, 0x50, 0xc4, 0x84, 0x24, 0xb8, 0xd8, 0xa1, 0xda, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x35, 0x2e, 0xbe, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc,
	0x94, 0xa0, 0xd4, 0xc4, 0xe4, 0x8c, 0xd4, 0x14, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x34,
	0x51, 0x21, 0x21, 0x2e, 0x16, 0x97, 0xc4, 0x92, 0x44, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x9e, 0x20,
	0x30, 0x1b, 0xc3, 0x41, 0x10, 0x2b, 0x41, 0x0e, 0x82, 0x38, 0x0d, 0xd5, 0x41, 0xc8, 0x62, 0x34,
	0x72, 0x90, 0x3d, 0x97, 0x24, 0xc2, 0x3d, 0x50, 0x85, 0xa4, 0x84, 0x13, 0x0e, 0x03, 0x88, 0xf7,
	0x97, 0x52, 0x27, 0x23, 0x97, 0x70, 0x80, 0x53, 0x3c, 0xc4, 0xad, 0xc5, 0xa9, 0xc8, 0x7a, 0x89,
	0x89, 0xa4, 0x00, 0x27, 0xe7, 0x9c, 0xc4, 0xe2, 0x62, 0x58, 0x98, 0x40, 0xb9, 0x42, 0x0a, 0x5c,
	0xdc, 0x41, 0x05, 0xc9, 0x6e, 0xa5, 0x39, 0x39, 0x7e, 0x89, 0xb9, 0xa9, 0xe0, 0x00, 0xe1, 0x0c,
	0x42, 0x16, 0xc2, 0x16, 0x1a, 0x4e, 0x82, 0x5c, 0x1c, 0x99, 0x45, 0x7a, 0xb9, 0xc5, 0x7a, 0x05,
	0x49, 0x1e, 0xcc, 0x01, 0x8c, 0x51, 0x8c, 0x15, 0x01, 0x0c, 0x49, 0x6c, 0xe0, 0xb4, 0x64, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xd7, 0xb0, 0x0a, 0x6e, 0x02, 0x00, 0x00,
}
