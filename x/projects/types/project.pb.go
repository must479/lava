// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/project.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProjectKey_KEY_TYPE int32

const (
	ProjectKey_NONE      ProjectKey_KEY_TYPE = 0
	ProjectKey_ADMIN     ProjectKey_KEY_TYPE = 1
	ProjectKey_DEVELOPER ProjectKey_KEY_TYPE = 2
)

var ProjectKey_KEY_TYPE_name = map[int32]string{
	0: "NONE",
	1: "ADMIN",
	2: "DEVELOPER",
}

var ProjectKey_KEY_TYPE_value = map[string]int32{
	"NONE":      0,
	"ADMIN":     1,
	"DEVELOPER": 2,
}

func (x ProjectKey_KEY_TYPE) String() string {
	return proto.EnumName(ProjectKey_KEY_TYPE_name, int32(x))
}

func (ProjectKey_KEY_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9f89a31663a330ce, []int{1, 0}
}

type Project struct {
	Index        string       `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Subscription string       `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Description  string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Enabled      bool         `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ProjectKeys  []ProjectKey `protobuf:"bytes,5,rep,name=project_keys,json=projectKeys,proto3" json:"project_keys"`
	Policy       Policy       `protobuf:"bytes,6,opt,name=policy,proto3" json:"policy"`
	UsedCu       uint64       `protobuf:"varint,7,opt,name=used_cu,json=usedCu,proto3" json:"used_cu,omitempty"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f89a31663a330ce, []int{0}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Project.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return m.Size()
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Project) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Project) GetProjectKeys() []ProjectKey {
	if m != nil {
		return m.ProjectKeys
	}
	return nil
}

func (m *Project) GetPolicy() Policy {
	if m != nil {
		return m.Policy
	}
	return Policy{}
}

func (m *Project) GetUsedCu() uint64 {
	if m != nil {
		return m.UsedCu
	}
	return 0
}

type ProjectKey struct {
	Key   string                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Types []ProjectKey_KEY_TYPE `protobuf:"varint,2,rep,packed,name=types,proto3,enum=lavanet.lava.projects.ProjectKey_KEY_TYPE" json:"types,omitempty"`
}

func (m *ProjectKey) Reset()         { *m = ProjectKey{} }
func (m *ProjectKey) String() string { return proto.CompactTextString(m) }
func (*ProjectKey) ProtoMessage()    {}
func (*ProjectKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f89a31663a330ce, []int{1}
}
func (m *ProjectKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectKey.Merge(m, src)
}
func (m *ProjectKey) XXX_Size() int {
	return m.Size()
}
func (m *ProjectKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectKey.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectKey proto.InternalMessageInfo

func (m *ProjectKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ProjectKey) GetTypes() []ProjectKey_KEY_TYPE {
	if m != nil {
		return m.Types
	}
	return nil
}

type Policy struct {
	ChainPolicies      []ChainPolicy `protobuf:"bytes,1,rep,name=chain_policies,json=chainPolicies,proto3" json:"chain_policies"`
	GeolocationProfile uint64        `protobuf:"varint,2,opt,name=geolocation_profile,json=geolocationProfile,proto3" json:"geolocation_profile,omitempty"`
	TotalCuLimit       uint64        `protobuf:"varint,3,opt,name=total_cu_limit,json=totalCuLimit,proto3" json:"total_cu_limit,omitempty"`
	EpochCuLimit       uint64        `protobuf:"varint,4,opt,name=epoch_cu_limit,json=epochCuLimit,proto3" json:"epoch_cu_limit,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f89a31663a330ce, []int{2}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetChainPolicies() []ChainPolicy {
	if m != nil {
		return m.ChainPolicies
	}
	return nil
}

func (m *Policy) GetGeolocationProfile() uint64 {
	if m != nil {
		return m.GeolocationProfile
	}
	return 0
}

func (m *Policy) GetTotalCuLimit() uint64 {
	if m != nil {
		return m.TotalCuLimit
	}
	return 0
}

func (m *Policy) GetEpochCuLimit() uint64 {
	if m != nil {
		return m.EpochCuLimit
	}
	return 0
}

type ChainPolicy struct {
	ChainId string   `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Apis    []string `protobuf:"bytes,2,rep,name=apis,proto3" json:"apis,omitempty"`
}

func (m *ChainPolicy) Reset()         { *m = ChainPolicy{} }
func (m *ChainPolicy) String() string { return proto.CompactTextString(m) }
func (*ChainPolicy) ProtoMessage()    {}
func (*ChainPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f89a31663a330ce, []int{3}
}
func (m *ChainPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainPolicy.Merge(m, src)
}
func (m *ChainPolicy) XXX_Size() int {
	return m.Size()
}
func (m *ChainPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ChainPolicy proto.InternalMessageInfo

func (m *ChainPolicy) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChainPolicy) GetApis() []string {
	if m != nil {
		return m.Apis
	}
	return nil
}

type ProtoString struct {
	String_ string `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
}

func (m *ProtoString) Reset()         { *m = ProtoString{} }
func (m *ProtoString) String() string { return proto.CompactTextString(m) }
func (*ProtoString) ProtoMessage()    {}
func (*ProtoString) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f89a31663a330ce, []int{4}
}
func (m *ProtoString) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtoString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtoString.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtoString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoString.Merge(m, src)
}
func (m *ProtoString) XXX_Size() int {
	return m.Size()
}
func (m *ProtoString) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoString.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoString proto.InternalMessageInfo

func (m *ProtoString) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

func init() {
	proto.RegisterEnum("lavanet.lava.projects.ProjectKey_KEY_TYPE", ProjectKey_KEY_TYPE_name, ProjectKey_KEY_TYPE_value)
	proto.RegisterType((*Project)(nil), "lavanet.lava.projects.Project")
	proto.RegisterType((*ProjectKey)(nil), "lavanet.lava.projects.ProjectKey")
	proto.RegisterType((*Policy)(nil), "lavanet.lava.projects.Policy")
	proto.RegisterType((*ChainPolicy)(nil), "lavanet.lava.projects.ChainPolicy")
	proto.RegisterType((*ProtoString)(nil), "lavanet.lava.projects.ProtoString")
}

func init() { proto.RegisterFile("projects/project.proto", fileDescriptor_9f89a31663a330ce) }

var fileDescriptor_9f89a31663a330ce = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x89, 0xe3, 0x24, 0xe3, 0x34, 0x8a, 0x96, 0x52, 0x0c, 0x12, 0xc6, 0x58, 0x20,
	0x59, 0x1c, 0x6c, 0x29, 0x1c, 0x39, 0x91, 0xd4, 0x48, 0xa5, 0x25, 0x89, 0x0c, 0x42, 0x2a, 0x17,
	0xcb, 0xb1, 0x97, 0x64, 0xa9, 0xeb, 0xb5, 0x62, 0x1b, 0xd5, 0x6f, 0xc1, 0x89, 0x3b, 0x6f, 0xd3,
	0x63, 0x4f, 0x88, 0x13, 0x42, 0xc9, 0x8b, 0x20, 0xaf, 0x37, 0x7f, 0x2a, 0x51, 0x71, 0xf2, 0xcc,
	0x37, 0xbf, 0x5d, 0xcf, 0x7c, 0xbb, 0x0b, 0x47, 0xc9, 0x92, 0x7d, 0x21, 0x41, 0x96, 0xda, 0x22,
	0xb0, 0x92, 0x25, 0xcb, 0x18, 0xbe, 0x1f, 0xf9, 0x5f, 0xfd, 0x98, 0x64, 0x56, 0xf9, 0xb5, 0x36,
	0xd0, 0xa3, 0xc3, 0x39, 0x9b, 0x33, 0x4e, 0xd8, 0x65, 0x54, 0xc1, 0xc6, 0x8f, 0x3a, 0xb4, 0xa6,
	0x15, 0x82, 0x0f, 0xa1, 0x49, 0xe3, 0x90, 0x5c, 0xa9, 0x48, 0x47, 0x66, 0xc7, 0xad, 0x12, 0x6c,
	0x40, 0x37, 0xcd, 0x67, 0x69, 0xb0, 0xa4, 0x49, 0x46, 0x59, 0xac, 0xd6, 0x79, 0xf1, 0x96, 0x86,
	0x75, 0x50, 0x42, 0xb2, 0x43, 0x1a, 0x1c, 0xd9, 0x97, 0xb0, 0x0a, 0x2d, 0x12, 0xfb, 0xb3, 0x88,
	0x84, 0xaa, 0xa4, 0x23, 0xb3, 0xed, 0x6e, 0x52, 0xfc, 0x16, 0xba, 0xa2, 0x47, 0xef, 0x82, 0x14,
	0xa9, 0xda, 0xd4, 0x1b, 0xa6, 0x32, 0x78, 0x6a, 0xfd, 0x73, 0x0a, 0x4b, 0xf4, 0x7a, 0x4a, 0x8a,
	0xa1, 0x74, 0xfd, 0xfb, 0x49, 0xcd, 0x55, 0x92, 0xad, 0x92, 0xe2, 0x57, 0x20, 0x27, 0x2c, 0xa2,
	0x41, 0xa1, 0xca, 0x3a, 0x32, 0x95, 0xc1, 0xe3, 0xbb, 0x76, 0xe1, 0x90, 0xd8, 0x41, 0x2c, 0xc1,
	0x0f, 0xa0, 0x95, 0xa7, 0x24, 0xf4, 0x82, 0x5c, 0x6d, 0xe9, 0xc8, 0x94, 0x5c, 0xb9, 0x4c, 0x47,
	0xb9, 0xf1, 0x1d, 0x01, 0xec, 0xfe, 0x8b, 0xfb, 0xd0, 0xb8, 0x20, 0x85, 0x30, 0xa9, 0x0c, 0xf1,
	0x1b, 0x68, 0x66, 0x45, 0x42, 0x52, 0xb5, 0xae, 0x37, 0xcc, 0xde, 0xe0, 0xc5, 0x7f, 0x7b, 0xb7,
	0x4e, 0x9d, 0x73, 0xef, 0xc3, 0xf9, 0xd4, 0x11, 0x2d, 0x54, 0xcb, 0x0d, 0x0b, 0xda, 0x9b, 0x02,
	0x6e, 0x83, 0x34, 0x9e, 0x8c, 0x9d, 0x7e, 0x0d, 0x77, 0xa0, 0xf9, 0xfa, 0xf8, 0xdd, 0xc9, 0xb8,
	0x8f, 0xf0, 0x01, 0x74, 0x8e, 0x9d, 0x8f, 0xce, 0xd9, 0x64, 0xea, 0xb8, 0xfd, 0xba, 0xf1, 0x13,
	0x81, 0x5c, 0x8d, 0x82, 0x27, 0xd0, 0x0b, 0x16, 0x3e, 0x8d, 0x3d, 0x3e, 0x0c, 0x25, 0xa9, 0x8a,
	0xb8, 0x8f, 0xc6, 0x1d, 0xbd, 0x8c, 0x4a, 0xf8, 0x96, 0x0d, 0x07, 0xc1, 0x56, 0xa2, 0x24, 0xc5,
	0x36, 0xdc, 0x9b, 0x13, 0x16, 0xb1, 0xc0, 0x2f, 0xcf, 0xcf, 0x4b, 0x96, 0xec, 0x33, 0x8d, 0x08,
	0x3f, 0x7d, 0xc9, 0xc5, 0x7b, 0xa5, 0x69, 0x55, 0xc1, 0xcf, 0xa0, 0x97, 0xb1, 0xcc, 0x8f, 0xbc,
	0x20, 0xf7, 0x22, 0x7a, 0x49, 0x33, 0x7e, 0x0d, 0x24, 0xb7, 0xcb, 0xd5, 0x51, 0x7e, 0x56, 0x6a,
	0x25, 0x45, 0x12, 0x16, 0x2c, 0x76, 0x94, 0x54, 0x51, 0x5c, 0x15, 0x94, 0x31, 0x04, 0x65, 0xaf,
	0x41, 0xfc, 0x10, 0xda, 0xd5, 0x70, 0x34, 0x14, 0xb6, 0xb7, 0x78, 0x7e, 0x12, 0x62, 0x15, 0x24,
	0x3f, 0xa1, 0x95, 0xf3, 0x1d, 0x31, 0x09, 0x57, 0x8c, 0xe7, 0xa0, 0x4c, 0xcb, 0x2b, 0xfe, 0x3e,
	0x5b, 0xd2, 0x78, 0x8e, 0x8f, 0x40, 0x4e, 0x79, 0x24, 0x76, 0x10, 0xd9, 0x70, 0x78, 0xbd, 0xd2,
	0xd0, 0xcd, 0x4a, 0x43, 0x7f, 0x56, 0x1a, 0xfa, 0xb6, 0xd6, 0x6a, 0x37, 0x6b, 0xad, 0xf6, 0x6b,
	0xad, 0xd5, 0x3e, 0x99, 0x73, 0x9a, 0x2d, 0xf2, 0x99, 0x15, 0xb0, 0x4b, 0x5b, 0x98, 0xc8, 0xbf,
	0xf6, 0x95, 0xbd, 0x7d, 0x79, 0xfc, 0xdc, 0x66, 0x32, 0x7f, 0x4b, 0x2f, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x81, 0xc4, 0xa7, 0x65, 0x92, 0x03, 0x00, 0x00,
}

func (m *Project) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Project) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Project) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UsedCu != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.UsedCu))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProject(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.ProjectKeys) > 0 {
		for iNdEx := len(m.ProjectKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Subscription) > 0 {
		i -= len(m.Subscription)
		copy(dAtA[i:], m.Subscription)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Subscription)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProjectKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Types) > 0 {
		dAtA3 := make([]byte, len(m.Types)*10)
		var j2 int
		for _, num := range m.Types {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintProject(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochCuLimit != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.EpochCuLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalCuLimit != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.TotalCuLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.GeolocationProfile != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.GeolocationProfile))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainPolicies) > 0 {
		for iNdEx := len(m.ChainPolicies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainPolicies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChainPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Apis) > 0 {
		for iNdEx := len(m.Apis) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Apis[iNdEx])
			copy(dAtA[i:], m.Apis[iNdEx])
			i = encodeVarintProject(dAtA, i, uint64(len(m.Apis[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintProject(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtoString) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtoString) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtoString) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.String_) > 0 {
		i -= len(m.String_)
		copy(dAtA[i:], m.String_)
		i = encodeVarintProject(dAtA, i, uint64(len(m.String_)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProject(dAtA []byte, offset int, v uint64) int {
	offset -= sovProject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Project) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Subscription)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.ProjectKeys) > 0 {
		for _, e := range m.ProjectKeys {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	l = m.Policy.Size()
	n += 1 + l + sovProject(uint64(l))
	if m.UsedCu != 0 {
		n += 1 + sovProject(uint64(m.UsedCu))
	}
	return n
}

func (m *ProjectKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if len(m.Types) > 0 {
		l = 0
		for _, e := range m.Types {
			l += sovProject(uint64(e))
		}
		n += 1 + sovProject(uint64(l)) + l
	}
	return n
}

func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainPolicies) > 0 {
		for _, e := range m.ChainPolicies {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if m.GeolocationProfile != 0 {
		n += 1 + sovProject(uint64(m.GeolocationProfile))
	}
	if m.TotalCuLimit != 0 {
		n += 1 + sovProject(uint64(m.TotalCuLimit))
	}
	if m.EpochCuLimit != 0 {
		n += 1 + sovProject(uint64(m.EpochCuLimit))
	}
	return n
}

func (m *ChainPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if len(m.Apis) > 0 {
		for _, s := range m.Apis {
			l = len(s)
			n += 1 + l + sovProject(uint64(l))
		}
	}
	return n
}

func (m *ProtoString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.String_)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	return n
}

func sovProject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProject(x uint64) (n int) {
	return sovProject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Project) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectKeys = append(m.ProjectKeys, ProjectKey{})
			if err := m.ProjectKeys[len(m.ProjectKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedCu", wireType)
			}
			m.UsedCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v ProjectKey_KEY_TYPE
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProject
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ProjectKey_KEY_TYPE(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Types = append(m.Types, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProject
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProject
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProject
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Types) == 0 {
					m.Types = make([]ProjectKey_KEY_TYPE, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ProjectKey_KEY_TYPE
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProject
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ProjectKey_KEY_TYPE(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Types = append(m.Types, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainPolicies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainPolicies = append(m.ChainPolicies, ChainPolicy{})
			if err := m.ChainPolicies[len(m.ChainPolicies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeolocationProfile", wireType)
			}
			m.GeolocationProfile = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GeolocationProfile |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCuLimit", wireType)
			}
			m.TotalCuLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCuLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCuLimit", wireType)
			}
			m.EpochCuLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochCuLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apis", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Apis = append(m.Apis, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtoString) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtoString: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtoString: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.String_ = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProject
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProject = fmt.Errorf("proto: unexpected end of group")
)