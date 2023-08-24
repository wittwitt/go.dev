package recover1

import (
	"fmt"
	"testing"
)

// 有defer，，即便有panic，，也会执行

func Test_defer_panic(t *testing.T) {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印1") }()
	defer func() { fmt.Println("打印2") }()
	defer func() { fmt.Println("打印3") }()

	var std *student
	fmt.Println(std.Name) // 触发panic

	fmt.Println("ok")
}

type student struct {
	Name string
	Age  int
}
