package main

import (
	"crypto/tls"
	"flag"
	"log"
	"strings"
	"time"
)

var checkD time.Duration
var addrs []string
var dialD time.Duration

func main() {
	var checkDStr, addrsStr, dailDstr string
	flag.StringVar(&checkDStr, "d", "10s", "-d check interval, h,m,s")
	flag.StringVar(&addrsStr, "addrs", "www.baidu.com:80", "-addrs ip:port;ip2:port2 ")
	flag.StringVar(&dailDstr, "t", "10s", "-t dial timeout h, m, s")
	flag.Parse()

	{
		var err error
		checkD, err = time.ParseDuration(checkDStr)
		if err != nil {
			log.Printf("checkD: %v", err)
			return
		}

		dialD, err = time.ParseDuration(dailDstr)
		if err != nil {
			log.Printf("dialD: %v", err)
			return
		}

		addrStrParts := strings.Split(addrsStr, ";")
		for _, addr := range addrStrParts {
			err = check(addr, dialD)
			if err != nil {
				log.Printf("check: %v, %s", err, addr)
				return
			}
			addrs = append(addrs, addr)
		}
	}
	log.Printf("check start, interval: %s", checkDStr)
	run()
}

func run() {
	timer := time.NewTimer(checkD)
	for range timer.C {
		for _, addr := range addrs {
			err := check(addr, dialD)
			if err != nil {
				log.Printf("check: %s, %v", addr, err)
			} else {
				log.Printf("check: %s, ok", addr)
			}
		}
		timer.Reset(checkD)
	}
}

func check(addr string, timeoutD time.Duration) error {
	defer func() {
		if rev := recover(); rev != nil {
			log.Printf("check %s err: %v", addr, rev)
		}
	}()

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	log.Println("start")
	conn, err := tls.Dial("tcp", "n.xmeer.com:16443", conf)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	conn.Write([]byte(`{"id":null,"method":"mining.subscribe","params":["1111","x.x"]}`))
	conn.Write([]byte("\n"))

	// cnn, err := net.DialTimeout("tcp", "n.xmeer.com:15443", timeoutD)
	// if err != nil {
	// 	return err
	// }
	// cnn.Close()
	return nil
}
