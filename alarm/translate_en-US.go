package alarm

import (
	"bytes"
	"strconv"
)

func genCpuUsageMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString("CPU usage rate ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% exceeds ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genMemoryUsageMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString("memory usage rate ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% exceeds ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genStoreUsageMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString("storage usage rate ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("% exceeds ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("%")
	return buf.String()
}

func genSubnetRadioMessageEn(ip, subnet string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" subnet:")
	buf.WriteString(subnet)
	buf.WriteString(" usage ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" exceeded ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func genQpsMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" QPS:")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" exceeded ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func genLpsMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" LPS:")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" exceeded ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func genHaTriggerMessageEn(cmd, role, master, slave string) string {
	buf := bytes.Buffer{}
	buf.WriteString("serve ")
	buf.WriteString(role)
	buf.WriteString(" switch from")
	if HaCmd(cmd) == HaCmdMasterUp {
		buf.WriteString(" slave node:")
		buf.WriteString(slave)
		buf.WriteString(" to the master node:")
		buf.WriteString(master)
	} else if HaCmd(cmd) == HaCmdMasterDown {
		buf.WriteString(" master node:")
		buf.WriteString(master)
		buf.WriteString(" to the slave node:")
		buf.WriteString(slave)
	}
	return buf.String()
}

func genNodeOfflineMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" offline")
	return buf.String()
}

func genServiceOfflineMessageEn(node, name string) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(node)
	buf.WriteString(" service ")
	buf.WriteString(name)
	buf.WriteString(" offline")
	return buf.String()
}

func genIpConflictMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("ip ")
	buf.WriteString(ip)
	buf.WriteString(" conflict")
	return buf.String()
}

func genSubnetConflictMessageEn(subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("subnet ")
	buf.WriteString(subnet)
	buf.WriteString(" conflict")
	return buf.String()
}

func genIllegalDhcpMessageEn(ip, mac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("found illegal DHCP service IP:")
	buf.WriteString(ip)
	buf.WriteString(" MAC:")
	buf.WriteString(mac)
	return buf.String()
}

func genIpMacObsoletedMessageEn(mac, obsolete, current string) string {
	buf := bytes.Buffer{}
	buf.WriteString(mac)
	buf.WriteString(" address changed from ")
	buf.WriteString(obsolete)
	buf.WriteString(" to ")
	buf.WriteString(current)
	return buf.String()
}

func genIpPortObsoletedMessageEn(port int, obsolete, current string) string {
	buf := bytes.Buffer{}
	buf.WriteString("the IP of the port:")
	buf.WriteString(strconv.FormatInt(int64(port), 10))
	buf.WriteString(" is changed from ")
	buf.WriteString(obsolete)
	buf.WriteString(" to ")
	buf.WriteString(current)
	return buf.String()
}

func enUSUnManagedIpMsg(ip, subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("subnet:")
	buf.WriteString(subnet)
	buf.WriteString(" of IP:")
	buf.WriteString(ip)
	buf.WriteString(" is unmanaged")
	return buf.String()
}

func genZombieIpMessageEn(ip string, timeOut int64) string {
	buf := bytes.Buffer{}
	buf.WriteString("zombie ip ")
	buf.WriteString(ip)
	buf.WriteString(" offline exceed ")
	buf.WriteString(strconv.FormatInt(timeOut, 10))
	buf.WriteString(" hours")
	return buf.String()
}

func genExpireIpMessageEn(ip string, timeOut int64) string {
	buf := bytes.Buffer{}
	buf.WriteString("online ip ")
	buf.WriteString(ip)
	buf.WriteString(" online exceed ")
	buf.WriteString(strconv.FormatInt(timeOut, 10))
	buf.WriteString(" hours")
	return buf.String()
}
