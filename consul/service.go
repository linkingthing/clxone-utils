package consul

import "fmt"

type serviceNode struct {
	AggregatedStatus string  `json:"AggregatedStatus"`
	Service          Service `json:"Service"`
}

type Service struct {
	ID      string   `json:"ID"`
	Service string   `json:"Service"`
	Address string   `json:"Address"`
	Port    int      `json:"Port"`
	Tags    []string `json:"Tags"`
}

func (c *Consul) GetServices(name string, localIP string) ([]Service, error) {
	var consulAddr string
	for _, addr := range c.Address {
		if addr == localIP {
			consulAddr = addr
			break
		}
	}

	if consulAddr == "" {
		consulAddr = c.Address[0]
	}

	var serviceNodes []serviceNode
	if err := c.httpClient.Get(fmt.Sprintf("http://%s:%d/v1/agent/health/service/name/%s",
		consulAddr, c.Port, name), &serviceNodes); err != nil {
		return nil, fmt.Errorf("get %s service failed:%s", name, err.Error())
	}

	var aliveService []Service
	for _, node := range serviceNodes {
		if node.AggregatedStatus == "passing" {
			aliveService = append(aliveService, node.Service)
		}
	}

	return aliveService, nil
}
