package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"sync/atomic"
)

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

type Client struct {
	conn      net.Conn
	lastIndex int32
}

func (c *Client) Send(msg []byte) error {
	_, err := c.conn.Write(msg)
	if err != nil {
		return err
	}

	return nil
}

// 发送消息,0:代表玩家获取到token,可以进入
// 大于0代表需要排队
// 编码格式: {index} {token}\n
func (c *Client) SendMessage(index int) error {
	if int(atomic.LoadInt32(&c.lastIndex)) == index {
		return nil
	}
	atomic.StoreInt32(&c.lastIndex, int32(index))

	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%+v", index))
	if index == 0 {
		b.WriteByte(' ')
		token := rand.Int31()
		b.WriteString(fmt.Sprintf("%+v", token))
	}
	b.WriteByte('\n')
	if err := c.Send(b.Bytes()); err != nil {
		return err
	}

	if index == 0 {
		return c.conn.Close()
	}
	return nil
}
