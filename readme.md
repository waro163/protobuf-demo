# compile

```
protoc  --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./proto/common/kind.proto 
protoc  --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./proto/user/people.proto ./proto/user/user.proto

```