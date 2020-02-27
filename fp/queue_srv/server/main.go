package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/profile"
)

func main() {
	// 测试cpu开启此行
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	// 测试内存
	defer profile.Start(profile.MemProfileRate(128), profile.ProfilePath("."), profile.NoShutdownHook).Stop()

	s := NewServer()
	if err := s.Run(); err != nil {
		log.Printf("start server fail\n")
		return
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
	s.Stop()
}
