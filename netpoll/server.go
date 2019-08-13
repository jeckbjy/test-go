package main

import (
	"github.com/mailru/easygo/netpoll"
	"log"
	"net"
	"os"
)

func main() {
	log.Printf("start server")

	p, err := netpoll.New(nil)
	if err != nil {
		return
	}

	l, err := net.Listen("tcp", ":6789")
	if err != nil {
		return
	}
	desc, err := netpoll.HandleListener(l, netpoll.EventRead)
	if err != nil {
		return
	}
	data := make([]byte, 1024)

	p.Start(desc, func(event netpoll.Event) {
		if event&netpoll.EventReadHup != 0 {
			p.Stop(desc)
			return
		}

		log.Printf("accept\n")
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
		}
		log.Printf("new conn\n")
		dd, err := netpoll.HandleRead(conn)
		if err == nil {
			p.Start(dd, func(ev netpoll.Event) {
				log.Printf("ev:%+v\n", ev)
				if ev&netpoll.EventReadHup != 0 {
					p.Stop(dd)
					return
				}

				log.Printf("read data\n")
				n, _ := conn.Read(data)
				//data, eee := ioutil.ReadAll(conn)
				//if eee != nil {
				//	log.Printf("aaa:%+v\n", eee)
				//}
				log.Printf("client:%s", data[:n])
				conn.Write([]byte("pong"))
			})
		}
	})

	log.Printf("ad\n")
	cc := make(chan os.Signal, 1)
	<-cc
}
