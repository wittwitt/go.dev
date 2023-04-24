package recover1

import (
	"fmt"
	"runtime"
	"testing"
)

// https://stackoverflow.com/questions/57486620/are-all-runtime-errors-recoverable-in-go

func TestRecover1(t *testing.T) {
	recover2()
}

func TestRecover2(t *testing.T) {
	recover1()
}

// Out of memory
func recover1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	_ = make([]int64, 1<<40)
}

// read and write map in mutiple thread
func recover2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	m := map[string]int{}

	go func() {
		for {
			m["x"] = 1
		}
	}()
	for {
		_ = m["x"]
	}
}

func recover3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var f func(a [1000]int64)
	f = func(a [1000]int64) {
		f(a)
	}
	f([1000]int64{})

}

// Attempting to launch a nil function as a goroutine
func recover4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var f func()
	go f()
}

// 5. All goroutines are asleep - deadlock
// 6. Thread limit exhaustion

func PrintStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
