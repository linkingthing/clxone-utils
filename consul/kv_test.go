package consul

import (
	"testing"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewKV(t *testing.T) {
	c, err := consulapi.NewClient(&consulapi.Config{
		Address:   "10.0.0.68:28500",
		Scheme:    "https",
		Token:     "b36ab7f4-7ca4-e97c-5f82-34df4f7ef38b",
		TLSConfig: consulapi.TLSConfig{InsecureSkipVerify: true},
	})
	require.NoError(t, err)

	kv := c.KV()
	n := &KNode{}
	ok, err := GetKV[*KNode](kv, "node", nil, n)
	assert.NoError(t, err)
	if ok {
		t.Log("before", n)
	} else {
		t.Log("before set there is no record")
	}

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
	assert.NoError(t, SetKV[*KNode](kv, nil, "node/"+n1.Ipv4, n1))
	assert.NoError(t, SetKV[*KNode](kv, nil, "node/"+n2.Ipv4, n2))

	n = &KNode{}
	ok, err = GetKV[*KNode](kv, "node/"+n1.Ipv4, nil, n)
	assert.NoError(t, err)
	if ok {
		t.Log("after", n)
	}

	kNodes, err := ListKVs[*KNode](kv, "node", nil)
	assert.NoError(t, err)
	for _, kNode := range kNodes {
		t.Log("list node", kNode)
	}

	assert.NoError(t, DeleteKV(kv, "node/"+n1.Ipv4, nil))
	kNodes, err = ListKVs[*KNode](kv, "node", nil)
	assert.NoError(t, err)
	t.Log("after delete key", len(kNodes))

	assert.NoError(t, DeleteKVTree(kv, "node", nil))
	kNodes, err = ListKVs[*KNode](kv, "node", nil)
	assert.NoError(t, err)
	t.Log("after delete prefix", len(kNodes))
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
