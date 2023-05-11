package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go f1()
	go f1()

	for {

		buf := make([]byte, 1024)
		runtime.Stack(buf, true)
		fmt.Printf("%s", string(buf))
		time.Sleep(1 * time.Second)
	}
}

func f1() {

	i := 0
	for {
		i = i + 1
		if i > 10 {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
