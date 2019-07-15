package consul

import (
	"errors"
	"registry"
	"sync"

	consul "github.com/hashicorp/consul/api"
)

func init() {

}

type cregistry struct {
	sync.Mutex
	client       *consul.Client
	opts         registry.Options
	connect      bool
	queryOptions *consul.QueryOptions
	address      string
}

func (c *cregistry) Init(opts ...registry.Option) error {
	// set opts
	for _, o := range opts {
		o(&c.opts)
	}

	// use default config
	config := consul.DefaultConfig()

	if c.opts.Context != nil {
		c.opts.Bind("consul_config", &config)
		c.opts.Bind("consul_connect", &c.connect)
		c.opts.Bind("consul_query_options", &c.queryOptions)
		c.opts.Bind("consul_allow_stale", &c.queryOptions.AllowStale)
	}

	//c.opts.Bind(
	//	"consul_config", &config,
	//	"consul_connect", &c.connect,
	//	"consul_query_options", &c.queryOptions,
	//	"consul_allow_stale", &c.queryOptions.AllowStale
	//	)
	//c.opts.Bind("consul_config", &c.connect, "consul_connect")

	// create the client
	client, _ := consul.NewClient(config)

	// set address/client
	c.address = config.Address
	c.client = client
	return nil
}

func (c *cregistry) Register(s *registry.Service, opts ...registry.RegisterOption) error {
	// use first node
	node := s.Nodes[0]

	// get existing hash and last checked time
	//c.Lock()
	//v, ok := c.register[s.Name]
	//lastChecked := c.lastChecked[s.Name]
	//c.Unlock()

	tags := []string{"rails"}
	var check *consul.AgentServiceCheck

	// register the service
	asr := &consul.AgentServiceRegistration{
		ID:      node.Id,
		Name:    s.Name,
		Tags:    tags,
		Port:    node.Port,
		Address: node.Address,
		Check:   check,
	}

	if err := c.client.Agent().ServiceRegister(asr); err != nil {
		return err
	}

	// pass the healthcheck
	return c.client.Agent().PassTTL("service:"+node.Id, "")
}

func (c *cregistry) Deregister(s *registry.Service) error {
	if len(s.Nodes) == 0 {
		return errors.New("Require at least one node")
	}

	// delete our hash and time check of the service
	//c.Lock()
	//delete(c.register, s.Name)
	//delete(c.lastChecked, s.Name)
	//c.Unlock()

	node := s.Nodes[0]
	return c.client.Agent().ServiceDeregister(node.Id)
}

func (c *cregistry) Query(name string) ([]*registry.Service, error) {
	var rsp []*consul.ServiceEntry
	var err error

	// if we're connect enabled only get connect services
	if c.connect {
		rsp, _, err = c.client.Health().Connect(name, "", false, c.queryOptions)
	} else {
		rsp, _, err = c.client.Health().Service(name, "", false, c.queryOptions)
	}
	if err != nil {
		return nil, err
	}

	serviceMap := map[string]*registry.Service{}

	for _, s := range rsp {
		if s.Service.Service != name {
			continue
		}

		// version is now a tag
		//version, _ := decodeVersion(s.Service.Tags)
		version := ""
		// service ID is now the node id
		id := s.Service.ID
		// key is always the version
		key := version

		// address is service address
		address := s.Service.Address

		// use node address
		if len(address) == 0 {
			address = s.Node.Address
		}

		svc, ok := serviceMap[key]
		if !ok {
			svc = &registry.Service{
				//Endpoints: decodeEndpoints(s.Service.Tags),
				Name:    s.Service.Service,
				Version: version,
			}
			serviceMap[key] = svc
		}

		var del bool

		for _, check := range s.Checks {
			// delete the node if the status is critical
			if check.Status == "critical" {
				del = true
				break
			}
		}

		// if delete then skip the node
		if del {
			continue
		}

		svc.Nodes = append(svc.Nodes, &registry.Node{
			Id:      id,
			Address: address,
			Port:    s.Service.Port,
			//Metadata: decodeMetadata(s.Service.Tags),
		})
	}

	var services []*registry.Service
	for _, service := range serviceMap {
		services = append(services, service)
	}
	return services, nil
}
