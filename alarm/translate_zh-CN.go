package alarm

import (
	"bytes"
	"strconv"
)

type HaCmd string

const (
	HaCmdStartHa    HaCmd = "start_ha"
	HaCmdMasterUp   HaCmd = "master_up"
	HaCmdMasterDown HaCmd = "master_down"
)

func GenCpuUsageMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 的CPU使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func GenMemoryUsageMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 的内存使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func GenStoreUsageMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 的磁盘使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func GenSubnetRadioMessageCh(ip, subnet string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 地址池 ")
	buf.WriteString(subnet)
	buf.WriteString(" 的使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func GenQpsMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 的QPS ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func GenLpsMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 的LPS ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func GenHaTriggerMessageCh(cmd, role, master, slave string) string {
	buf := bytes.Buffer{}
	buf.WriteString("服务 ")
	buf.WriteString(role)
	buf.WriteString(" 由")
	if HaCmd(cmd) == HaCmdMasterUp {
		buf.WriteString("辅节点 ")
		buf.WriteString(slave)
		buf.WriteString("切换到主节点 ")
		buf.WriteString(master)
	} else if HaCmd(cmd) == HaCmdMasterDown {
		buf.WriteString("主节点 ")
		buf.WriteString(master)
		buf.WriteString("切换到辅节点 ")
		buf.WriteString(slave)
	}
	return buf.String()
}

func GenNodeOfflineMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString(" 离线")
	return buf.String()
}

func GenServiceOfflineMessageCh(node, name string) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(node)
	buf.WriteString(" 服务 ")
	buf.WriteString(name)
	buf.WriteString(" 离线")
	return buf.String()
}

func GenDbOfflineMessageCh(name string) string {
	buf := bytes.Buffer{}
	buf.WriteString("数据库 ")
	buf.WriteString(name)
	buf.WriteString(" 离线")
	return buf.String()
}

func GenSubnetConflictMessageCh(subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("子网 ")
	buf.WriteString(subnet)
	buf.WriteString(" 冲突")
	return buf.String()
}

func GenIllegalDhcpMessageCh(ip, mac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("发现非法DHCP服务器IP ")
	buf.WriteString(ip)
	buf.WriteString(" MAC ")
	buf.WriteString(mac)
	return buf.String()
}

func GenIpMacObsoletedMessageCh(device, ip, oldMac, newMac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("终端 ")
	buf.WriteString(device)
	buf.WriteString(" 的IP ")
	buf.WriteString(ip)
	buf.WriteString(" 上所绑定的MAC从 ")
	buf.WriteString(oldMac)
	buf.WriteString(" 变更为 ")
	buf.WriteString(newMac)
	return buf.String()
}

func GenIpPortObsoletedMessageCh(equip, port, obsolete, current string) string {
	buf := bytes.Buffer{}
	buf.WriteString("设备 ")
	buf.WriteString(equip)
	buf.WriteString(" 的端口 ")
	buf.WriteString(port)
	buf.WriteString(" 上的IP由 ")
	buf.WriteString(obsolete)
	buf.WriteString(" 变更为 ")
	buf.WriteString(current)
	return buf.String()
}

func GenUnManagedIpMsgCh(ip, subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("IP ")
	buf.WriteString(ip)
	buf.WriteString(" 的所属的子网 ")
	buf.WriteString(subnet)
	buf.WriteString(" 不在系统规划范围内")
	return buf.String()
}

func GenZombieIpMessageCh(ip string, timeOut int64) string {
	buf := bytes.Buffer{}
	buf.WriteString("僵尸地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 离线超过 ")
	buf.WriteString(strconv.FormatInt(timeOut, 10))
	buf.WriteString(" 小时")
	return buf.String()
}

func GenExpireIpMessageCh(ip string, timeOut int64) string {
	days := strconv.FormatInt(timeOut/24, 10)
	hours := strconv.FormatInt(timeOut%24, 10)
	buf := bytes.Buffer{}
	buf.WriteString(ip)
	buf.WriteString(" 在线状态超过")
	if days != "0" {
		buf.WriteString(days)
		buf.WriteString("天")
	}
	if hours != "0" {
		buf.WriteString(hours)
		buf.WriteString("小时")
	}
	buf.WriteString("未更新，IP过期")
	return buf.String()
}

func GenReservedIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("静态子网中预留的IP地址 ")
	buf.WriteString(ip)
	buf.WriteString(" ，采集到该IP地址在线，产生冲突告警")
	return buf.String()
}

func GenDhcpExcludeIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("DHCP子网中的排除地址 ")
	buf.WriteString(ip)
	buf.WriteString(" ，采集到该IP地址在线，产生IP冲突告警")
	return buf.String()
}

func GenDhcpDynamicMacIpConflictMessageCh(ip, ipMac, collectMac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("DHCP租赁地址中动态地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 的MAC[")
	buf.WriteString(ipMac)
	buf.WriteString("] 与采集的MAC[")
	buf.WriteString(collectMac)
	buf.WriteString("] 不匹配，产生IP冲突告警")
	return buf.String()
}

func GenDhcpReservationMacIpConflictMessageCh(ip, ipMac, collectMac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("DHCP租赁地址中固定地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 的MAC[")
	buf.WriteString(ipMac)
	buf.WriteString("] 与采集的MAC[")
	buf.WriteString(collectMac)
	buf.WriteString("] 不匹配，产生IP冲突告警")
	return buf.String()
}

func GenDhcpDynamicIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("动态IP地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 无租赁，但采集到该IP在线，产生IP冲突告警")
	return buf.String()
}

func GenDhcpReservationIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("固定IP地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 无租赁，但采集到该IP在线，产生IP冲突告警")
	return buf.String()
}

func GenDhcpReservedIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("DHCP保留地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 未登记，但采集到该IP在线，产生IP冲突告警")
	return buf.String()
}
