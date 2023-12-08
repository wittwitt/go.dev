package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/rpc/v0", handleRequest)
	fmt.Println("Proxy server listening on :17080")
	http.ListenAndServe(":17080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Print the incoming request details
	fmt.Printf("Received request: %s %s\n", r.Method, r.URL)

	rAddr := "http://172.31.20.11:17080"
	//rAddr := "http://127.0.0.1:17081"

	targetURL, err := url.Parse(rAddr)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	r.Host = "172.31.20.11:17080"

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Set("Content-Type", "application/json")
		return nil
	}

	proxy.ServeHTTP(w, r)
}
