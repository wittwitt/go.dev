package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("go")

	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "client":
		mutiplThread()
	case "server":
		server()
	}
}
