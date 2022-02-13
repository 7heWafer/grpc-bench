package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"example.com/example/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const requestTimeout time.Duration = 500 * time.Millisecond

type impl struct {
	httpPath   string
	httpClient *http.Client
	grpcPath   string
	grpcClient api.PingClient
	grpcConn   *grpc.ClientConn
}

type benchmarkClient interface {
	PingHTTP() (*api.PingReply, error)
	PingGRPC(ctx context.Context) (*api.PingReply, error)
	Done()
}

func newBenchmarkClient(addr string, httpPort int, grpcPort int) benchmarkClient {
	i := &impl{
		httpPath: fmt.Sprintf("http://%s:%d", addr, httpPort),
		grpcPath: fmt.Sprintf("%s:%d", addr, grpcPort),
	}
	i.httpClient = &http.Client{Timeout: requestTimeout}
	conn, err := grpc.Dial(i.grpcPath, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	i.grpcConn = conn
	i.grpcClient = api.NewPingClient(conn)
	return i
}

func (impl impl) PingHTTP() (*api.PingReply, error) {
	req, err := http.NewRequest(http.MethodGet, impl.httpPath+"/?msg=ping", nil)
	if err != nil {
		return nil, err
	}
	response, err := impl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &api.PingReply{Msg: string(body)}, err
}

func (impl impl) PingGRPC(ctx context.Context) (*api.PingReply, error) {
	return impl.grpcClient.ResolvePing(ctx, &api.PingRequest{Msg: "ping"})
}

func (impl impl) Done() {
	impl.grpcConn.Close()
}
