# grpc-bench

## Prerequisites
- [Go](https://go.dev/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [Make](https://www.gnu.org/software/make/)
    - or just run the commands found in the Makefile

```sh
# Terminal 1
make server
# Terminal 2
make benchmark
```

## Results
```
go test -benchmem -bench . example.com/example/client
goos: darwin
goarch: amd64
pkg: example.com/example/client
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkPingHTTP-12               16353             77944 ns/op            4221 B/op         54 allocs/op
BenchmarkPingGRPC-12               12873             93254 ns/op            4933 B/op         99 allocs/op
PASS
ok      example.com/example/client      4.314s
```
