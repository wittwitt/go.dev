package main

import (
	"fmt"

	"github.com/wittwitt/go.dev/demo/wire1/node"
)

func main() {
	e, err := node.IniPool()
	if err != nil {
		fmt.Println(err)
		return
	}
	e.Start()
	ch := make(chan struct{})
	<-ch
}
