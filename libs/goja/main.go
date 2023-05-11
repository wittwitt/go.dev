package main

import (
	"fmt"

	"github.com/dop251/goja"
)

// Goja is an implementation of ECMAScript 5.1 in pure Go with emphasis on standard compliance and performance.

func main() {
	ObjectJS()
	go run()
	select {}
}

func run() {
	vm := goja.New()

	console := vm.NewObject()
	console.Set("log", log)
	vm.Set("console", console)

	type Web3 struct {
		Name string
	}
	web3 := &Web3{}
	vm.Set("web3", web3)

	for {

		str := ""
		fmt.Scanln(&str)

		goaV, err := vm.RunScript("a", str)

		if err != nil {
			fmt.Println(goaV, err)
			continue
		}

		// c, err := vm.RunScript("a", `web3.Name="haha,js"`)
		// fmt.Println(c, err)
		// fmt.Println(web3.Name)
		// v, err := vm.RunString("2 + 2")
		// if err != nil {
		// 	panic(err)
		// }
		// if num := v.Export().(int64); num != 4 {
		// 	panic(num)
		// }
	}
}

// ObjectJS ObjectJS
func ObjectJS() {
	vm := goja.New()

	o := vm.NewObject()
	o.Set("Name", "")
	o.Set("Age", 1)

	vm.Set("zhangsan", o)

	zsV, err := vm.RunScript("zs", `zhangsan.Name="zs";zhangsan.Age=20;`)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(zsV)

	fmt.Println(o.Get("Name"))

}

// log console.log
func log(call goja.FunctionCall) goja.Value {
	str := call.Argument(0)
	fmt.Print(str.String())
	return str
}
