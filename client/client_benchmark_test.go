package client

import (
	"context"
	"log"
	"testing"

	"example.com/example/config"
)

func BenchmarkPingHTTP(b *testing.B) {
	client := newBenchmarkClient(config.Addr, config.PortHTTP, config.PortGRPC)
	defer client.Done()
	for i := 0; i < b.N; i++ {
		if _, err := client.PingHTTP(); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkPingGRPC(b *testing.B) {
	client := newBenchmarkClient(config.Addr, config.PortHTTP, config.PortGRPC)
	defer client.Done()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		if _, err := client.PingGRPC(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
