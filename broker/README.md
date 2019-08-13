# message queue
常见的消息队列:ActiveMQ，RabbitMQ，ZeroMQ，Kafka，MetaMQ，RocketMQ，kafaka,nats

- 消息队列常见模型
  - 点对点模式:一个消息只能由一个消费者执行
  - Pub/Sub模式:一个消息可以有多个消费者执行,支持分组，同一组内是点对点模式

- nats  
nsts的设计是Subject-based,consumer通过subject通配符去绑定主题,通过指定queue去实现工作队列,相同的Queue组内,同一个消息只会被一个消费  
这种设计可能会有一些效率问题,直接在跟目录上进行通配符匹配  
还有一种问题,消息队列的分发模式是由Consumer决定的,但是有些时候并不是完全由客户端决定,比如一个GM系统，有些消息希望由某个消费者执行，
而有些消息则是希望所有消费者都能执行,这就需要定义两种队列才能实现,

nats的api设计比较简单，publish不用关心消费模式

- RabbitMQ