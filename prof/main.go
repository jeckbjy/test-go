package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

// demo1:https://www.jianshu.com/p/c20d8b250f00
// demo2:https://blog.csdn.net/moxiaomomo/article/details/77096814
// 其他:https://www.cnblogs.com/li-peng/p/9391543.html
// 安装Graphviz: http://macappstore.org/graphviz-2/
func main() {
	demo2()
}

func hello() {
	for {
		go func() {
			fmt.Println("hello word")
		}()
		time.Sleep(time.Millisecond * 1)
	}
}

func demo1() {
	// 开启pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	go hello()
	select {}
}

func Counter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 200)
		counter++
	}
	wg.Done()
}

func demo2() {
	flag.Parse()

	//远程获取pprof数据
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()

	// sleep 10mins, 在程序退出之前可以查看性能参数.
	time.Sleep(60 * time.Second)
}
