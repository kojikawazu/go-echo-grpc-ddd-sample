
## ディレクトリ構成

```
backend/
├── internal/
│   └── interfaces/
│       └── auth/
│       │   ├── auth_handler.go
│       │   └── auth.proto
│       └── todo/
│       │   ├── todo_handler.go
│       │   └── todo.proto
│       └── user/
│           ├── user_handler.go
│           └── user.proto
├── proto/github.com/grpc/backend/proto
│   ├── auth.pb.go
│   ├── todo.pb.go
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

- `Makefile` に設定済。`make generate`コマンドで実施可能。
