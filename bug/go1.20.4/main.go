package main

import "fmt"

func foo() {
	println("foo")
}

var list = []func(){
	func() {
		println("1")
	},
	func() {
		println("2")
	},
	func() {
		println("3")
	},
}

var list2 = []string{
	"s1", "s2", "s3",
}

func main() {
	var fn func() = foo
	for _, fn = range list {
		fn()
	}
	// go.1.20.4之前，打印 foo foo foo 

	var ss string = "s0"
	for _, ss = range list2 {
		fmt.Println(ss)
	}
}
