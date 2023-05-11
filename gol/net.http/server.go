package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type clientResponse struct {
	ID string
}

func Indexhandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Println(url.ParseQuery(r.URL.RequestURI()))

	// fmt.Fprintln(w, "Tmh3je9zbnHAvPfwwHhQsFSJmKkeRTtKqmV", r.URL.RequestURI())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("ccccc", "Keep-Alive")

	ddd, _ := json.Marshal(&clientResponse{ID: "adddd"})

	_, err := w.Write(ddd)
	if err != nil {
		fmt.Println("err")
	}

}

func server() {

	addr := "0.0.0.0:6060"

	ls, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("could not listen: %w", err)
		return
	}

	srv := &http.Server{
		Handler: http.DefaultServeMux,
		BaseContext: func(listener net.Listener) context.Context {
			// ctx, _ := tag.New(context.Background(), tag.Upsert(metrics.APIInterface, id))
			// return ctx
			return context.TODO()
		},
		ReadTimeout:       90 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	// srv.SetKeepAlivesEnabled(false)

	http.HandleFunc("/", Indexhandler)

	err = srv.Serve(ls)
	if err != http.ErrServerClosed {
		log.Printf("rpc server failed: %s", err)
	}
}
