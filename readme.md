# compile

```
protoc  --go_out=paths=source_relative:. --go-grpc_out=. ./proto/common/kind.proto 
protoc  --go_out=paths=source_relative:. --go-grpc_out=. ./proto/user/people.proto ./proto/user/user.proto

```