package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: false, // true,
	}
	conn, err := tls.Dial("tcp", "meer.meerpool.com:15443", conf)
	// conn, err := tls.Dial("tcp", "127.0.0.1:9001", conf)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	var str string
	for {

		_, err := fmt.Scanln(&str)
		if err != nil {
			fmt.Println(err)
			continue
		}

		n, err := conn.Write([]byte("client:" + str + "\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
		buf := make([]byte, 100)
		n, err = conn.Read(buf)
		if err != nil {
			log.Println(n, err)
			return
		}
		println(string(buf[:n]))

	}
}
