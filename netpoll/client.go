package main

import (
	"github.com/mailru/easygo/netpoll"
	"log"
	"net"
	"os"
)

func main() {
	log.Printf("start client")
	p, err := netpoll.New(nil)
	if err != nil {
		return
	}
	conn, err := net.Dial("tcp", "localhost:6789")
	if err != nil {
		return
	}

	conn.Write([]byte("ping"))

	desc, err := netpoll.HandleRead(conn)
	if err != nil {
		return
	}

	data := make([]byte, 1024)

	p.Start(desc, func(event netpoll.Event) {
		if event&netpoll.EventReadHup != 0 {
			p.Stop(desc)
			log.Printf("close")
			return
		}

		n, _ := conn.Read(data)
		//data, err := ioutil.ReadAll(conn)
		//if err != nil {
		//	return
		//}
		log.Printf("server:%s", data[:n])
		conn.Write([]byte("ping"))
	})

	log.Printf("aa")
	cc := make(chan os.Signal, 1)
	<-cc
}
