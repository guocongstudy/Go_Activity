package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var ctx context.Context
var cancle context.CancelFunc

func init() {
	ctx, cancle = context.WithCancel(context.Background())
}
func listenSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case sig := <-ch:
			fmt.Printf("recover signal %d\n", sig)
			cancle()
			return
		}
	}
}

func listenHttp(port int) {
	server := &http.Server{
		Addr: ":" + strconv.Itoa(port), Handler: nil,
	}
	go func() {
		select {
		case <-ctx.Done():
			server.Close()
			return
		}
	}()
	//
	server.ListenAndServe()
}
