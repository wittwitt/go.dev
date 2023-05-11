// reverse_proxy.go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// 需要直接写一个 tcp server , 具体看 net.Listen 函数的官方实例
var addr = "http://127.0.0.1:2003"

// func NewSingleHostReverseProxy(lb algorithm.LoadBalance) *httputil.ReverseProxy {
func NewSingleHostReverseProxy() *httputil.ReverseProxy {

	director := func(r *http.Request) {
		// 对addr 进行解析
		rurl, err := url.Parse(addr)
		if err != nil {
			return
		}
		tr := rurl.RawQuery
		r.URL.Host = rurl.Host
		r.URL.Scheme = rurl.Scheme

		r.URL.Path = singleJoiningSlash(rurl.Path, r.URL.Path)
		// r.url = 127.0.0.1:8081/?id=123
		// tr = 127.0.0.1:8082/?coms=23
		if tr == "" || r.URL.RawQuery != "" {
			r.URL.RawQuery = tr + r.URL.RawQuery
		} else {
			r.URL.RawQuery = tr + "&" + r.URL.RawQuery
		}

		if _, ok := r.Header["User-Agent"]; !ok {
			r.Header.Set("User-Agent", "")
		}
	}

	modifyresponse := func(response *http.Response) error {
		if response.StatusCode != 200 {
			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			by := []byte("coms is good" + string(b))
			response.Body = ioutil.NopCloser(bytes.NewBuffer(by))
			response.ContentLength = int64(len(by))
			response.Header.Set("Content-Length", fmt.Sprintf("%v", response.ContentLength))
		}
		return nil
	}

	errhandler := func(w http.ResponseWriter, re *http.Request, err error) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	return &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyresponse,
		ErrorHandler:   errhandler}
}

// url 匹配模块
func singleJoiningSlash(dist, src string) string {
	// 获取 dist 的后缀 , 比如 ： 【127.0.0.1/】就是匹配 “/” 是否存在
	s := strings.HasSuffix(dist, "/")
	// 获取 src 的前缀 ，比如：【/coms】匹配 “/” 是否存在
	p := strings.HasPrefix(src, "/")

	switch {
	case s && p:
		return dist + src[1:]
	case !s && !p:
		return dist + "/" + src
	}

	return dist + src
}

//
func get() {

}
