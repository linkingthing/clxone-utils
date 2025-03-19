package alarm

import (
	"bytes"
	"strconv"
	"strings"

	pb "github.com/linkingthing/clxone-utils/alarm/proto"
)

type DHCPClient interface {
	GetHwAddress() string
	GetHostname() string
	GetDuid() string
	GetMessageType() string
}

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

func (c *DHCPClient4) GetHwAddress() string {
	return c.HwAddress
}

func (c *DHCPClient4) GetHostname() string {
	return c.Hostname
}
func (c *DHCPClient4) GetDuid() string {
	return ""
}
func (c *DHCPClient4) GetMessageType() string {
	return c.MessageType
}

type DHCPClient6 struct {
	Duid                  string
	HwAddress             string
	HwAddressOrganization string
	Hostname              string
	Fingerprint           string
	VendorId              string
	OperatingSystem       string
	ClientType            string
	RequestSourceAddr     string
	MessageType           string
}

func (c *DHCPClient6) GetHwAddress() string {
	return c.HwAddress
}

func (c *DHCPClient6) GetHostname() string {
	return c.Hostname
}
func (c *DHCPClient6) GetDuid() string {
	return c.Duid
}
func (c *DHCPClient6) GetMessageType() string {
	return c.MessageType
}

// illegal_packet_with_opcode_alarm
func (a *Alarm) AddDhcpIllegalPacketWithOpcodeAlarm(client DHCPClient, opcode string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalPacketWithOpcodeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalPacketWithOpcodeMessageEn(client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), opcode),
		GenDhcpIllegalPacketWithOpcodeMessageCh(client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), opcode),
		CmdDhcpIllegalPacketWithOpcodeAlarm)
}
func GenDhcpIllegalPacketWithOpcodeMessageCh(mac, hostname, messageType, opcode string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 MAC：")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 的非法报文 ")
	buf.WriteString(messageType)
	buf.WriteString("，")
	buf.WriteString("不支持操作码 ")
	buf.WriteString(opcode)
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
	buf.WriteString(". The operation code ")
	buf.WriteString(opcode)
	buf.WriteString(" is not supported.")
	return buf.String()
}

// illegal_inform_packet_without_source_addr_alarm
func (a *Alarm) AddDhcpIllegalInformPacketWithoutSourceAddrAlarm(client DHCPClient) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalInformPacketWithoutSourceAddrAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalInformPacketWithoutSourceAddrAlarmMessageEn(client.GetHwAddress(), client.GetHostname(), client.GetMessageType()),
		GenDhcpIllegalInformPacketWithoutSourceAddrAlarmMessageCh(client.GetHwAddress(), client.GetHostname(), client.GetMessageType()),
		CmdDhcpIllegalInformPacketWithoutSourceAddrAlarm)
}
func GenDhcpIllegalInformPacketWithoutSourceAddrAlarmMessageCh(mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 MAC：")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 的非法报文 ")
	buf.WriteString(messageType)
	buf.WriteString("，请求源地址为空")
	return buf.String()
}
func GenDhcpIllegalInformPacketWithoutSourceAddrAlarmMessageEn(mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("an illegal message of ")
	buf.WriteString(messageType)
	buf.WriteString("was received from the client with")
	if mac != "" {
		buf.WriteString(" Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString("The source address of the request is empty")
	return buf.String()
}

// illegal_client_alarm
func (a *Alarm) AddDhcpIllegalClientAlarm(client DHCPClient) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalClientAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalClientAlarmMessageEn(client.GetHwAddress(), client.GetDuid(), client.GetHostname()),
		GenDhcpIllegalClientAlarmMessageCh(client.GetHwAddress(), client.GetDuid(), client.GetHostname()),
		CmdDhcpIllegalClientAlarm)
}
func GenDhcpIllegalClientAlarmMessageCh(mac, duid, hostname string) string {
	buf := bytes.Buffer{}
	buf.WriteString("客户端 ")
	if duid != "" {
		buf.WriteString("DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 正在尝试非法接入")
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
func (a *Alarm) AddDhcpIllegalClientWithHighQpsAlarm(client DHCPClient, rateLimit uint32) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalClientWithHighQpsAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalClientWithHighQpsAlarmMessageEn(client.GetHwAddress(), client.GetDuid(), client.GetHostname(), rateLimit),
		GenDhcpIllegalClientWithHighQpsAlarmMessageCh(client.GetHwAddress(), client.GetDuid(), client.GetHostname(), rateLimit),
		CmdDhcpIllegalClientWithHighQpsAlarm)
}
func GenDhcpIllegalClientWithHighQpsAlarmMessageCh(mac, duid, hostname string, rateLimit uint32) string {
	buf := bytes.Buffer{}
	buf.WriteString("客户端 ")
	if duid != "" {
		buf.WriteString("DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 请求速度超过阈值 ")
	buf.WriteString(strconv.Itoa(int(rateLimit)))
	buf.WriteString("r/s")
	return buf.String()
}
func GenDhcpIllegalClientWithHighQpsAlarmMessageEn(mac, duid, hostname string, rateLimit uint32) string {
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
	buf.WriteString(strconv.Itoa(int(rateLimit)))
	buf.WriteString(" r/s")
	return buf.String()
}

// illegal_option_with_unexpected_message_type_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUnexpectedMessageTypeAlarm(client DHCPClient) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUnexpectedMessageTypeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType()),
		GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType()),
		CmdDhcpIllegalOptionWithUnexpectedMessageTypeAlarm)
}
func GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageCh(duid, mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString("DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	if duid == "" && mac != "" {
		buf.WriteString(" 报文携带非法OPTION 53")
	}
	if duid != "" {
		buf.WriteString(" 非法报文")
	}
	buf.WriteString("，不支持报文类型 ")
	buf.WriteString(messageType)
	return buf.String()
}
func GenDhcpIllegalOptionWithUnexpectedMessageTypeAlarmMessageEn(duid, mac, hostname, messageType string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with ")
	if duid != "" {
		buf.WriteString(" DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" and Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" and hostname, ")
		buf.WriteString(hostname)
	}
	if duid == "" && mac != "" {
		buf.WriteString(" carrying an illegal OPTION 53,")
	}
	buf.WriteString(" The message type is not supported")
	return buf.String()
}

// illegal_option_with_ultra_short_lease_time_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUltraShortLeaseTimeAlarm(client DHCPClient, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUltraShortLeaseTimeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageEn(client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageCh(client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		CmdDhcpIllegalOptionWithUltraShortLeaseTimeAlarm)
}
func GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageCh(mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 MAC：")
	buf.WriteString(mac)

	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString("，请求租约时间 ")
	buf.WriteString(optionData)
	buf.WriteString(" 过短")
	return buf.String()
}
func GenDhcpIllegalOptionWithUltraShortLeaseTimeAlarmMessageEn(mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with Mac address ")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(". The requested lease time ")
	buf.WriteString(optionData)
	buf.WriteString(" is too short")
	return buf.String()
}

// illegal_option_with_ultra_long_lease_time_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUltraLongLeaseTimeAlarm(client DHCPClient, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUltraLongLeaseTimeAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageEn(client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageCh(client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		CmdDhcpIllegalOptionWithUltraLongLeaseTimeAlarm)
}
func GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageCh(mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 MAC：")
	buf.WriteString(mac)

	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString("，请求租约时间 ")
	buf.WriteString(optionData)
	buf.WriteString(" 过长")
	return buf.String()
}
func GenDhcpIllegalOptionWithUltraLongLeaseTimeAlarmMessageEn(mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with Mac address ")
	buf.WriteString(mac)
	if hostname != "" {
		buf.WriteString(" and hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(". The requested lease time ")
	buf.WriteString(optionData)
	buf.WriteString(" is too long")
	return buf.String()
}

// illegal_option_with_invalid_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithInvalidServerIdAlarm(client DHCPClient, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithInvalidServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		CmdDhcpIllegalOptionWithInvalidServerIdAlarm)
}
func GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString("，server id ")
	buf.WriteString(optionData)
	buf.WriteString(" 格式错误")
	return buf.String()
}
func GenDhcpIllegalOptionWithInvalidServerIdAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with")
	if duid != "" {
		buf.WriteString(" DUID ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(", Mac address ")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(", hostname ")
		buf.WriteString(hostname)
	}
	buf.WriteString(" , carrying an illegal OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(". The format of the server ID ")
	buf.WriteString(optionData)
	buf.WriteString(" is incorrect")
	return buf.String()
}

// illegal_option_with_unexpected_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithUnexpectedServerIdAlarm(client DHCPClient, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithUnexpectedServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		CmdDhcpIllegalOptionWithUnexpectedServerIdAlarm)
}
func GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString("，server id ")
	buf.WriteString(optionData)
	buf.WriteString(" 与当前服务器不匹配")
	return buf.String()
}
func GenDhcpIllegalOptionWithUnexpectedServerIdAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
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
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(". The server ID ")
	buf.WriteString(optionData)
	buf.WriteString(" does not match the current server address")
	return buf.String()
}

// illegal_option_with_forbidden_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithForbiddenServerIdAlarm(client DHCPClient, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithForbiddenServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		CmdDhcpIllegalOptionWithForbiddenServerIdAlarm)
}
func GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString("，禁止携带 server id ")
	buf.WriteString(optionData)
	return buf.String()
}
func GenDhcpIllegalOptionWithForbiddenServerIdAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
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
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString(". The server ID ")
	buf.WriteString(optionData)
	buf.WriteString(" is not allowed")
	return buf.String()
}

// illegal_option_with_mandatory_server_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithMandatoryServerIdAlarm(client DHCPClient, optionCode uint32) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithMandatoryServerIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode),
		GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode),
		CmdDhcpIllegalOptionWithMandatoryServerIdAlarm)
}
func GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString("，未携带OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	return buf.String()
}
func GenDhcpIllegalOptionWithMandatoryServerIdAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32) string {
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
	buf.WriteString(strconv.Itoa(int(optionCode)))
	return buf.String()
}

// illegal_option_with_invalid_client_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithInvalidClientIdAlarm(client DHCPClient, optionCode uint32, optionData string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithInvalidClientIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode, optionData),
		CmdDhcpIllegalOptionWithInvalidClientIdAlarm)
}
func GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString(" DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	buf.WriteString("，client id ")
	buf.WriteString(optionData)
	buf.WriteString(" 格式错误")
	return buf.String()
}
func GenDhcpIllegalOptionWithInvalidClientIdAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32, optionData string) string {
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
	buf.WriteString(",The format of the client ID")
	buf.WriteString(optionData)
	buf.WriteString(" is incorrect")
	return buf.String()
}

// illegal_option_with_empty_client_id_alarm
func (a *Alarm) AddDhcpIllegalOptionWithMandatoryClientIdAlarm(client DHCPClient, optionCode uint32) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionWithMandatoryClientIdAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionWithMandatoryClientIdAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode),
		GenDhcpIllegalOptionWithMandatoryClientIdAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), optionCode),
		CmdDhcpIllegalOptionWithMandatoryClientIdAlarm)
}
func GenDhcpIllegalOptionWithMandatoryClientIdAlarmMessageCh(duid, mac, hostname, messageType string, optionCode uint32) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString("DUID：")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString("，未携带OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	return buf.String()
}
func GenDhcpIllegalOptionWithMandatoryClientIdAlarmMessageEn(duid, mac, hostname, messageType string, optionCode uint32) string {
	buf := bytes.Buffer{}
	buf.WriteString("Received a message of type ")
	buf.WriteString(messageType)
	buf.WriteString(" from the client with")
	if duid != "" {
		buf.WriteString(" DUID ")
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
	buf.WriteString(", The message does not carry OPTION ")
	buf.WriteString(strconv.Itoa(int(optionCode)))
	return buf.String()
}

// illegal_options
func (a *Alarm) AddDhcpIllegalOptionsAlarm(client DHCPClient, illegalOptions map[uint32]string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpIllegalOptionsAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpIllegalOptionsAlarmMessageEn(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), illegalOptions),
		GenDhcpIllegalOptionsAlarmMessageCh(client.GetDuid(), client.GetHwAddress(), client.GetHostname(), client.GetMessageType(), illegalOptions),
		CmdDhcpIllegalOptionsAlarm)
}
func GenDhcpIllegalOptionsAlarmMessageCh(duid, mac, hostname, messageType string, illegalOptions map[uint32]string) string {
	buf := bytes.Buffer{}
	buf.WriteString("收到客户端 ")
	if duid != "" {
		buf.WriteString("DUID： ")
		buf.WriteString(duid)
	}
	if mac != "" {
		buf.WriteString(" MAC：")
		buf.WriteString(mac)
	}
	if hostname != "" {
		buf.WriteString(" 主机名：")
		buf.WriteString(hostname)
	}
	buf.WriteString(" 报文 ")
	buf.WriteString(messageType)
	buf.WriteString(" 携带非法OPTION ")
	for code, data := range illegalOptions {
		buf.WriteString(strconv.Itoa(int(code)))
		buf.WriteString(" ")
		buf.WriteString(data)
		buf.WriteString("，")
	}
	return strings.TrimRight(buf.String(), "，")
}
func GenDhcpIllegalOptionsAlarmMessageEn(duid, mac, hostname, messageType string, illegalOptions map[uint32]string) string {
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
	for code, data := range illegalOptions {
		buf.WriteString(strconv.Itoa(int(code)))
		buf.WriteString(" ")
		buf.WriteString(data)
		buf.WriteString(",")
	}
	return strings.TrimRight(buf.String(), ",")
}

func (a *Alarm) AddDhcpLeaseExceptionAlarm(subnet, ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpLeaseExceptionAlarm)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		GenDhcpLeaseExceptionAlarmMessageEn(subnet, ip),
		GenDhcpLeaseExceptionAlarmMessageCh(subnet, ip),
		CmdDhcpLeaseExceptionAlarm)
}
func GenDhcpLeaseExceptionAlarmMessageCh(subnet, ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("子网 ")
	buf.WriteString(subnet)
	buf.WriteString(" 地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 租约异常，地址已被其他客户端使用")
	return buf.String()
}
func GenDhcpLeaseExceptionAlarmMessageEn(subnet, ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("The lease for address ")
	buf.WriteString(ip)
	buf.WriteString(" in the subnet ")
	buf.WriteString(subnet)
	buf.WriteString(" is abnormal. The address has been used by another client")
	return buf.String()
}
