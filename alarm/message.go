package alarm

import (
	"context"

	"github.com/golang/protobuf/proto"
	kg "github.com/segmentio/kafka-go"

	pb "github.com/linkingthing/clxone-utils/alarm/proto"
)

func (a *Alarm) GetThreshold(name pb.ThresholdName) *Threshold {
	a.Lock()
	defer a.Unlock()
	threshold, ok := a.ThresholdMap[name]
	if !ok {
		return nil
	} else if threshold.Enabled == false {
		return nil
	}
	return threshold
}

func (a *Alarm) sendAlarmToKafka(threshold *Threshold, message, messageDic, cmd string) error {
	data, err := proto.Marshal(&pb.Alarm{
		Name:              threshold.GetName().String(),
		Level:             threshold.GetLevel().String(),
		SendMail:          threshold.SendMail,
		Message:           message,
		MessageDictionary: messageDic,
	})
	if err != nil {
		return err
	}

	return a.alarmWriter.WriteMessages(context.Background(),
		kg.Message{Topic: AlarmTopic, Key: []byte(cmd), Value: data})
}

func (a *Alarm) AddCpuUsageAlarm(ip string, value uint64) error {
	threshold := a.GetThreshold(pb.ThresholdName_cpuUsedRatio)
	if threshold == nil {
		return nil
	} else if value < threshold.Value {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genCpuUsageMessageEn(ip, value, threshold.Value),
		genCpuUsageMessageCh(ip, value, threshold.Value),
		CmdCpuUsageAlarm)
}

func (a *Alarm) AddMemoryUsageAlarm(ip string, value uint64) error {
	threshold := a.GetThreshold(pb.ThresholdName_memoryUsedRatio)
	if threshold == nil {
		return nil
	} else if value < threshold.Value {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genMemoryUsageMessageEn(ip, value, threshold.Value),
		genMemoryUsageMessageCh(ip, value, threshold.Value),
		CmdMemoryUsageAlarm)
}

func (a *Alarm) AddStorageUsageAlarm(ip string, value uint64) error {
	threshold := a.GetThreshold(pb.ThresholdName_storageUsedRatio)
	if threshold == nil {
		return nil
	} else if value < threshold.Value {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genStoreUsageMessageEn(ip, value, threshold.Value),
		genStoreUsageMessageCh(ip, value, threshold.Value),
		CmdStorageUsageAlarm)
}

func (a *Alarm) AddSubnetRadioAlarm(ip, subnet string, value uint64) error {
	threshold := a.GetThreshold(pb.ThresholdName_subnetUsedRatio)
	if threshold == nil {
		return nil
	} else if value < threshold.Value {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genSubnetRadioMessageEn(ip, subnet, value, threshold.Value),
		genSubnetRadioMessageCh(ip, subnet, value, threshold.Value),
		CmdSubnetRadioAlarm)
}

func (a *Alarm) AddSubnetConflictAlarm(name string) error {
	threshold := a.GetThreshold(pb.ThresholdName_subnetConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genSubnetConflictMessageEn(name),
		genSubnetConflictMessageCh(name),
		CmdConflictSubnetAlarm)
}

func (a *Alarm) AddQPSAlarm(ip string, value uint64) error {
	threshold := a.GetThreshold(pb.ThresholdName_qps)
	if threshold == nil {
		return nil
	} else if value < threshold.Value {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genQpsMessageEn(ip, value, threshold.Value),
		genQpsMessageCh(ip, value, threshold.Value),
		CmdQpsAlarm)
}

func (a *Alarm) AddLPSAlarm(ip string, value uint64) error {
	threshold := a.GetThreshold(pb.ThresholdName_lps)
	if threshold == nil {
		return nil
	} else if value < threshold.Value {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genLpsMessageEn(ip, value, threshold.Value),
		genLpsMessageCh(ip, value, threshold.Value),
		CmdLpsAlarm)
}

func (a *Alarm) AddHaTriggerAlarm(cmd, role, master, slave string) error {
	threshold := a.GetThreshold(pb.ThresholdName_haTrigger)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genHaTriggerMessageEn(cmd, role, master, slave),
		genHaTriggerMessageCh(cmd, role, master, slave),
		CmdHaTriggerAlarm)
}

func (a *Alarm) AddNodeOfflineAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_nodeOffline)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genNodeOfflineMessageEn(ip),
		genNodeOfflineMessageCh(ip),
		CmdNodeOfflineAlarm)
}

func (a *Alarm) AddServiceOfflineAlarm(node, name string) error {
	threshold := a.GetThreshold(pb.ThresholdName_serviceOffline)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genServiceOfflineMessageEn(node, name),
		genServiceOfflineMessageCh(node, name),
		CmdServiceOfflineAlarm)
}

func (a *Alarm) AddIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_ipConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genIpConflictMessageEn(ip),
		genIpConflictMessageCh(ip),
		CmdConflictIpAlarm)
}

func (a *Alarm) AddIllegalDHCPAlarm(ip, mac string) error {
	threshold := a.GetThreshold(pb.ThresholdName_illegalDhcp)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genIllegalDhcpMessageEn(ip, mac),
		genIllegalDhcpMessageCh(ip, mac),
		CmdIllegalDhcpAlarm)
}

func (a *Alarm) AddIpMacObsoletedAlarm(mac, obsolete, current string) error {
	threshold := a.GetThreshold(pb.ThresholdName_ipMacObsoleted)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genIpMacObsoletedMessageEn(mac, obsolete, current),
		genIpMacObsoletedMessageCh(mac, obsolete, current),
		CmdIpMacObsoletedAlarm)
}

func (a *Alarm) AddIpPortObsoletedAlarm(port int, obsolete, current string) error {
	threshold := a.GetThreshold(pb.ThresholdName_ipPortObsoleted)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genIpPortObsoletedMessageEn(port, obsolete, current),
		genIpPortObsoletedMessageCh(port, obsolete, current),
		CmdIpPortObsoletedAlarm)
}

func (a *Alarm) AddUnmanagedIpAlarm(ip, subnet string) error {
	threshold := a.GetThreshold(pb.ThresholdName_ipUnmanaged)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		enUSUnManagedIpMsg(ip, subnet),
		zhCNUnManagedIpMsg(ip, subnet),
		CmdIpUnmanagedAlarm)
}

func (a *Alarm) AddZombieIpAlarm(ip string, timeOut int64) error {
	threshold := a.GetThreshold(pb.ThresholdName_zombieIp)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genZombieIpMessageEn(ip, timeOut),
		genZombieIpMessageCh(ip, timeOut),
		CmdZombieIpAlarm)
}

func (a *Alarm) AddExpireIpAlarm(ip string, timeOut int64) error {
	threshold := a.GetThreshold(pb.ThresholdName_onlineExpiredIp)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genExpireIpMessageEn(ip, timeOut),
		genExpireIpMessageCh(ip, timeOut),
		CmdOnlineExpiredIpAlarm)
}

func (a *Alarm) AddReservedIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_reservedIpConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genReservedIpConflictMessageEn(ip),
		genReservedIpConflictMessageCh(ip),
		CmdReservedIpConflictAlarm)
}

func (a *Alarm) AddDhcpExcludeIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpExcludeIpConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genDhcpExcludeIpConflictMessageEn(ip),
		genDhcpExcludeIpConflictMessageCh(ip),
		CmdDhcpExcludeIpConflictAlarm)
}

func (a *Alarm) AddDhcpDynamicMacIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpDynamicMacIpConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genDhcpDynamicMacIpConflictMessageEn(ip),
		genDhcpDynamicMacIpConflictMessageCh(ip),
		CmdDhcpDynamicMacIpConflictAlarm)
}

func (a *Alarm) AddDhcpReservationMacIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpReservationMacIpConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genDhcpReservationMacIpConflictMessageEn(ip),
		genDhcpReservationMacIpConflictMessageCh(ip),
		CmdDhcpReservationMacIpConflictAlarm)
}

func (a *Alarm) AddDhcpDynamicIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpDynamicIpConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genDhcpDynamicIpConflictMessageEn(ip),
		genDhcpDynamicIpConflictMessageCh(ip),
		CmdDhcpDynamicIpConflictAlarm)
}

func (a *Alarm) AddDhcpReservedIpConflictAlarm(ip string) error {
	threshold := a.GetThreshold(pb.ThresholdName_dhcpReservedIpConflict)
	if threshold == nil {
		return nil
	}

	return a.sendAlarmToKafka(threshold,
		genDhcpReservedIpConflictMessageEn(ip),
		genDhcpReservedIpConflictMessageCh(ip),
		CmdDhcpReservedIpConflictAlarm)
}
