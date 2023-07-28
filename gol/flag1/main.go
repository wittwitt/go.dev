package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)

	flag.String("car", "./car", "car path")
	flag.Parse()
	fmt.Println("cccc", flag.Args())

}
