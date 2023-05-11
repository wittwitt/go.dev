package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

var bench_N = 500

func mutiplThread() {
	var wg sync.WaitGroup
	for i := 0; i < bench_N; i++ {
		wg.Add(1)
		go func() {
			res, err := httpClient(context.TODO(), "http://192.168.56.101:6060/") // "http://127.0.0.1:6060/")
			if err != nil {
				fmt.Println(err)

			} else {
				fmt.Println(res.ID)
			}
			wg.Done()

			// time.Sleep(100 * time.Millisecond)
		}()
	}
	wg.Wait()
}

var _defaultHTTPClient = &http.Client{
	Transport: &http.Transport{
		// ！！！！注意，这个参数影响了 Keep-Alive，，
		// 1. tcp连接不会复用
		// 2. http header 永远是 Connection：close
		// 3. hreq.Header.Set("Connection", "Keep-Alive") // 这个没用

		// Proxy: http.ProxyFromEnvironment,

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		// ForceAttemptHTTP2:     true,
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		// TLSHandshakeTimeout: 10 * time.Second,
		// ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives: false,
	},
	Timeout: 30 * time.Second,
}

var requestHeader http.Header = http.Header{}

func httpClient(ctx context.Context, addr string) (clientResponse, error) {
	// b, err := json.Marshal(&cr.req)
	b, err := json.Marshal(&clientResponse{ID: "abc"})
	if err != nil {
		return clientResponse{}, fmt.Errorf("mershaling requset: %w", err)
	}

	hreq, err := http.NewRequest("GET", addr, bytes.NewReader(b))
	if err != nil {
		return clientResponse{}, fmt.Errorf("req: %v", err) //  err
	}

	hreq.Close = false

	hreq.Header = requestHeader.Clone()
	if ctx != nil {
		hreq = hreq.WithContext(ctx)
	}

	hreq.Header.Set("User-Agent", "firefox v10000")
	hreq.Header.Add("Content-Type", "application/json")
	hreq.Header.Add("hah", "thishis req")

	// if config.transport != nil {
	// 	_defaultHTTPClient.Transport = config.transport
	// }
	httpResp, err := _defaultHTTPClient.Do(hreq)
	if err != nil {
		return clientResponse{}, fmt.Errorf("do: %v", err) // err
	}

	// ！！！！http 及时关闭 ，防止大量wait-time
	// 1. blody一定要关闭,,或者读取完
	defer httpResp.Body.Close()
	// defer func() {
	// 	io.Copy(ioutil.Discard, httpResp.Body) // 读取
	// 	httpResp.Body.Close()
	// }()

	var resp clientResponse

	ff := json.NewDecoder(httpResp.Body)
	if err := ff.Decode(&resp); err != nil {
		return clientResponse{}, fmt.Errorf("http status %s unmarshaling response: %w", httpResp.Status, err)
	}

	return resp, nil
}
