package registry

import (
	"context"
	"crypto/tls"
	"github.com/micro/go-micro/registry/consul"
	"reflect"
	"time"
)

type BaseOptions struct {
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

func (o *BaseOptions) Bind(key string, value interface{}) {
	//v := o.Context.Value(key)
	//t := reflect.TypeOf(value)
	//vv := reflect.ValueOf(v)
	//if t.Kind() == reflect.Ptr && t.Elem().Kind() ==  {
	//
	//}
}

type Options struct {
	BaseOptions
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config
}

type RegisterOptions struct {
	BaseOptions
	TTL time.Duration
}

type WatchOptions struct {
	BaseOptions
	// Specify a service to watch
	// If blank, the watch is for all services
	Service string
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)
