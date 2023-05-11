package http1

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

func Test_client(t *testing.T) {

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout: 2 * time.Second,
				// Deadline:  time.Now(), // time.Now().Add(23 * time.Second),
				KeepAlive: 20 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 2 * time.Second,
			MaxIdleConnsPerHost: 2,
			MaxConnsPerHost:     2,
			MaxIdleConns:        2,
		},
		Timeout: 5 * time.Second,
	}

	for {
		fmt.Println("====")
		time.Sleep(1 * time.Second)

		req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/p1", nil)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Do", err)
			continue
		}

		// 不读body内容就response.Body.Close()，那么连接会被主动关闭，得不到复用
		data, err := io.ReadAll(res.Body)
		fmt.Println(string(data), err)
		res.Body.Close()
	}
}
