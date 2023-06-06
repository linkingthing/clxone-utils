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

func genCpuUsageMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点")
	buf.WriteString(ip)
	buf.WriteString("的CPU使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genMemoryUsageMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点")
	buf.WriteString(ip)
	buf.WriteString("的内存使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genStoreUsageMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点")
	buf.WriteString(ip)
	buf.WriteString("的磁盘使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genSubnetRadioMessageCh(ip, subnet string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点")
	buf.WriteString(ip)
	buf.WriteString("地址池 ")
	buf.WriteString(subnet)
	buf.WriteString("的使用率")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% 超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genQpsMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点")
	buf.WriteString(ip)
	buf.WriteString("的QPS ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func genLpsMessageCh(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点")
	buf.WriteString(ip)
	buf.WriteString("的LPS ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("超过了")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func genHaTriggerMessageCh(cmd, role, master, slave string) string {
	buf := bytes.Buffer{}
	buf.WriteString("服务 ")
	buf.WriteString(role)
	buf.WriteString("由")
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

func genNodeOfflineMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(ip)
	buf.WriteString("离线")
	return buf.String()
}

func genServiceOfflineMessageCh(node, name string) string {
	buf := bytes.Buffer{}
	buf.WriteString("节点 ")
	buf.WriteString(node)
	buf.WriteString(" 服务 ")
	buf.WriteString(name)
	buf.WriteString("离线")
	return buf.String()
}

func genIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("IP ")
	buf.WriteString(ip)
	buf.WriteString("冲突")
	return buf.String()
}

func genSubnetConflictMessageCh(subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("子网 ")
	buf.WriteString(subnet)
	buf.WriteString("冲突")
	return buf.String()
}

func genIllegalDhcpMessageCh(ip, mac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("发现非法DHCP服务器IP ")
	buf.WriteString(ip)
	buf.WriteString(" MAC ")
	buf.WriteString(mac)
	return buf.String()
}

func genIpMacObsoletedMessageCh(mac, obsolete, current string) string {
	buf := bytes.Buffer{}
	buf.WriteString(mac)
	buf.WriteString("的IP地址由 ")
	buf.WriteString(obsolete)
	buf.WriteString("变更为 ")
	buf.WriteString(current)
	return buf.String()
}

func genIpPortObsoletedMessageCh(port int, obsolete, current string) string {
	buf := bytes.Buffer{}
	buf.WriteString("端口 ")
	buf.WriteString(strconv.FormatInt(int64(port), 10))
	buf.WriteString(" 的IP由 ")
	buf.WriteString(obsolete)
	buf.WriteString("变更为 ")
	buf.WriteString(current)
	return buf.String()
}

func zhCNUnManagedIpMsg(ip, subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("IP ")
	buf.WriteString(ip)
	buf.WriteString("的所属的子网 ")
	buf.WriteString(subnet)
	buf.WriteString("不在系统规划范围内")
	return buf.String()
}

func genZombieIpMessageCh(ip string, timeOut int64) string {
	buf := bytes.Buffer{}
	buf.WriteString("僵尸地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 离线超过 ")
	buf.WriteString(strconv.FormatInt(timeOut, 10))
	buf.WriteString(" 小时")
	return buf.String()
}

func genExpireIpMessageCh(ip string, timeOut int64) string {
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

func genReservedIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("IP ")
	buf.WriteString(ip)
	buf.WriteString(" 在线且无分配信息，判定地址冲突")
	return buf.String()
}

func genDhcpExcludeIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 在线且地址为排除地址，判定地址冲突")
	return buf.String()
}

func genDhcpDynamicMacIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("DHCP地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 的MAC与采集到的MAC不符，判定地址冲突")
	return buf.String()
}

func genDhcpReservationMacIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("DHCP地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 的MAC与采集到的MAC不符，判定地址冲突")
	return buf.String()
}

func genDhcpDynamicIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("动态地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 在线且无DHCP租赁信息，判定地址冲突")
	return buf.String()
}

func genDhcpReservationIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("固定地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 在线且无DHCP租赁信息，判定地址冲突")
	return buf.String()
}

func genDhcpReservedIpConflictMessageCh(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("保留地址 ")
	buf.WriteString(ip)
	buf.WriteString(" 在线且未登记，判定地址冲突")
	return buf.String()
}
