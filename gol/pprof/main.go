package main

import (
	"fmt"
	"time"

	"net/http"
	_ "net/http/pprof"
)

var catCh chan string

func main() {

	catCh = make(chan string)

	go func() {
		for {
			catCh <- ""
		}
	}()

	go loop()

	http.ListenAndServe("0.0.0.0:19900", nil)
}

type Cat struct {
	Name string
	Ls   []byte
}

func loop() {
	index := 0
	for {
		index = index + 1
		str1 := fmt.Sprintf("index: %d", index)
		select {
		case cc := <-catCh:
			var c Cat
			c.Name = str1 + cc
			c.Ls = make([]byte, 1024*1024*1024)
			go handle(c)
			// time.Sleep(1 * time.Second)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func handle(cat Cat) {
	fmt.Println(cat.Name)
}

/*
1. CPU profiling（CPU 性能分析）：这是最常使用的一种类型。用于分析函数或方法的执行耗时；
2. Memory profiling：这种类型也常使用。用于分析程序的内存占用情况；
3. Block profiling：这是 Go 独有的，用于记录 goroutine 在等待共享资源花费的时间；
4. Mutex profiling：与 Block profiling 类似，但是只记录因为锁竞争导致的等待或延迟。
*/

// http://127.0.0.1:19900/debug/pprof/
// go tool pprof --text http://localhost:19900/debug/pprof/heap
// go tool pprof --text http://localhost:19900/debug/pprof/profile
// go tool pprof --text http://localhost:19900/debug/pprof/block
// go tool pprof /home/sh/pprof/pprof.loop.alloc_objects.alloc_space.inuse_objects.inuse_space.022.pb.gz
// web
// go tool pprof -svg http://192.168.56.101:8080/debug/pprof/profile >  cpu.svg
// go tool pprof -svg ./pprof_runtime cpu.pprof.201801301415 > cpu.svg

// Total：总共采样次数，这里是2525次。
// Flat：函数在样本中处于运行状态的次数。简单来说就是函数出现在栈顶的次数，而函数在栈顶则意味着它在使用CPU。
// Flat%：Flat / Total。
// Sum%：自己以及所有前面的Flat%的累积值。解读方式：表中第3行Sum% 32.4%，意思是前3个函数（运行状态）的计数占了总样本数的32.4%
// Cum：函数在样本中出现的次数。只要这个函数出现在栈中那么就算进去，这个和Flat不同（必须是栈顶才能算进去）。也可以解读为这个函数的调用次数。
// Cum%：Cum / Total

// go tool pprof -http=:8081 ~/pprof_file

// go tool pprof [binary] [source]

// - 一个有用的命令是 topN，它列出最耗时间的地方

// top10

// allocs	内存分配情况的采样信息
// blocks	阻塞操作情况的采样信息
// cmdline	显示程序启动命令参数及其参数
// goroutine	显示当前所有协程的堆栈信息
// heap	堆上的内存分配情况的采样信息
// mutex	锁竞争情况的采样信息
// profile	cpu占用情况的采样信息，点击会下载文件
// threadcreate	系统线程创建情况的采样信息
// trace	程序运行跟踪信息
