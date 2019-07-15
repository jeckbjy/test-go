package consul

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"registry"
	"runtime"
	"time"

	consul "github.com/hashicorp/consul/api"
)

func configure(opts *registry.Options) *consul.Config {
	// use default config
	config := consul.DefaultConfig()

	//if opts.Context != nil {
	//	// Use the consul config passed in the options, if available
	//	if co, ok := opts.Context.Value("consul_config").(*consul.Config); ok {
	//		config = co
	//	}
	//	if cn, ok := opts.Context.Value("consul_connect").(bool); ok {
	//		c.connect = cn
	//	}
	//
	//	// Use the consul query options passed in the options, if available
	//	if qo, ok := c.opts.Context.Value("consul_query_options").(*consul.QueryOptions); ok && qo != nil {
	//		c.queryOptions = qo
	//	}
	//	if as, ok := c.opts.Context.Value("consul_allow_stale").(bool); ok {
	//		c.queryOptions.AllowStale = as
	//	}
	//}

	// check if there are any addrs
	if len(opts.Addrs) > 0 {
		addr, port, err := net.SplitHostPort(opts.Addrs[0])
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			port = "8500"
			addr = opts.Addrs[0]
			config.Address = fmt.Sprintf("%s:%s", addr, port)
		} else if err == nil {
			config.Address = fmt.Sprintf("%s:%s", addr, port)
		}
	}

	if config.HttpClient == nil {
		config.HttpClient = new(http.Client)
	}

	// requires secure connection?
	if opts.Secure || opts.TLSConfig != nil {

		config.Scheme = "https"
		// We're going to support InsecureSkipVerify
		config.HttpClient.Transport = newTransport(opts.TLSConfig)
	}

	// set timeout
	if opts.Timeout > 0 {
		config.HttpClient.Timeout = opts.Timeout
	}

	return config
}

func newTransport(config *tls.Config) *http.Transport {
	if config == nil {
		config = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     config,
	}
	runtime.SetFinalizer(&t, func(tr **http.Transport) {
		(*tr).CloseIdleConnections()
	})
	return t
}
