package consul

import (
	"fmt"

	consulsd "github.com/go-kit/kit/sd/consul"
	consulapi "github.com/hashicorp/consul/api"
)

var defaultCheck = &consulapi.AgentServiceCheck{
	Interval:                       "10s",
	Timeout:                        "3s",
	DeregisterCriticalServiceAfter: "168h",
	TLSSkipVerify:                  true,
}

func RegisterHttp(conf *consulapi.Config, registration consulapi.AgentServiceRegistration) (*Registrar, error) {
	if check := registration.Check; check == nil {
		check = defaultCheck
		check.GRPC = ""
		check.HTTP = fmt.Sprintf("http://%v:%v/health", registration.Address, registration.Port)
		registration.Check = check
	}

	return register(conf, registration)
}

func RegisterGrpc(conf *consulapi.Config, registration consulapi.AgentServiceRegistration) (*Registrar, error) {
	if check := registration.Check; check == nil {
		check = defaultCheck
		check.HTTP = ""
		check.GRPC = fmt.Sprintf("%v:%v", registration.Address, registration.Port)
		registration.Check = check
	}

	return register(conf, registration)
}

func register(conf *consulapi.Config, registration consulapi.AgentServiceRegistration) (*Registrar, error) {
	consulClient, err := consulapi.NewClient(conf)
	if err != nil {
		return nil, fmt.Errorf("new consul client failed: %s", err.Error())
	}

	return NewRegistrar(consulsd.NewClient(consulClient), &registration), nil
}

// Registrar registers service instance liveness information to Consul.
type Registrar struct {
	client       consulsd.Client
	registration *consulapi.AgentServiceRegistration
}

// NewRegistrar returns a Consul Registrar acting on the provided catalog
// registration.
func NewRegistrar(client consulsd.Client, r *consulapi.AgentServiceRegistration) *Registrar {
	return &Registrar{
		client:       client,
		registration: r,
	}
}

// Register implements sd.Registrar interface.
func (p *Registrar) Register() error {
	return p.client.Register(p.registration)
}

// Deregister implements sd.Registrar interface.
func (p *Registrar) Deregister() error {
	return p.client.Deregister(p.registration)
}
