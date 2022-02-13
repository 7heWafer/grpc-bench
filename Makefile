protoc:
	rm -f api/*.pb.go
	PATH=$(PATH):$(go env GOPATH)/bin
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/api.proto
.PHONY: protoc

server:
	go run server/server.go 
.PHONY: server

benchmark:
	go test -benchmem -bench . example.com/example/client
.PHONY: benchmark
