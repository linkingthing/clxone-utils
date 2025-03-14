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

	CmdCpuUsageAlarm                                   = "cpu_usage_alarm"
	CmdMemoryUsageAlarm                                = "memory_usage_alarm"
	CmdStorageUsageAlarm                               = "storage_usage_alarm"
	CmdSubnetRadioAlarm                                = "subnet_radio_alarm"
	CmdLpsAlarm                                        = "lps_alarm"
	CmdQpsAlarm                                        = "qps_alarm"
	CmdHaTriggerAlarm                                  = "ha_trigger_alarm"
	CmdNodeOfflineAlarm                                = "node_offline_alarm"
	CmdServiceOfflineAlarm                             = "service_offline_alarm"
	CmdIllegalDhcpAlarm                                = "illegal_dhcp_alarm"
	CmdConflictSubnetAlarm                             = "conflict_subnet_alarm"
	CmdConflictIpAlarm                                 = "conflict_ip_alarm"
	CmdIpMacObsoletedAlarm                             = "ip_mac_obsoleted_alarm"
	CmdIpPortObsoletedAlarm                            = "ip_port_obsoleted_alarm"
	CmdIpUnmanagedAlarm                                = "ip_unmanaged_alarm"
	CmdZombieIpAlarm                                   = "zombie_ip_alarm"
	CmdOnlineExpiredIpAlarm                            = "online_expired_ip_alarm"
	CmdReservedIpConflictAlarm                         = "reserved_ip_conflict_alarm"
	CmdDhcpExcludeIpConflictAlarm                      = "dhcp_exclude_ip_conflict"
	CmdDhcpDynamicMacIpConflictAlarm                   = "dhcp_dynamic_mac_ip_conflict_alarm"
	CmdDhcpReservationMacIpConflictAlarm               = "dhcp_reservation_mac_ip_conflict_alarm"
	CmdDhcpDynamicIpConflictAlarm                      = "dhcp_dynamic_ip_conflict_alarm"
	CmdDhcpReservationIpConflictAlarm                  = "dhcp_reservation_ip_conflict_alarm"
	CmdDhcpReservedIpConflictAlarm                     = "dhcp_reserved_ip_conflict_alarm"
	CmdDhcpIllegalPacketWithOpcodeAlarm                = "dhcp_illegal_packet_with_opcode_alarm"
	CmdDhcpIllegalInformPacketWithoutSourceAddrAlarm   = "dhcp_illegal_inform_packet_without_source_addr_alarm"
	CmdDhcpIllegalClientAlarm                          = "dhcp_illegal_client_alarm"
	CmdDhcpIllegalClientWithHighQpsAlarm               = "dhcp_illegal_client_with_high_qps_alarm"
	CmdDhcpIllegalOptionWithUnexpectedMessageTypeAlarm = "dhcp_illegal_option_with_unexpected_message_type_alarm"
	CmdDhcpIllegalOptionWithUltraShortLeaseTimeAlarm   = "dhcp_illegal_option_with_ultra_short_lease_time_alarm"
	CmdDhcpIllegalOptionWithUltraLongLeaseTimeAlarm    = "dhcp_illegal_option_with_ultra_long_lease_time_alarm"
	CmdDhcpIllegalOptionWithInvalidServerIdAlarm       = "dhcp_illegal_option_with_invalid_server_id_alarm"
	CmdDhcpIllegalOptionWithUnexpectedServerIdAlarm    = "dhcp_illegal_option_with_unexpected_server_id_alarm"
	CmdDhcpIllegalOptionWithForbiddenServerIdAlarm     = "dhcp_illegal_option_with_forbidden_server_id_alarm"
	CmdDhcpIllegalOptionWithMandatoryServerIdAlarm     = "dhcp_illegal_option_with_mandatory_server_id_alarm"
	CmdDhcpIllegalOptionWithInvalidClientIdAlarm       = "dhcp_illegal_option_with_invalid_client_id_alarm"
	CmdDhcpIllegalOptionWithMandatoryClientIdAlarm     = "dhcp_illegal_option_with_mandatory_client_id_alarm"
	CmdDhcpIllegalOptionsAlarm                         = "dhcp_illegal_options_alarm"

	CmdDatabaseOfflineAlarm = "database_offline_alarm"
)
