package main

import (
	"container/list"
	"log"
	"net"
	"sync"
	"time"
)

const (
	listenAddr            = ":5678"
	defaultPassInterval   = time.Millisecond * 100 // 每隔固定时间放发放1个token
	defaultNotifyInterval = time.Second            // 每隔固定时间刷新所有序列
	defaultUrgentNum      = 10                     // 立即通知前n个排队序号
)

func NewServer() *Server {
	s := &Server{passInterval: defaultPassInterval, notifyInterval: defaultNotifyInterval}
	return s
}

// Server 定时
type Server struct {
	passInterval   time.Duration // 每隔固定时间放行1个
	notifyInterval time.Duration // 固定时间间隔通知客户端
	mux            sync.Mutex    //
	queue          *list.List    // 等待队列
	cursor         *list.Element // 当前通知游标
	listener       net.Listener  // 监听socket
	quit           chan struct{} // 退出
}

func (s *Server) Run() error {
	s.quit = make(chan struct{})
	s.queue = list.New()
	if err := s.serve(); err != nil {
		return err
	}

	go s.process()
	go s.notify()

	log.Printf("queue service start")
	return nil
}

func (s *Server) Stop() {
	_ = s.listener.Close()
	close(s.quit)
	log.Printf("queue service stop")
}

func (s *Server) serve() error {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	s.listener = l

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			}
			c := NewClient(conn)
			s.mux.Lock()
			s.queue.PushBack(c)
			c.SendMessage(s.queue.Len())
			s.mux.Unlock()
		}
	}()
	return nil
}

func (s *Server) process() {
	t := time.NewTicker(s.passInterval)
	for {
		select {
		case <-t.C:
			// 定时放行1个
			//log.Printf("pass one\n")
			s.mux.Lock()
			if s.queue.Len() > 0 {
				elem := s.queue.Front()
				c := elem.Value.(*Client)
				c.SendMessage(0)
				s.queue.Remove(elem)
				if elem == s.cursor {
					s.cursor = nil
				}
				// 立刻通知前n个
				index := 1
				for iter := s.queue.Front(); iter != nil; iter = iter.Next() {
					if index > defaultUrgentNum {
						break
					}
					client := iter.Value.(*Client)
					client.SendMessage(index)
					index++
				}
			}
			s.mux.Unlock()
		case <-s.quit:
			return
		}
	}
}

// 定期通知客户端
func (s *Server) notify() {
	t := time.NewTicker(s.notifyInterval)
	for {
		select {
		case <-t.C:
			// 通知客户端当前队列
			//log.Printf("update index\n")
			s.mux.Lock()
			s.cursor = nil
			if s.queue.Len() > 0 {
				s.cursor = s.queue.Front()
			}
			cursor := s.cursor
			index := 1
			s.mux.Unlock()
			for cursor != nil {
				err := cursor.Value.(*Client).SendMessage(index)
				index++
				s.mux.Lock()
				if s.cursor == nil {
					// 说明被重置了,重新等待下次通知
					cursor = nil
				} else {
					temp := cursor
					cursor = cursor.Next()
					s.cursor = cursor
					if err != nil {
						// 删除失效的socket
						s.queue.Remove(temp)
					}
				}
				s.mux.Unlock()
			}

		case <-s.quit:
			return
		}
	}
}
