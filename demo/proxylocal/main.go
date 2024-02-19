package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type proxy struct {
	toAddr string
}

func (p *proxy) ServeHTTP(wr http.ResponseWriter, req *http.Request) {

	requrl := fmt.Sprintf("%s%s", p.toAddr, req.URL.String())

	log.Printf("%s, %s => %s", req.Method, req.URL.String(), requrl)

	newReq, _ := http.NewRequest(req.Method, requrl, req.Body)
	newReq.Header = req.Header
	resp, err := http.DefaultClient.Do(newReq)
	if err != nil {
		http.Error(wr, "Server Error", http.StatusInternalServerError)
		log.Fatal("ServeHTTP:", err)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	for k, vv := range resp.Header {
		for _, v := range vv {
			wr.Header().Add(k, v)
		}
	}
	wr.WriteHeader(resp.StatusCode)

	wr.Write(data)
}

func main() {
	var listenAddr string
	px := &proxy{}

	flag.StringVar(&listenAddr, "l", "0.0.0.0:8800", "listen")
	flag.StringVar(&px.toAddr, "to", "http://192.168.50.1:80", "to addr")
	flag.Parse()

	log.Printf("start: listen: %s, to http server: %s \n", listenAddr, px.toAddr)

	if err := http.ListenAndServe(listenAddr, px); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
