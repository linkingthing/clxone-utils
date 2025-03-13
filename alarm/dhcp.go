package alarm

import (
	"bytes"
	"strconv"

	pb "github.com/linkingthing/clxone-utils/alarm/proto"
)

type DHCPClient4 struct {
	HwAddress             string
	HwAddressOrganization string
	Hostname              string
	ClientId              string
	Fingerprint           string
	VendorId              string
	OperatingSystem       string
	ClientType            string
	MessageType           string
}

type DHCPClient6 struct {
	Duid                  string
	HwAddress             string
	HwAddressOrganization string
	Hostname              string
	ClientId              string
	Fingerprint           string
	VendorId              string
	OperatingSystem       string
	ClientType            string
	RequestSourceAddr     string
	MessageType           string
}

// illegal_packet_with_opcode_alarm
func (a *Alarm) AddDhcpIllegalPacketWithOpcodeAlarm(mac, hostname, messageType, opcode string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalPacketWithOpcodeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalPacketWithOpcodeMessageEn(mac, hostname, messageType, opcode),
		GenDhcpIllegalPacketWithOpcodeMessageCh(mac, hostname, messageType, opcode),
		CmdDhcpIllegalPacketWithOpcodeAlarm)
}
func GenDhcpIllegalPacketWithOpcodeMessageCh(mac, hostname, messageType, opcode string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端MAC：")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 的非法报文")
	buf.WriteString(messageType)
	buf.WriteString("，")
	if len(opcode) > 0 {
		buf.WriteString("不支持操作码 ")
		buf.WriteString(opcode)
	} else {
		buf.WriteString("请求源地址为空")
	}
	return buf.String()
}
func GenDhcpIllegalPacketWithOpcodeMessageEn(mac, hostname, messageType, opcode string) string {
	buf := bytes.Buffer{}
	buf.WriteString("an illegal message of ")
	buf.WriteString(messageType)
	buf.WriteString("was received from the client with Mac address")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(".")
	if len(opcode) > 0 {
		buf.WriteString(". The operation code ")
		buf.WriteString(opcode)
		buf.WriteString(" is not supported.")
	} else {
		buf.WriteString("The source address of the request is empty")
	}

	return buf.String()
}

// illegal_client_alarm
func (a *Alarm) AddDhcpIllegalClientAlarm(mac, duid, hostname string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalClientAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalClientAlarmMessageEn(mac, duid, hostname),
		GenDhcpIllegalClientAlarmMessageCh(mac, duid, hostname),
		CmdDhcpIllegalClientAlarm)
}
func GenDhcpIllegalClientAlarmMessageCh(mac, duid, hostname string) string {
	buf := bytes.Buffer{}
	buf.WriteString("客户端 ")
	if mac != "" {
		buf.WriteString("Mac：")
		buf.WriteString(mac)
	}
	if duid != "" {
		buf.WriteString("DUID：")
		buf.WriteString(duid)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 非法接入")
	return buf.String()
}
func GenDhcpIllegalClientAlarmMessageEn(mac, duid, hostname string) string {
	buf := bytes.Buffer{}
	buf.WriteString("client with ")
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString("has illegally accessed the system")
	return buf.String()
}

// illegal_client_with_high_qps_alarm
func (a *Alarm) AddDhcpIllegalClientWithHighQpsAlarm(mac, duid, hostname, rateLimit string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalClientWithHighQpsAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalClientWithHighQpsAlarmMessageEn(mac, duid, hostname, rateLimit),
		GenDhcpIllegalClientWithHighQpsAlarmMessageCh(mac, duid, hostname, rateLimit),
		CmdDhcpIllegalClientWithHighQpsAlarm)
}
func GenDhcpIllegalClientWithHighQpsAlarmMessageCh(mac, duid, hostname, rateLimit string) string {
	buf := bytes.Buffer{}
	buf.WriteString("客户端 ")
	if mac != "" {
		buf.WriteString("Mac：")
		buf.WriteString(mac)
	}
	if duid != "" {
		buf.WriteString("DUID：")
		buf.WriteString(duid)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 请求速度超过阈值 ")
	buf.WriteString(rateLimit)
	return buf.String()
}
func GenDhcpIllegalClientWithHighQpsAlarmMessageEn(mac, duid, hostname, rateLimit string) string {
	buf := bytes.Buffer{}
	buf.WriteString("client with ")
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" has exceeded the rate limit")
	buf.WriteString(rateLimit)
	return buf.String()
}

// illegal_option_with_unexpected_message_type_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUnexpectedMessageTypeAlarm(mac, hostname, messageType, value string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUnexpectedMessageTypeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageEn(mac, hostname, messageType, value),
		GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageCh(mac, hostname, messageType, value),
		CmdDhcpIllegalOptionWithUnexpectedMessageTypeAlarm)
}
func GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageCh(mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端Mac：")
	buf.WriteString(mac)

	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文")
	buf.WriteString(messageType)
	buf.WriteString("携带非法OPTION 53，不支持报文类型")
	buf.WriteString(value)
	return buf.String()
}
func GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageEn(mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with Mac address ")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION 53. The message type ")
	buf.WriteString(value)
	buf.WriteString(" is not supported")
	return buf.String()
}

// illegal_option_with_ultra_short_lease_time_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUltraShortLeaseTimeAlarm(mac, hostname, messageType, value string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUltraShortLeaseTimeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageEn(mac, hostname, messageType, value),
		GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageCh(mac, hostname, messageType, value),
		CmdDhcpIllegalOptionWithUltraShortLeaseTimeAlarm)
}
func GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageCh(mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端Mac：")
	buf.WriteString(mac)

	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文")
	buf.WriteString(messageType)
	buf.WriteString("携带非法OPTION 51，请求租约时间")
	buf.WriteString(value)
	buf.WriteString(" 过短")
	return buf.String()
}
func GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageEn(mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with Mac address ")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION 51. The requested lease time ")
	buf.WriteString(value)
	buf.WriteString(" is too short")
	return buf.String()
}

// illegal_option_with_ultra_long_lease_time_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUltraLongLeaseTimeAlarm(mac, hostname, messageType, value string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUltraLongLeaseTimeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageEn(mac, hostname, messageType, value),
		GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageCh(mac, hostname, messageType, value),
		CmdDhcpIllegalOptionWithUltraLongLeaseTimeAlarm)
}
func GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageCh(mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端Mac：")
	buf.WriteString(mac)

	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文")
	buf.WriteString(messageType)
	buf.WriteString("携带非法OPTION 51，请求租约时间")
	buf.WriteString(value)
	buf.WriteString(" 过长")
	return buf.String()
}
func GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageEn(mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with Mac address ")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION 51. The requested lease time ")
	buf.WriteString(value)
	buf.WriteString(" is too long")
	return buf.String()
}

// illegal_option_with_invalid_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithInvalidServerIdAlarm(duid, mac, hostname, messageType, value string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithInvalidServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageEn(duid, mac, hostname, messageType, value),
		GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageCh(duid, mac, hostname, messageType, value),
		CmdDhcpIllegalOptionWithInvalidServerIdAlarm)
}
func GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageCh(duid, mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端DUID：")
	buf.WriteString(duid)
	if mac != "" {
		buf.WriteString(" Mac：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文")
	buf.WriteString(messageType)
	buf.WriteString("携带非法OPTION 2，server id ")
	buf.WriteString(value)
	buf.WriteString(" 格式错误")
	return buf.String()
}
func GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageEn(duid, mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with DUID ")
	buf.WriteString(duid)
	if mac != "" {
		buf.WriteString(", Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(", hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION 2. The format of the server ID ")
	buf.WriteString(value)
	buf.WriteString(" is incorrect")
	return buf.String()
}

// illegal_option_with_unexpected_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUnexpectedServerIdAlarm(duid, mac, hostname, messageType, value string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUnexpectedServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageEn(duid, mac, hostname, messageType, value),
		GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageCh(duid, mac, hostname, messageType, value),
		CmdDhcpIllegalOptionWithUnexpectedServerIdAlarm)
}
func GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageCh(duid, mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文")
	buf.WriteString(messageType)
	buf.WriteString("携带非法OPTION ")
	if mac != "" {
		buf.WriteString("54")
	}
	if duid != "" {
		buf.WriteString("2")
	}
	buf.WriteString("，server id ")
	buf.WriteString(value)
	buf.WriteString(" 与当前服务器地址不匹配")
	return buf.String()
}
func GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageEn(duid, mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION ")
	if mac != "" {
		buf.WriteString("54")
	}
	if duid != "" {
		buf.WriteString("2")
	}
	buf.WriteString(". The server ID ")
	buf.WriteString(value)
	buf.WriteString(" does not match the current server address")
	return buf.String()
}

// illegal_option_with_forbidden_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithForbiddenServerIdAlarm(duid, mac, hostname, messageType, value string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithForbiddenServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageEn(duid, mac, hostname, messageType, value),
		GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageCh(duid, mac, hostname, messageType, value),
		CmdDhcpIllegalOptionWithForbiddenServerIdAlarm)
}
func GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageCh(duid, mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文")
	buf.WriteString(messageType)
	buf.WriteString("携带非法OPTION ")
	if mac != "" {
		buf.WriteString("54")
	}
	if duid != "" {
		buf.WriteString("2")
	}
	buf.WriteString("，禁止携带 server id ")
	buf.WriteString(value)
	return buf.String()
}
func GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageEn(duid, mac, hostname, messageType, value string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION ")
	if mac != "" {
		buf.WriteString("54")
	}
	if duid != "" {
		buf.WriteString("2")
	}
	buf.WriteString(". The server ID ")
	buf.WriteString(value)
	buf.WriteString(" is not allowed")
	return buf.String()
}

// illegal_option_with_mandatory_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithMandatoryServerIdAlarm(duid, mac, hostname, messageType string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithMandatoryServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageEn(duid, mac, hostname, messageType),
		GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageCh(duid, mac, hostname, messageType),
		CmdDhcpIllegalOptionWithMandatoryServerIdAlarm)
}
func GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageCh(duid, mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 未携带OPTION ")
	if mac != "" {
		buf.WriteString("54")
	}
	if duid != "" {
		buf.WriteString("2")
	}
	return buf.String()
}
func GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageEn(duid, mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , The message does not carry OPTION ")
	if mac != "" {
		buf.WriteString("54")
	}
	if duid != "" {
		buf.WriteString("2")
	}

	return buf.String()
}

// illegal_option_with_invalid_client_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithInvalidClientIdAlarm(duid, mac, hostname, messageType, clientId string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithInvalidClientIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageEn(duid, mac, hostname, messageType, clientId),
		GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageCh(duid, mac, hostname, messageType, clientId),
		CmdDhcpIllegalOptionWithInvalidClientIdAlarm)
}
func GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageCh(duid, mac, hostname, messageType, clientId string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION 1，client id ")
	buf.WriteString(clientId)
	buf.WriteString(" 格式错误")
	return buf.String()
}
func GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageEn(duid, mac, hostname, messageType, clientId string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" ,The message carries an illegal OPTION 1, The format of the client ID")
	buf.WriteString(clientId)
	buf.WriteString(" is incorrect")
	return buf.String()
}

// illegal_option_with_empty_client_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithEmptyClientIdAlarm(mac, hostname, messageType string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithEmptyClientIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithEmptyClientIdAlarmMessageCh(mac, hostname, messageType),
		GenDhcpIllegalOptionWithEmptyClientIdAlarmMessageEn(mac, hostname, messageType),
		CmdDhcpIllegalOptionWithEmptyClientIdAlarm)
}
func GenDhcpIllegalOptionWithEmptyClientIdAlarmMessageCh(mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 未携带OPTION 61")
	return buf.String()
}
func GenDhcpIllegalOptionWithEmptyClientIdAlarmMessageEn(mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" ,The message does not carry OPTION 61")
	return buf.String()
}

// illegal_options
func (a *Alarm) AddDhcpIllegalOptionsAlarm(duid, mac, hostname, messageType string, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionsAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionsAlarmMessageCh(duid, mac, hostname, messageType, optionCode, optionData),
		GenDhcpIllegalOptionsAlarmMessageEn(duid, mac, hostname, messageType, optionCode, optionData),
		CmdDhcpIllegalOptionsAlarm)
}
func GenDhcpIllegalOptionsAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(" ")
	buf.WriteString(optionData)
	return buf.String()
}
func GenDhcpIllegalOptionsAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if duid != "" {
		buf.WriteString("DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString("Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" ,The message carries an illegal OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(" ")
	buf.WriteString(optionData)
	return buf.String()
}
