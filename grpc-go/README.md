# README

Go×gRPCでのデモプログラムです。

## Generate gRPC client

```bash
protoc --go_out=./pkg/ --go_opt=paths=source_relative \
  --go-grpc_out=./pkg/ --go-grpc_opt=paths=source_relative \
  proto/service.proto
```

## How to send request from console

Use `grpcurl`

```bash
grpcurl -plaintext localhost:8080 list

grpcurl -plaintext localhost:8080 list service.GreetingService

grpcurl -plaintext -d '{"name":"aqyuki"}' localhost:8080 service.GreetingService.Hello
```
