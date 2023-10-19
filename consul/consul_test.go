package consul

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"testing"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

func TestRegisterHttpService(t *testing.T) {
	if apiRegister, err := RegisterHttp(&consulapi.Config{
		Address:   "10.0.0.66:28500",
		Token:     "b36ab7f4-7ca4-e97c-5f82-34df4f7ef38b",
		TLSConfig: consulapi.TLSConfig{InsecureSkipVerify: true},
	}, consulapi.AgentServiceRegistration{
		ID:      "clxone-ipam-api-10.0.0.66",
		Name:    "clxone-ipam-api",
		Address: "10.0.0.66",
		Port:    28083,
	}); err != nil {
		t.Error(err)
		return
	} else {
		if err := apiRegister.Register(); err != nil {
			t.Error(err)
			return
		} else {
			defer apiRegister.Deregister()
		}
	}
}

func TestRegisterGRPCService(t *testing.T) {
	if grpcRegister, err := RegisterGrpc(&consulapi.Config{
		Address:   "10.0.0.66:28500",
		Token:     "b36ab7f4-7ca4-e97c-5f82-34df4f7ef38b",
		TLSConfig: consulapi.TLSConfig{InsecureSkipVerify: true},
	}, consulapi.AgentServiceRegistration{
		ID:      "clxone-ipam-grpc-10.0.0.66",
		Name:    "clxone-ipam-grpc",
		Address: "10.0.0.66",
		Port:    28883,
	}); err != nil {
		t.Error(err)
		return
	} else {
		if err := grpcRegister.Register(); err != nil {
			t.Error(err)
			return
		} else {
			defer grpcRegister.Deregister()
		}
	}
}

func TestNewGrpcConn(t *testing.T) {
	conn, err := NewGrpcConn(&consulapi.Config{
		Address:   "10.0.0.66:28500",
		Token:     "b36ab7f4-7ca4-e97c-5f82-34df4f7ef38b",
		TLSConfig: consulapi.TLSConfig{InsecureSkipVerify: true},
	}, "clxone-monitor-grpc")
	defer conn.Close()
	if err != nil {
		t.Error(err)
		return
	}

	//new protobuf Client
	//gMonitorGrpcClient = &MonitorGrpcClient{client: pbmonitor.NewMonitorServiceClient(conn)}
}

func TestGetEndpoints(t *testing.T) {
	endpoints, err := GetEndpoints(&consulapi.Config{
		Address:   "10.0.0.66:28500",
		Token:     "b36ab7f4-7ca4-e97c-5f82-34df4f7ef38b",
		TLSConfig: consulapi.TLSConfig{InsecureSkipVerify: true},
	}, "clxone-dhcp-agent-grpc")
	if err != nil {
		t.Error(err)
		return
	}

	for _, endpoint := range endpoints {
		response, err := endpoint(context.Background(), struct{}{})
		if err != nil {
			t.Error(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, response.(string), grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		//new protobuf Client and call grpc service method
		//client := pbdhcpagent.NewDHCPManagerClient(conn)
		//resp, err := client.GetDHCPNodes(ctx, &pbdhcpagent.GetDHCPNodesRequest{})
		//if err != nil {
		//	return nil, err
		//}
	}
}

func TestKV(t *testing.T) {
	conf := &consulapi.Config{
		Address:   "10.0.0.68:28500",
		Scheme:    "https",
		Token:     "b36ab7f4-7ca4-e97c-5f82-34df4f7ef38b",
		TLSConfig: consulapi.TLSConfig{InsecureSkipVerify: true},
	}
	apiClient, err := consulapi.NewClient(conf)
	if err != nil {
		t.Error(err)
		return
	}

	kv, meta, err := apiClient.KV().Get("node", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("before", kv, meta)

	n1 := &KNode{
		Mac:          "00:00:00:00:00:01",
		Ipv4:         "10.0.0.1",
		DeployType:   DeployTypeHa,
		Grouped:      1,
		ServiceRoles: []ServiceRole{ServiceRoleIpam, ServiceRoleDc},
		Version:      "v2.7.0",
		IsAlive:      true,
		HaVirtualIp:  "10.0.0.3",
		IsNodeHa:     true,
	}
	n2 := &KNode{
		Mac:          "00:00:00:00:00:02",
		Ipv4:         "10.0.0.2",
		DeployType:   DeployTypeHa,
		Grouped:      1,
		ServiceRoles: []ServiceRole{ServiceRoleIpam, ServiceRoleDc},
		Version:      "v2.7.0",
		IsAlive:      true,
		IsNodeHa:     true,
	}

	b1, err := json.Marshal(n1)
	if err != nil {
		t.Error(err)
		return
	}
	b2, err := json.Marshal(n2)
	if err != nil {
		t.Error(err)
		return
	}
	p1 := &consulapi.KVPair{Key: "node/" + n1.Ipv4, Value: b1}
	p2 := &consulapi.KVPair{Key: "node/" + n2.Ipv4, Value: b2}
	if _, err := apiClient.KV().Put(p1, nil); err != nil {
		t.Error(err)
		return
	}
	if _, err := apiClient.KV().Put(p2, nil); err != nil {
		t.Error(err)
		return
	}

	kv, meta, err = apiClient.KV().Get("node/"+n1.Ipv4, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("after", kv, meta)

	list, _, err := apiClient.KV().List("node", nil)
	if err != nil {
		t.Error(err)
		return
	}

	for k, pair := range list {
		var value *KNode
		if err := json.Unmarshal(pair.Value, &value); err != nil {
			t.Error(err)
			return
		}
		t.Log(k, pair.Key, value)
	}
}

type KNode struct {
	Mac          string        `json:"mac"`
	Name         string        `json:"name"`
	Ipv4         string        `json:"ipv4"`
	Ipv6         string        `json:"ipv6"`
	DeployType   DeployType    `json:"deployType"`
	Grouped      int           `json:"grouped"`
	ServiceRoles []ServiceRole `json:"serviceRoles"`
	Version      string        `json:"version"`
	HaVirtualIp  string        `json:"haVirtualIp"`
	IsNodeHa     bool          `json:"isNodeHa"`
	IsAlive      bool          `json:"isAlive"`
}

type ServiceRole = string

const (
	ServiceRoleIpam     ServiceRole = "ipam"
	ServiceRoleDhcp     ServiceRole = "dhcp"
	ServiceRoleDns      ServiceRole = "dns"
	ServiceRoleDc       ServiceRole = "dc"
	ServiceRoleFlow     ServiceRole = "flow"
	ServiceRoleDetector ServiceRole = "detector"
)

type DeployType string

const (
	DeployTypeSingle  DeployType = "singleton"
	DeployTypeHa      DeployType = "ha"
	DeployTypeAnycast DeployType = "anycast"
	DeployTypeCluster DeployType = "cluster"
)
