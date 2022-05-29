package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type ServerInfo struct {
	Addr string
	port int
	mux  *http.ServeMux
}

var servers []ServerInfo

func registerServer() {

	servers = make([]ServerInfo, 0)

	servers = append(servers, ServerInfo{Addr: "127.0.0.1", port: 8081, mux: http.NewServeMux()})
	servers = append(servers, ServerInfo{Addr: "127.0.0.1", port: 8082, mux: http.NewServeMux()})
	servers = append(servers, ServerInfo{Addr: "127.0.0.1", port: 8083, mux: http.NewServeMux()})

}

func (s *ServerInfo) Start(ctx context.Context) error {

	fmt.Printf("Start Addr:%s, port: %d\n", s.Addr, s.port)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.Addr, s.port), s.mux)
}

func (s *ServerInfo) Stop(ctx context.Context) error {
	fmt.Printf("Stop Addr:%s, port: %d\n", s.Addr, s.port)
	return fmt.Errorf("sss")
}

func main() {
	registerServer()
	eg, ctx := errgroup.WithContext(context.Background())
	wg := sync.WaitGroup{}
	for _, srv := range servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done()
			stopCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
			defer cancel()
			return srv.Stop(stopCtx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(context.WithValue(ctx, "ServerName", "aaa"))
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				fmt.Println("stop app")
				return fmt.Errorf("stop app")
			}
		}
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("get error: %v", err)
	}
}
