package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/prestonTao/upnp"
)

func main() {

	mapping := new(upnp.Upnp)

	err := mapping.AddPortMapping(55789, 24455, "TCP")
	if err == nil {
		fmt.Println("success !")
		// remove port mapping in gatway
		//mapping.Reclaim()
	} else {
		fmt.Println("fail !", err)
	}

	go run()

	ch := make(chan int)
	<-ch
}

func run() {
	fmt.Println("abc")
	service := "0.0.0.0:55789"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		fmt.Println("ccccccc")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn)
		go func(cnn net.Conn) {
			defer cnn.Close()

			daytime := time.Now().String()
			cnn.Write([]byte(daytime)) // don't care about return value

			var reader = bufio.NewReader(cnn)

			for {
				dat, _, err := reader.ReadLine()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(dat) + "aaaaaaaa")
			}

		}(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
