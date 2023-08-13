// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: thresholds.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ThresholdType int32

const (
	ThresholdType_values  ThresholdType = 0
	ThresholdType_ratio   ThresholdType = 1
	ThresholdType_trigger ThresholdType = 2
)

// Enum value maps for ThresholdType.
var (
	ThresholdType_name = map[int32]string{
		0: "values",
		1: "ratio",
		2: "trigger",
	}
	ThresholdType_value = map[string]int32{
		"values":  0,
		"ratio":   1,
		"trigger": 2,
	}
)

func (x ThresholdType) Enum() *ThresholdType {
	p := new(ThresholdType)
	*p = x
	return p
}

func (x ThresholdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThresholdType) Descriptor() protoreflect.EnumDescriptor {
	return file_thresholds_proto_enumTypes[0].Descriptor()
}

func (ThresholdType) Type() protoreflect.EnumType {
	return &file_thresholds_proto_enumTypes[0]
}

func (x ThresholdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThresholdType.Descriptor instead.
func (ThresholdType) EnumDescriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{0}
}

type ThresholdLevel int32

const (
	ThresholdLevel_critical ThresholdLevel = 0
	ThresholdLevel_major    ThresholdLevel = 1
	ThresholdLevel_minor    ThresholdLevel = 2
	ThresholdLevel_warning  ThresholdLevel = 3
)

// Enum value maps for ThresholdLevel.
var (
	ThresholdLevel_name = map[int32]string{
		0: "critical",
		1: "major",
		2: "minor",
		3: "warning",
	}
	ThresholdLevel_value = map[string]int32{
		"critical": 0,
		"major":    1,
		"minor":    2,
		"warning":  3,
	}
)

func (x ThresholdLevel) Enum() *ThresholdLevel {
	p := new(ThresholdLevel)
	*p = x
	return p
}

func (x ThresholdLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThresholdLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_thresholds_proto_enumTypes[1].Descriptor()
}

func (ThresholdLevel) Type() protoreflect.EnumType {
	return &file_thresholds_proto_enumTypes[1]
}

func (x ThresholdLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThresholdLevel.Descriptor instead.
func (ThresholdLevel) EnumDescriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{1}
}

type ThresholdName int32

const (
	ThresholdName_cpuUsedRatio                 ThresholdName = 0
	ThresholdName_memoryUsedRatio              ThresholdName = 1
	ThresholdName_storageUsedRatio             ThresholdName = 2
	ThresholdName_qps                          ThresholdName = 3
	ThresholdName_lps                          ThresholdName = 4
	ThresholdName_subnetUsedRatio              ThresholdName = 5
	ThresholdName_haTrigger                    ThresholdName = 6
	ThresholdName_nodeOffline                  ThresholdName = 7
	ThresholdName_serviceOffline               ThresholdName = 8
	ThresholdName_subnetConflict               ThresholdName = 10
	ThresholdName_illegalDhcp                  ThresholdName = 11
	ThresholdName_ipMacObsoleted               ThresholdName = 12
	ThresholdName_ipPortObsoleted              ThresholdName = 13
	ThresholdName_ipUnmanaged                  ThresholdName = 14
	ThresholdName_zombieIp                     ThresholdName = 15
	ThresholdName_onlineExpiredIp              ThresholdName = 16
	ThresholdName_dhcpDynamicIpConflict        ThresholdName = 17
	ThresholdName_dhcpReservationIpConflict    ThresholdName = 18
	ThresholdName_dhcpReservedIpConflict       ThresholdName = 19
	ThresholdName_dhcpDynamicMacIpConflict     ThresholdName = 20
	ThresholdName_dhcpReservationMacIpConflict ThresholdName = 21
	ThresholdName_dhcpExcludeIpConflict        ThresholdName = 22
	ThresholdName_reservedIpConflict           ThresholdName = 23
	ThresholdName_addressAudit                 ThresholdName = 24
	ThresholdName_asAudit                      ThresholdName = 25
)

// Enum value maps for ThresholdName.
var (
	ThresholdName_name = map[int32]string{
		0:  "cpuUsedRatio",
		1:  "memoryUsedRatio",
		2:  "storageUsedRatio",
		3:  "qps",
		4:  "lps",
		5:  "subnetUsedRatio",
		6:  "haTrigger",
		7:  "nodeOffline",
		8:  "serviceOffline",
		10: "subnetConflict",
		11: "illegalDhcp",
		12: "ipMacObsoleted",
		13: "ipPortObsoleted",
		14: "ipUnmanaged",
		15: "zombieIp",
		16: "onlineExpiredIp",
		17: "dhcpDynamicIpConflict",
		18: "dhcpReservationIpConflict",
		19: "dhcpReservedIpConflict",
		20: "dhcpDynamicMacIpConflict",
		21: "dhcpReservationMacIpConflict",
		22: "dhcpExcludeIpConflict",
		23: "reservedIpConflict",
		24: "addressAudit",
		25: "asAudit",
	}
	ThresholdName_value = map[string]int32{
		"cpuUsedRatio":                 0,
		"memoryUsedRatio":              1,
		"storageUsedRatio":             2,
		"qps":                          3,
		"lps":                          4,
		"subnetUsedRatio":              5,
		"haTrigger":                    6,
		"nodeOffline":                  7,
		"serviceOffline":               8,
		"subnetConflict":               10,
		"illegalDhcp":                  11,
		"ipMacObsoleted":               12,
		"ipPortObsoleted":              13,
		"ipUnmanaged":                  14,
		"zombieIp":                     15,
		"onlineExpiredIp":              16,
		"dhcpDynamicIpConflict":        17,
		"dhcpReservationIpConflict":    18,
		"dhcpReservedIpConflict":       19,
		"dhcpDynamicMacIpConflict":     20,
		"dhcpReservationMacIpConflict": 21,
		"dhcpExcludeIpConflict":        22,
		"reservedIpConflict":           23,
		"addressAudit":                 24,
		"asAudit":                      25,
	}
)

func (x ThresholdName) Enum() *ThresholdName {
	p := new(ThresholdName)
	*p = x
	return p
}

func (x ThresholdName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThresholdName) Descriptor() protoreflect.EnumDescriptor {
	return file_thresholds_proto_enumTypes[2].Descriptor()
}

func (ThresholdName) Type() protoreflect.EnumType {
	return &file_thresholds_proto_enumTypes[2]
}

func (x ThresholdName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThresholdName.Descriptor instead.
func (ThresholdName) EnumDescriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{2}
}

type BaseThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  ThresholdName  `protobuf:"varint,1,opt,name=name,proto3,enum=ThresholdName" json:"name,omitempty"`
	Level ThresholdLevel `protobuf:"varint,2,opt,name=level,proto3,enum=ThresholdLevel" json:"level,omitempty"`
	Type  ThresholdType  `protobuf:"varint,3,opt,name=type,proto3,enum=ThresholdType" json:"type,omitempty"`
}

func (x *BaseThreshold) Reset() {
	*x = BaseThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thresholds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseThreshold) ProtoMessage() {}

func (x *BaseThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_thresholds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseThreshold.ProtoReflect.Descriptor instead.
func (*BaseThreshold) Descriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{0}
}

func (x *BaseThreshold) GetName() ThresholdName {
	if x != nil {
		return x.Name
	}
	return ThresholdName_cpuUsedRatio
}

func (x *BaseThreshold) GetLevel() ThresholdLevel {
	if x != nil {
		return x.Level
	}
	return ThresholdLevel_critical
}

func (x *BaseThreshold) GetType() ThresholdType {
	if x != nil {
		return x.Type
	}
	return ThresholdType_values
}

type RegisterThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseThreshold *BaseThreshold `protobuf:"bytes,1,opt,name=base_threshold,json=baseThreshold,proto3" json:"base_threshold,omitempty"`
	Value         uint64         `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	SendMail      bool           `protobuf:"varint,3,opt,name=send_mail,json=sendMail,proto3" json:"send_mail,omitempty"`
	Enabled       bool           `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *RegisterThreshold) Reset() {
	*x = RegisterThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thresholds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterThreshold) ProtoMessage() {}

func (x *RegisterThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_thresholds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterThreshold.ProtoReflect.Descriptor instead.
func (*RegisterThreshold) Descriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterThreshold) GetBaseThreshold() *BaseThreshold {
	if x != nil {
		return x.BaseThreshold
	}
	return nil
}

func (x *RegisterThreshold) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *RegisterThreshold) GetSendMail() bool {
	if x != nil {
		return x.SendMail
	}
	return false
}

func (x *RegisterThreshold) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     ThresholdName `protobuf:"varint,1,opt,name=name,proto3,enum=ThresholdName" json:"name,omitempty"`
	Value    uint64        `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	SendMail bool          `protobuf:"varint,3,opt,name=send_mail,json=sendMail,proto3" json:"send_mail,omitempty"`
	Enabled  bool          `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *UpdateThreshold) Reset() {
	*x = UpdateThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thresholds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateThreshold) ProtoMessage() {}

func (x *UpdateThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_thresholds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateThreshold.ProtoReflect.Descriptor instead.
func (*UpdateThreshold) Descriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateThreshold) GetName() ThresholdName {
	if x != nil {
		return x.Name
	}
	return ThresholdName_cpuUsedRatio
}

func (x *UpdateThreshold) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *UpdateThreshold) GetSendMail() bool {
	if x != nil {
		return x.SendMail
	}
	return false
}

func (x *UpdateThreshold) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type DeRegisterThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name ThresholdName `protobuf:"varint,1,opt,name=name,proto3,enum=ThresholdName" json:"name,omitempty"`
}

func (x *DeRegisterThreshold) Reset() {
	*x = DeRegisterThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thresholds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeRegisterThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeRegisterThreshold) ProtoMessage() {}

func (x *DeRegisterThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_thresholds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeRegisterThreshold.ProtoReflect.Descriptor instead.
func (*DeRegisterThreshold) Descriptor() ([]byte, []int) {
	return file_thresholds_proto_rawDescGZIP(), []int{3}
}

func (x *DeRegisterThreshold) GetName() ThresholdName {
	if x != nil {
		return x.Name
	}
	return ThresholdName_cpuUsedRatio
}

var File_thresholds_proto protoreflect.FileDescriptor

var file_thresholds_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x82, 0x01, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x39, 0x0a, 0x13, 0x44, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x33, 0x0a, 0x0d,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x10,
	0x02, 0x2a, 0x41, 0x0a, 0x0e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x10, 0x03, 0x2a, 0x91, 0x04, 0x0a, 0x0d, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x63, 0x70, 0x75, 0x55, 0x73, 0x65,
	0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x55, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x71, 0x70, 0x73, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x6c, 0x70, 0x73, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x68, 0x61,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x6e, 0x6f, 0x64,
	0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x08, 0x12, 0x12,
	0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74,
	0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x69, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x44, 0x68, 0x63,
	0x70, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x69, 0x70, 0x4d, 0x61, 0x63, 0x4f, 0x62, 0x73, 0x6f,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x69, 0x70, 0x50, 0x6f, 0x72,
	0x74, 0x4f, 0x62, 0x73, 0x6f, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b,
	0x69, 0x70, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x10, 0x0e, 0x12, 0x0c, 0x0a,
	0x08, 0x7a, 0x6f, 0x6d, 0x62, 0x69, 0x65, 0x49, 0x70, 0x10, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x49, 0x70, 0x10, 0x10,
	0x12, 0x19, 0x0a, 0x15, 0x64, 0x68, 0x63, 0x70, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x64,
	0x68, 0x63, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x70,
	0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0x12, 0x12, 0x1a, 0x0a, 0x16, 0x64, 0x68,
	0x63, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x6c, 0x69, 0x63, 0x74, 0x10, 0x13, 0x12, 0x1c, 0x0a, 0x18, 0x64, 0x68, 0x63, 0x70, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x61, 0x63, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69,
	0x63, 0x74, 0x10, 0x14, 0x12, 0x20, 0x0a, 0x1c, 0x64, 0x68, 0x63, 0x70, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x63, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x6c, 0x69, 0x63, 0x74, 0x10, 0x15, 0x12, 0x19, 0x0a, 0x15, 0x64, 0x68, 0x63, 0x70, 0x45, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10,
	0x16, 0x12, 0x16, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x49, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0x17, 0x12, 0x10, 0x0a, 0x0c, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x10, 0x18, 0x12, 0x0b, 0x0a, 0x07, 0x61,
	0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x10, 0x19, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thresholds_proto_rawDescOnce sync.Once
	file_thresholds_proto_rawDescData = file_thresholds_proto_rawDesc
)

func file_thresholds_proto_rawDescGZIP() []byte {
	file_thresholds_proto_rawDescOnce.Do(func() {
		file_thresholds_proto_rawDescData = protoimpl.X.CompressGZIP(file_thresholds_proto_rawDescData)
	})
	return file_thresholds_proto_rawDescData
}

var file_thresholds_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_thresholds_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_thresholds_proto_goTypes = []interface{}{
	(ThresholdType)(0),          // 0: ThresholdType
	(ThresholdLevel)(0),         // 1: ThresholdLevel
	(ThresholdName)(0),          // 2: ThresholdName
	(*BaseThreshold)(nil),       // 3: BaseThreshold
	(*RegisterThreshold)(nil),   // 4: RegisterThreshold
	(*UpdateThreshold)(nil),     // 5: UpdateThreshold
	(*DeRegisterThreshold)(nil), // 6: DeRegisterThreshold
}
var file_thresholds_proto_depIdxs = []int32{
	2, // 0: BaseThreshold.name:type_name -> ThresholdName
	1, // 1: BaseThreshold.level:type_name -> ThresholdLevel
	0, // 2: BaseThreshold.type:type_name -> ThresholdType
	3, // 3: RegisterThreshold.base_threshold:type_name -> BaseThreshold
	2, // 4: UpdateThreshold.name:type_name -> ThresholdName
	2, // 5: DeRegisterThreshold.name:type_name -> ThresholdName
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_thresholds_proto_init() }
func file_thresholds_proto_init() {
	if File_thresholds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thresholds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseThreshold); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thresholds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterThreshold); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thresholds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateThreshold); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thresholds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeRegisterThreshold); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thresholds_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thresholds_proto_goTypes,
		DependencyIndexes: file_thresholds_proto_depIdxs,
		EnumInfos:         file_thresholds_proto_enumTypes,
		MessageInfos:      file_thresholds_proto_msgTypes,
	}.Build()
	File_thresholds_proto = out.File
	file_thresholds_proto_rawDesc = nil
	file_thresholds_proto_goTypes = nil
	file_thresholds_proto_depIdxs = nil
}
