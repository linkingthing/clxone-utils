package consul

import (
	"context"
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
