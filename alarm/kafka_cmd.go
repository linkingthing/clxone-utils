package alarm

const (
	ThresholdTopic        = "threshold"
	ThresholdIpamTopic    = "threshold_ipam"
	ThresholdDnsTopic     = "threshold_dns"
	ThresholdDhcpTopic    = "threshold_dhcp"
	ThresholdMonitorTopic = "threshold_monitor"
	ThresholdHaTopic      = "threshold_ha"

	RegisterThreshold   = "register_threshold"
	UpdateThreshold     = "update_threshold"
	DeRegisterThreshold = "de_register_threshold"
)

const (
	AlarmTopic = "alarm"

	CmdCpuUsageAlarm        = "cpu_usage_alarm"
	CmdMemoryUsageAlarm     = "memory_usage_alarm"
	CmdStorageUsageAlarm    = "storage_usage_alarm"
	CmdSubnetRadioAlarm     = "subnet_radio_alarm"
	CmdLpsAlarm             = "lps_alarm"
	CmdQpsAlarm             = "qps_alarm"
	CmdHaTriggerAlarm       = "ha_trigger_alarm"
	CmdNodeOfflineAlarm     = "node_offline_alarm"
	CmdServiceOfflineAlarm  = "service_offline_alarm"
	CmdIllegalDhcpAlarm     = "illegal_dhcp_alarm"
	CmdConflictSubnetAlarm  = "conflict_subnet_alarm"
	CmdConflictIpAlarm      = "conflict_ip_alarm"
	CmdIpMacObsoletedAlarm  = "ip_mac_obsoleted_alarm"
	CmdIpPortObsoletedAlarm = "ip_port_obsoleted_alarm"
	CmdIpUnmanagedAlarm     = "ip_unmanaged_alarm"
	CmdZombieIpAlarm        = "zombie_ip_alarm"
	CmdOnlineExpiredIpAlarm = "online_expired_ip_alarm"
)
