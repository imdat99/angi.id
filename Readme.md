```bash
protoc --go_out=. protos/*.proto
protoc --go_out=. --go_opt=paths=source_relative protos/*.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/*.proto
```