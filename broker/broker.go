package broker

// TODO:事务支持,需要么?
type IBroker interface {
	Dial(...Option) error
	Close()
	Publish(topic string, v interface{}, opts ...PublishOption) error
	Subscribe(topic string, opts ...SubscribeOption) (ISubscriber, error)
	String() string
}

type ISubscriber interface {
	Next() (chan *Message, error)
	Ack() error
	Unsubscribe() error
}

type IMarshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
	String() string
}

type Message struct {
	Topic string
	Reply string
	Data  []byte
}
