package alarm

import (
	"bytes"
	"strconv"
)

func GenCpuUsageMessageEn(ip string, value, limit uint64) string {
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

func GenMemoryUsageMessageEn(ip string, value, limit uint64) string {
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

func GenStoreUsageMessageEn(ip string, value, limit uint64) string {
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

func GenTemperatureUsageMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString("cpu temperature value ")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString("°C exceeds ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	buf.WriteString("°C")
	return buf.String()
}

func GenSubnetRadioMessageEn(ip, subnet string, value, limit uint64) string {
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

func GenQpsMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" QPS:")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" exceeded ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func GenLpsMessageEn(ip string, value, limit uint64) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" LPS:")
	buf.WriteString(strconv.FormatUint(value, 10))
	buf.WriteString(" exceeded ")
	buf.WriteString(strconv.FormatUint(limit, 10))
	return buf.String()
}

func GenHaTriggerMessageEn(cmd, role, master, slave string) string {
	buf := bytes.Buffer{}
	buf.WriteString("Ha switch from")
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

func GenHaStateTriggerMessageEn(cmd string) string {
	buf := bytes.Buffer{}
	if HaCmd(cmd) == HaCmdMasterUp {
		buf.WriteString("The HA node switches from slave to master, and the HA high-availability status returns to normal")
	} else if HaCmd(cmd) == HaCmdMasterDown {
		buf.WriteString("The HA node switches from master to slave, and the HA high-availability status is abnormal")
	}
	return buf.String()
}

func GenBackupTriggerMessageEn(ip string, start bool) string {
	buf := bytes.Buffer{}
	if start {
		buf.WriteString("The cluster is unavailable, switch to the disaster recovery node ")
		buf.WriteString(ip)
	} else {
		buf.WriteString("The cluster has been restored. Switch back from the disaster recovery node ")
		buf.WriteString(ip)
		buf.WriteString(" to the cluster")
	}
	return buf.String()
}

func GenBackupStateTriggerMessageEn(start bool) string {
	buf := bytes.Buffer{}
	if start {
		buf.WriteString("The cluster switches to the disaster recovery node, and the cluster's high-availability status is abnormal")
	} else {
		buf.WriteString("The cluster switches back to the primary node, and the cluster's high-availability status returns to normal")
	}
	return buf.String()
}

func GenNodeOfflineMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(ip)
	buf.WriteString(" offline")
	return buf.String()
}

func GenServiceOfflineMessageEn(node, name string) string {
	buf := bytes.Buffer{}
	buf.WriteString("node ")
	buf.WriteString(node)
	buf.WriteString(" service ")
	buf.WriteString(name)
	buf.WriteString(" offline")
	return buf.String()
}

func GenDatabaseOfflineMessageEn(name string) string {
	buf := bytes.Buffer{}
	buf.WriteString("database ")
	buf.WriteString(name)
	buf.WriteString(" offline")
	return buf.String()
}

func GenSubnetConflictMessageEn(subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("subnet ")
	buf.WriteString(subnet)
	buf.WriteString(" conflict")
	return buf.String()
}

func GenIllegalDhcpMessageEn(ip, mac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("found illegal DHCP service IP:")
	buf.WriteString(ip)
	buf.WriteString(" MAC:")
	buf.WriteString(mac)
	return buf.String()
}

func GenIpMacObsoletedMessageEn(device, ip, oldMac, newMac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("the mac bound onto ip ")
	buf.WriteString(ip)
	buf.WriteString(" for device ")
	buf.WriteString(device)
	buf.WriteString(" from ")
	buf.WriteString(oldMac)
	buf.WriteString(" changed to ")
	buf.WriteString(newMac)
	return buf.String()
}

func GenIpPortObsoletedMessageEn(equip, port, obsolete, current string) string {
	buf := bytes.Buffer{}
	buf.WriteString("the IP of the port ")
	buf.WriteString(port)
	buf.WriteString(" on equipment ")
	buf.WriteString(equip)
	buf.WriteString(" is changed from ")
	buf.WriteString(obsolete)
	buf.WriteString(" to ")
	buf.WriteString(current)
	return buf.String()
}

func GenUnManagedIpMsg(ip, subnet string) string {
	buf := bytes.Buffer{}
	buf.WriteString("subnet:")
	buf.WriteString(subnet)
	buf.WriteString(" of IP:")
	buf.WriteString(ip)
	buf.WriteString(" is unmanaged")
	return buf.String()
}

func GenZombieIpMessageEn(ip string, timeOut int64) string {
	buf := bytes.Buffer{}
	buf.WriteString("zombie ip ")
	buf.WriteString(ip)
	buf.WriteString(" offline exceed ")
	buf.WriteString(strconv.FormatInt(timeOut, 10))
	buf.WriteString(" hours")
	return buf.String()
}

func GenExpireIpMessageEn(ip string, timeOut int64) string {
	buf := bytes.Buffer{}
	buf.WriteString("online ip ")
	buf.WriteString(ip)
	buf.WriteString(" online exceed ")
	buf.WriteString(strconv.FormatInt(timeOut, 10))
	buf.WriteString(" hours")
	return buf.String()
}

func GenReservedIpConflictMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("ip ")
	buf.WriteString(ip)
	buf.WriteString(" is online without allocation information")
	return buf.String()
}

func GenDhcpExcludeIpConflictMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("online ip")
	buf.WriteString(ip)
	buf.WriteString(" is exclusion address")
	return buf.String()
}

func GenDhcpDynamicMacIpConflictMessageEn(ip, ipMac, collectMac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("the mac [")
	buf.WriteString(ipMac)
	buf.WriteString("] of ip ")
	buf.WriteString(ip)
	buf.WriteString(" is different from collected mac [")
	buf.WriteString(collectMac)
	buf.WriteString("]")
	return buf.String()
}

func GenDhcpReservationMacIpConflictMessageEn(ip, ipMac, collectMac string) string {
	buf := bytes.Buffer{}
	buf.WriteString("the mac [")
	buf.WriteString(ipMac)
	buf.WriteString("] of ip ")
	buf.WriteString(ip)
	buf.WriteString(" is different from collected mac [")
	buf.WriteString(collectMac)
	buf.WriteString("]")
	return buf.String()
}

func GenDhcpDynamicIpConflictMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("no lease info for online dynamic ip")
	buf.WriteString(ip)
	return buf.String()
}

func GenDhcpReservationIpConflictMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("no lease info for online reservation ip")
	buf.WriteString(ip)
	return buf.String()
}

func GenDhcpReservedIpConflictMessageEn(ip string) string {
	buf := bytes.Buffer{}
	buf.WriteString("reserved online ip")
	buf.WriteString(ip)
	buf.WriteString(" isn't registered")
	return buf.String()
}
