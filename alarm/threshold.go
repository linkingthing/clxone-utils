package alarm

import (
	pb "github.com/linkingthing/clxone-utils/alarm/proto"
)

type Threshold struct {
	*pb.BaseThreshold
	Enabled  bool
	SendMail bool
	Value    uint64
}

var baseThresholdMap = map[pb.ThresholdName]*Threshold{
	pb.ThresholdName_cpuUsedRatio: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_cpuUsedRatio,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_ratio,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_memoryUsedRatio: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_memoryUsedRatio,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_ratio,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_storageUsedRatio: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_storageUsedRatio,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_ratio,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_qps: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_qps,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_values,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_lps: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_lps,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_values,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_subnetUsedRatio: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_subnetUsedRatio,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_ratio,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_haTrigger: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_haTrigger,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_nodeOffline: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_nodeOffline,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_serviceOffline: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_serviceOffline,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_subnetConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_subnetConflict,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_illegalDhcp: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_illegalDhcp,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_ipMacObsoleted: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_ipMacObsoleted,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_ipPortObsoleted: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_ipPortObsoleted,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_ipUnmanaged: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_ipUnmanaged,
			Level: pb.ThresholdLevel_critical,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_zombieIp: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_zombieIp,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_values,
		},
		Enabled:  true,
		SendMail: false,
		Value:    168,
	},
	pb.ThresholdName_onlineExpiredIp: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_onlineExpiredIp,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_values,
		},
		Enabled:  true,
		SendMail: false,
		Value:    720,
	},
	pb.ThresholdName_reservedIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_reservedIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpExcludeIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpExcludeIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpDynamicIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpDynamicIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpDynamicMacIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpDynamicMacIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpReservationMacIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpReservationMacIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpReservedIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpReservedIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpReservationIpConflict: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpReservationIpConflict,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: true,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalPacketWithOpcodeAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalPacketWithOpcodeAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalClientAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalClientAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalClientWithHighQpsAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalClientWithHighQpsAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithUnexpectedMessageTypeAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithUnexpectedMessageTypeAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithUltraShortLeaseTimeAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithUltraShortLeaseTimeAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithUltraLongLeaseTimeAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithUltraLongLeaseTimeAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithInvalidServerIdAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithInvalidServerIdAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithUnexpectedServerIdAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithUnexpectedServerIdAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithForbiddenServerIdAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithForbiddenServerIdAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithMandatoryServerIdAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithMandatoryServerIdAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithInvalidClientIdAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithInvalidClientIdAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionWithEmptyClientIdAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionWithEmptyClientIdAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
	pb.ThresholdName_dhcpIllegalOptionsAlarm: {
		BaseThreshold: &pb.BaseThreshold{
			Name:  pb.ThresholdName_dhcpIllegalOptionsAlarm,
			Level: pb.ThresholdLevel_major,
			Type:  pb.ThresholdType_trigger,
		},
		Enabled:  true,
		SendMail: false,
		Value:    0,
	},
}
