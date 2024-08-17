package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ch := make(chan string)
	go NewServer(ch).ListenAndServe()
	go NewMockServer(ch).ListenAndServe()

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case msg := <-ch:
				fmt.Printf("[server] %s\n", msg)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				Host:   "localhost:8081",
			}),
		},
	}

	req, err := client.Get("http://localhost:8080/request")
	if err != nil {
		fmt.Println(err)
		done()
		wg.Wait()
		return
	}
	defer req.Body.Close()

	msg, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		done()
		wg.Wait()
		return
	}

	fmt.Printf("Response status: %s\n", req.Status)
	fmt.Printf("Body content: %s\n", msg)

	<-ctx.Done()
	wg.Wait()
}

func NewServer(ch chan<- string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		ch <- "real server"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a real server"))
	})
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func NewMockServer(ch chan<- string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		ch <- "mock server"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a mock server"))
	})
	return &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
}
