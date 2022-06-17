package consul

import (
	"fmt"

	"github.com/linkingthing/clxone-utils/http"
)

type register struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Tags    []string  `json:"tags"`
	Address string    `json:"address"`
	Port    int       `json:"port"`
	Check   CheckConf `json:"check"`
}

type CheckConf struct {
	Interval                       string `json:"interval"`
	Timeout                        string `json:"timeout"`
	DeregisterCriticalServiceAfter string `json:"deregister_critical_service_after"`
	Http                           string `json:"http"`
	GRPC                           string `json:"grpc"`
	TLSSkipVerify                  bool   `json:"tls_skip_verify"`
}

type RegisterService struct {
	IP       string
	Name     string
	Tags     []string
	HttpPort int
	GrpcPort int
}

const (
	defaultInterval                       = "10s"
	defaultTimeout                        = "3s"
	defaultDeregisterCriticalServiceAfter = "168h"
)

type Consul struct {
	Address []string
	Port    int
	CheckConf
	httpClient     *http.Client
	registeredGrpc bool
}

func NewConsul(addresses []string, port int) *Consul {
	c := &Consul{
		Address:    addresses,
		Port:       port,
		httpClient: http.NewHttpClient(),
		CheckConf: CheckConf{
			Interval:                       defaultInterval,
			Timeout:                        defaultTimeout,
			DeregisterCriticalServiceAfter: defaultDeregisterCriticalServiceAfter,
		},
	}

	return c
}

func (c *Consul) SetCheckConfig(config CheckConf) *Consul {
	c.CheckConf = config
	return c
}

func (c *Consul) RegisterAPIService(s RegisterService) error {
	name := s.Name + "-api"
	consul := register{
		ID:      name + "-" + s.IP,
		Name:    name,
		Tags:    s.Tags,
		Address: s.IP,
		Port:    s.HttpPort,
		Check: CheckConf{
			Interval:                       c.Interval,
			Timeout:                        c.Timeout,
			DeregisterCriticalServiceAfter: c.DeregisterCriticalServiceAfter,
			Http:                           fmt.Sprintf("http://%s:%d/health", s.IP, s.HttpPort),
		},
	}

	for _, addr := range c.Address {
		if err := c.httpClient.Put(
			fmt.Sprintf("http://%s:%d/v1/agent/service/register?replace-existing-checks=true", addr, c.Port),
			consul, nil); err != nil {
			return fmt.Errorf("register service %s failed:%s", s.Name, err.Error())
		}
	}

	c.registeredGrpc = true
	return nil
}

func (c *Consul) RegisterGRPCService(s RegisterService) error {
	name := s.Name + "-grpc"
	consul := register{
		ID:      name + "-" + s.IP,
		Name:    name,
		Tags:    s.Tags,
		Address: s.IP,
		Port:    s.GrpcPort,
		Check: CheckConf{
			Interval:                       c.Interval,
			Timeout:                        c.Timeout,
			DeregisterCriticalServiceAfter: c.DeregisterCriticalServiceAfter,
			GRPC:                           fmt.Sprintf("%s:%d/%s", s.IP, s.GrpcPort, s.Name),
		},
	}

	for _, addr := range c.Address {
		if err := c.httpClient.Put(
			fmt.Sprintf("http://%s:%d/v1/agent/service/register?replace-existing-checks=true", addr, c.Port),
			consul, nil); err != nil {
			return fmt.Errorf("register service %s failed:%s", s.Name, err.Error())
		}
	}

	return nil
}

func (c *Consul) DeRegister(s RegisterService) error {
	var response interface{}
	for _, addr := range c.Address {
		if s.HttpPort != 0 {
			if err := c.httpClient.Put(
				fmt.Sprintf("http://%s:%d/v1/agent/service/deregister/%s", addr, c.Port, s.Name+"-api-"+s.IP),
				nil, response); err != nil {
				fmt.Printf("deregister service %s failed:%s\n", s.Name, err.Error())
			}
		}
		if s.GrpcPort != 0 {
			if err := c.httpClient.Put(
				fmt.Sprintf("http://%s:%d/v1/agent/service/deregister/%s", addr, c.Port, s.Name+"-grpc-"+s.IP),
				nil, response); err != nil {
				fmt.Printf("deregister service %s failed:%s\n", s.Name, err.Error())
			}
		}
	}

	return nil
}
