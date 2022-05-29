package alarm

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	kg "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"

	pb "github.com/linkingthing/clxone-utils/alarm/proto"
)

type Alarm struct {
	ThresholdMap    map[pb.ThresholdName]*Threshold
	thresholdReader *kg.Reader
	thresholdWriter *kg.Writer
	alarmWriter     *kg.Writer
	sync.RWMutex
}

type KafkaConf struct {
	Addresses []string
	Topic     string
	GroupId   string
	Username  string
	Password  string
}

func RegisterAlarm(kafkaConfig KafkaConf, thresholds ...*Threshold) (*Alarm, error) {
	a := &Alarm{
		ThresholdMap: make(map[pb.ThresholdName]*Threshold),
	}
	a.initKafka(kafkaConfig)

	for _, threshold := range thresholds {
		if _, ok := baseThresholdMap[threshold.GetName()]; !ok {
			return nil, fmt.Errorf("unknown threshold name :%s", threshold.GetName())
		}
		data, err := proto.Marshal(&pb.RegisterThreshold{
			BaseThreshold: threshold.BaseThreshold,
			Value:         threshold.Value,
			SendMail:      threshold.SendMail,
			Enabled:       threshold.Enabled,
		})
		if err != nil {
			return nil, fmt.Errorf("register threshold %s marshal failed: %s ",
				threshold.GetName(), err.Error())
		}
		if err := a.sendThresholdMessage(data, RegisterThreshold); err != nil {
			return nil, fmt.Errorf("register threshold failed:%s", err.Error())
		}

		a.ThresholdMap[threshold.GetName()] = threshold
	}

	go a.listenThreshold()
	return a, nil
}

func (a *Alarm) initKafka(kafkaConfig KafkaConf) {
	a.thresholdReader = kg.NewReader(kg.ReaderConfig{
		Brokers:  kafkaConfig.Addresses,
		Topic:    kafkaConfig.Topic,
		GroupID:  kafkaConfig.GroupId,
		MinBytes: 10,
		MaxBytes: 10e6,
		Dialer: &kg.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
			SASLMechanism: plain.Mechanism{
				Username: kafkaConfig.Username,
				Password: kafkaConfig.Password,
			},
			TLS: &tls.Config{InsecureSkipVerify: true},
		},
	})

	a.alarmWriter = &kg.Writer{
		Transport: &kg.Transport{
			SASL: plain.Mechanism{
				Username: kafkaConfig.Username,
				Password: kafkaConfig.Password,
			},
			TLS: &tls.Config{InsecureSkipVerify: true},
		},
		Addr:      kg.TCP(kafkaConfig.Addresses...),
		BatchSize: 1,
		Balancer:  &kg.LeastBytes{},
	}

	a.thresholdWriter = &kg.Writer{
		Transport: &kg.Transport{
			SASL: plain.Mechanism{
				Username: kafkaConfig.Username,
				Password: kafkaConfig.Password,
			},
			TLS: &tls.Config{InsecureSkipVerify: true},
		},
		Addr:      kg.TCP(kafkaConfig.Addresses...),
		BatchSize: 1,
		Balancer:  &kg.LeastBytes{},
	}
}

func (a *Alarm) Close() {
	if a.alarmWriter != nil {
		a.alarmWriter.Close()
	}
	if a.thresholdWriter != nil {
		a.thresholdWriter.Close()
	}
}

func (a *Alarm) listenThreshold() {
	defer a.thresholdReader.Close()

	for {
		message, err := a.thresholdReader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("read threshold message from kg failed: %s\n", err.Error())
			continue
		}

		switch string(message.Key) {
		case UpdateThreshold:
			var req pb.UpdateThreshold
			if err := proto.Unmarshal(message.Value, &req); err != nil {
				fmt.Printf("unmarshal update threshold failed: %s\n", err.Error())
			} else {
				a.updateThreshold(&req)
			}
		}
	}
}

func (a *Alarm) updateThreshold(req *pb.UpdateThreshold) {
	a.Lock()
	defer a.Unlock()

	if _, ok := a.ThresholdMap[req.GetName()]; !ok {
		return
	}
	a.ThresholdMap[req.GetName()].Value = req.GetValue()
	a.ThresholdMap[req.GetName()].SendMail = req.GetSendMail()
	a.ThresholdMap[req.GetName()].Enabled = req.GetEnabled()
}

func (a *Alarm) sendThresholdMessage(data []byte, cmd string) error {
	return a.thresholdWriter.WriteMessages(context.Background(),
		kg.Message{Topic: ThresholdTopic, Key: []byte(cmd), Value: data})
}
