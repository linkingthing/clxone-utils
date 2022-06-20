package consul

import (
	"context"
	"io"
	"math"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func NewGrpcConn(conf *consulapi.Config, serviceName string, tags ...string) (*grpc.ClientConn, error) {
	endPointer, err := getDefaultEndPointer(conf, serviceName, tags...)
	if err != nil {
		return nil, err
	}

	defer endPointer.Close()
	balancer := lb.NewRoundRobin(endPointer)
	roundRobinEndPoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}

	response, err := roundRobinEndPoint(context.Background(), struct{}{})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, response.(string), grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetEndpoints(conf *consulapi.Config, serviceName string) ([]endpoint.Endpoint, error) {
	endPointer, err := getDefaultEndPointer(conf, serviceName)
	if err != nil {
		return nil, err
	}

	defer endPointer.Close()
	return endPointer.Endpoints()
}

func getDefaultEndPointer(conf *consulapi.Config, serviceName string, tags ...string) (*sd.DefaultEndpointer, error) {
	apiClient, err := consulapi.NewClient(conf)
	if err != nil {
		return nil, err
	}

	client := consul.NewClient(apiClient)
	logger := kitlog.With(kitlog.NewLogfmtLogger(os.Stderr), "timestamp", kitlog.DefaultTimestampUTC)
	instance := consul.NewInstancer(client, logger, serviceName, tags, true)
	defer instance.Stop()

	return sd.NewEndpointer(instance, getFactory, logger), nil
}

func getFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	return func(context.Context, interface{}) (interface{}, error) {
		return instance, nil
	}, nil, nil
}
