package addr

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestAddr(t *testing.T) {
	addr, err := Extract("")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(addr)
	}
}

func TestEcho(t *testing.T) {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	go func() {
		c, err := l.Accept()
		if err != nil {
			return
		}

		c.Write([]byte("Ping"))
		var rr [100]byte
		n, err := c.Read(rr[:])
		if err != nil {
			t.Error("aa")
			return
		}
		t.Logf("%s", rr[:n])
	}()

	addr, err := Extract("")
	if err != nil {
		return
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%+v", addr, "8080"))
	if err != nil {
		t.Error(err)
		return
	}

	var bytes [100]byte
	n, err := conn.Read(bytes[:])
	if err != nil {
		return
	}

	t.Logf("%s", bytes[:n])

	conn.Write([]byte("Pong"))

	time.Sleep(10000)
}
