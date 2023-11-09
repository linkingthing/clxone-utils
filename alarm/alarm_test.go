package alarm

import (
	pb "github.com/linkingthing/clxone-utils/alarm/proto"
	"testing"
)

func TestRegisterAlarm(t *testing.T) {
	thresholds := []*Threshold{
		&Threshold{
			BaseThreshold: &pb.BaseThreshold{
				Name:  pb.ThresholdName_subnetConflict,
				Level: pb.ThresholdLevel_critical,
				Type:  pb.ThresholdType_trigger,
			},
			Value:    0,
			SendMail: false,
			Enabled:  true,
		},
		{
			BaseThreshold: &pb.BaseThreshold{
				Name:  pb.ThresholdName_ipMacObsoleted,
				Level: pb.ThresholdLevel_critical,
				Type:  pb.ThresholdType_trigger,
			},
			Value:    0,
			SendMail: false,
			Enabled:  true,
		},
		{
			BaseThreshold: &pb.BaseThreshold{
				Name:  pb.ThresholdName_ipPortObsoleted,
				Level: pb.ThresholdLevel_critical,
				Type:  pb.ThresholdType_trigger,
			},
			Value:    0,
			SendMail: false,
			Enabled:  true,
		},
		{
			BaseThreshold: &pb.BaseThreshold{
				Name:  pb.ThresholdName_ipUnmanaged,
				Level: pb.ThresholdLevel_critical,
				Type:  pb.ThresholdType_trigger,
			},
			Value:    0,
			SendMail: false,
			Enabled:  true,
		},
	}

	alarm, err := RegisterAlarm(KafkaConf{
		Addresses: []string{"10.0.0.66:29092"},
		Topic:     ThresholdIpamTopic,
		GroupId:   "ipam_group_10.0.0.65",
	}, thresholds...)
	if err != nil {
		t.Errorf("register alarm failed:%s", err.Error())
		return
	}

	if err := alarm.AddSubnetConflictAlarm("20.0.0.0/24"); err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := alarm.AddIpMacObsoletedAlarm("devicd", "20.0.0.1", "11:11:11:11:11:11", "22:22:22:22:22:22"); err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := alarm.AddIpPortObsoletedAlarm("equip", "x", "11.0.0.1", "11.0.0.2"); err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := alarm.AddUnmanagedIpAlarm("11.0.0.1", "11.0.0.0/24"); err != nil {
		t.Errorf(err.Error())
		return
	}
}
