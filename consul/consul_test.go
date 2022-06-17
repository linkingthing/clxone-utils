package consul

import "testing"

func genTestConsul() (*Consul, RegisterService) {
	s := RegisterService{
		Name:     "clxone-ipam",
		Tags:     []string{},
		HttpPort: 28005,
		GrpcPort: 28885,
		IP:       "10.0.0.135",
	}

	return NewConsul([]string{"10.0.0.66"}, 28500), s
}

func TestRegisterAPIService(t *testing.T) {
	consul, s := genTestConsul()
	if err := consul.RegisterAPIService(s); err != nil {
		t.Error(err)
		return
	}
}

func TestRegisterGRPCService(t *testing.T) {
	consul, s := genTestConsul()
	if err := consul.RegisterGRPCService(s); err != nil {
		t.Error(err)
		return
	}
}

func TestDeRegister(t *testing.T) {
	consul, s := genTestConsul()
	if err := consul.DeRegister(s); err != nil {
		t.Error(err)
		return
	}
}

func TestGetServiceNodes(t *testing.T) {
	consul, _ := genTestConsul()
	nodes, err := consul.GetServices("clxone-dhcp-agent-grpc", "10.0.0.66")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%#v", nodes)
}
