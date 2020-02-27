package main

import (
	"flag"
	"log"
	"net"
	"strings"
	"sync"
)

var showIndex = false

func main() {
	np := flag.Int("n", 100, "client max count")
	showIndexP := flag.Bool("s", false, "show index")
	flag.Parse()
	showIndex = *showIndexP

	wg := &sync.WaitGroup{}
	for i := 0; i < *np; i++ {
		c := &Client{id: i + 1, wg: wg}
		c.Run()
	}

	wg.Wait()
}

type Client struct {
	id   int
	wg   *sync.WaitGroup
	conn net.Conn
}

func (c *Client) Run() {
	c.wg.Add(1)
	conn, err := net.Dial("tcp", "localhost:5678")
	if err != nil {
		log.Printf("dial fail,%+v\n", err)
		c.wg.Done()
		return
	}
	c.conn = conn
	go c.loop()
}

func (c *Client) loop() {
	data := ""
	for {
		temp := make([]byte, 128)
		n, err := c.conn.Read(temp)
		if err != nil {
			c.wg.Done()
			break
		}

		data += string(temp[:n])
		for len(data) != 0 {
			index := strings.IndexByte(data, '\n')
			if index == -1 {
				break
			}

			msg := data[:index]
			data = data[index+1:]
			tokens := strings.SplitN(msg, " ", 2)
			queueID := tokens[0]
			token := ""
			if queueID == "0" {
				if len(tokens) > 1 {
					token = tokens[1]
				}
				log.Printf("conn gain token:%+v,%+v\n", c.id, token)
				c.wg.Done()
				return
			} else {
				// 选择是否需要展示客户端的当前排队ID,如果打开会一直刷屏
				if showIndex {
					log.Printf("conn wait index:%+v,%+v\n", c.id, queueID)
				}
			}
		}
	}
}
