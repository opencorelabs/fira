.DEFAULT_GOAL := bin/server

bin/server: gen
	@echo "Building server..."
	@mkdir -p bin
	@go build -o bin/server ./cmd/server

.PHONY: gen
gen: protoreqs
	@echo "Generating..."
	@go generate ./...
	@buf generate
	@buf build
	@go mod tidy

.PHONY: reqs
reqs: protoreqs clientreqs
	@echo "Installing dependencies..."
	@go mod download

.PHONY: protoreqs
protoreqs:
	@buf --version >/dev/null 2>&1 || go install github.com/bufbuild/buf/cmd/buf@latest
	@which protoc-gen-go >/dev/null 2>&1 || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc >/dev/null 2>&1 || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@which protoc-gen-grpc-gateway >/dev/null 2>&1 || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@which protoc-gen-openapiv2 >/dev/null 2>&1 || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@which protoc-gen-buf-breaking >/dev/null 2>&1 || go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest
	@which protoc-gen-buf-lint >/dev/null 2>&1 || go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest

.PHONY: clientreqs
clientreqs:
	@which yarn >/dev/null 2>&1 || npm install -g yarn
	@cd client && yarn install

.PHONY: dev
dev:
	@echo "Starting dev server in docker..."
	@docker-compose up -d --build --remove-orphans dev
	@docker-compose logs -f dev
