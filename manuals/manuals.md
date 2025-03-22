
## ディレクトリ構成

```
backend/
├── internal/
│   └── interfaces/
│       └── user/
│           └── user.proto
├── proto/
│   └── user.pb.go
```

## Homebrew でインストール

```bash
brew install protobuf
```

## goでインストール

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## .proto を protoc で Go に変換する

```bash
protoc \
  --go_out=./proto \
  --go-grpc_out=./proto \
  internal/interfaces/user/user.proto
```
