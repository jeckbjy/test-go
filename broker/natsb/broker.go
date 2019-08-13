package natsb

import (
	"errors"
	"github.com/nats-io/nats.go"
	"sync"
	"test-go/broker"
)

type Broker struct {
	opts  broker.Options
	nopts nats.Options
	conn  *nats.Conn
	drain bool
	mux   sync.Mutex
}

func (b *Broker) Dial(opts ...broker.Option) error {
	b.mux.Lock()
	defer b.mux.Unlock()

	status := nats.CLOSED
	if b.conn != nil {
		status = b.conn.Status()
	}

	switch status {
	case nats.CONNECTED, nats.RECONNECTING, nats.CONNECTING:
		return nil
	default:
		// DISCONNECTED or CLOSED or DRAINING
		c, err := b.nopts.Connect()
		if err != nil {
			return err
		}
		b.conn = c
		return nil
	}
}

func (b *Broker) Close() error {
	b.mux.Lock()
	defer b.mux.Unlock()
	if b.drain {
		_ = b.conn.Drain()
	}
	b.conn.Close()
	return nil
}

func (b *Broker) Publish(topic string, v interface{}, opts ...broker.PublishOption) error {
	var data []byte
	switch v.(type) {
	case []byte:
		data = v.([]byte)
	case string:
		data = []byte(v.(string))
	default:
		var err error
		data, err = b.opts.Codec.Marshal(v)
		if err != nil {
			return err
		}
	}

	return b.conn.Publish(topic, data)
}

func (b *Broker) Subscribe(topic string, opts ...broker.SubscribeOption) (broker.ISubscriber, error) {
	if b.conn == nil {
		return nil, errors.New("not connected")
	}

	conf := broker.SubscribeOptions{}

	var sub *nats.Subscription
	var err error
	// 使用同步阻塞调用
	if len(conf.Queue) > 0 {
		//b.conn.QueueSubscribeSyncWithChan(topic, conf.Queue, )
		sub, err = b.conn.QueueSubscribeSync(topic, conf.Queue)
	} else {
		sub, err = b.conn.SubscribeSync(topic)
	}

	if err != nil {
		return nil, err
	}

	return &Subscriber{sub: sub}, nil
}
