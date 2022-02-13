package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"example.com/example/api"
	"example.com/example/config"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type grpcserver struct {
	api.UnimplementedPingServer
}

func (g grpcserver) ResolvePing(ctx context.Context, r *api.PingRequest) (*api.PingReply, error) {
	return &api.PingReply{Msg: "pong"}, nil
}

func grpcServer(wg *sync.WaitGroup, addr string, port int) {
	defer wg.Done()
	s := grpc.NewServer()
	addrs := fmt.Sprintf("%s:%d", addr, port)
	api.RegisterPingServer(s, &grpcserver{})
	l, err := net.Listen("tcp", addrs)
	if err != nil {
		panic(err)
	}
	log.Println("starting grpc server", addrs)
	log.Println(s.Serve(l))
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("msg")
	_, _ = w.Write([]byte("pong"))
}

func httpServer(wg *sync.WaitGroup, addr string, port int) {
	defer wg.Done()
	r := mux.NewRouter()
	r.HandleFunc("/", handleHTTP).Methods(http.MethodGet)
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", addr, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	log.Println("starting http server", s.Addr)
	log.Println(s.ListenAndServe())
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go httpServer(wg, config.Addr, config.PortHTTP)
	go grpcServer(wg, config.Addr, config.PortGRPC)
	wg.Wait()
}
