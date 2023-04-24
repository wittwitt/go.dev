package err1

import (
	"fmt"
	"testing"
)

// 有defer，，即便有panic，，也会执行

func Test_f1(t *testing.T) {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印1") }()
	defer func() { fmt.Println("打印2") }()
	defer func() { fmt.Println("打印3") }()
	// panic("触发异常")

	var std *student
	fmt.Println(std.Name)
}

type student struct {
	Name string
	Age  int
}
