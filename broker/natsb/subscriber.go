package natsb

import (
	"errors"
	"github.com/nats-io/nats.go"
	"test-go/broker"
)

type Subscriber struct {
	sub  *nats.Subscription
	nmsg chan *nats.Msg
	bmsg chan *broker.Message
}

func (s *Subscriber) Next() (chan *broker.Message, error) {
	m, ok := <-s.nmsg
	if !ok {
		return nil, errors.New("close")
	}

	s.bmsg <- &broker.Message{Topic: m.Subject, Reply: m.Reply, Data: m.Data}

	return s.bmsg, nil
}

func (s *Subscriber) Ack() {
}

func (s *Subscriber) Unsubscribe() error {
	return s.sub.Unsubscribe()
}
