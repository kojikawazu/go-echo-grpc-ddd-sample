# Go のバージョン
GO_VERSION := 1.20
APP_NAME := go-echo-grpc-ddd-sample
CMD_PATH := cmd/server/main.go
BUILD_PATH := ./bin/$(APP_NAME)

# 依存関係の取得
.PHONY: init
init:
	@echo "Initializing project..."
	go mod tidy

.PHONY: generate
generate:
	@echo "Generating gRPC code..."
	protoc --go_out=./proto --go_opt=paths=import \
	       --go-grpc_out=./proto --go-grpc_opt=paths=import \
	       internal/interfaces/auth/auth.proto
	protoc --go_out=./proto --go_opt=paths=import \
	       --go-grpc_out=./proto --go-grpc_opt=paths=import \
	       internal/interfaces/todo/todo.proto
	protoc --go_out=./proto --go_opt=paths=import \
	       --go-grpc_out=./proto --go-grpc_opt=paths=import \
	       internal/interfaces/user/user.proto
# ビルド
.PHONY: build
build: 
	@echo "Building the application..."
	mkdir -p bin
	go build -o $(BUILD_PATH) $(CMD_PATH)

# アプリの実行
.PHONY: run
run:
	@echo "Running the application..."
	go run $(CMD_PATH)

# テストの実行
.PHONY: test
test:
	@echo "Running tests..."
	@TEST_MODE=true go test ./internal/test/... -v

# Linter チェック (golangci-lint を使用)
.PHONY: lint
lint:
	@echo "Running linter..."
	golangci-lint run ./...
