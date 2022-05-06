
folder: ## create folder structure
	@mkdir -p cmd
	@mkdir -p cmd/server
	@mkdir -p cmd/client
	@mkdir -p config
	@mkdir -p docs
	@mkdir -p external
	@mkdir -p internal
	@mkdir -p internal/adapter
	@mkdir -p internal/adapter/model
	@mkdir -p internal/api
	@mkdir -p internal/repository
	@mkdir -p internal/repository/model
	@mkdir -p internal/dto
	@mkdir -p internal/helper
	@mkdir -p internal/registry
	@mkdir -p internal/usecase
	@mkdir -p internal/util
	@mkdir -p pb
	@mkdir -p proto
	@mkdir -p vendor

dep: ## Install required packages
	@GO111MODULE=on go get -v github.com/golangci/golangci-lint/cmd/golangci-lint
	@@GO111MODULE=off go get -v github.com/rakyll/statik
	@@GO111MODULE=off go get github.com/golang/protobuf/protoc-gen-go \
		github.com/gogo/protobuf/protoc-gen-gogo \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/favadi/protoc-go-inject-tag \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

mod: ## Install go library
	@go mod tidy
	@go mod vendor

generate: ## Generate proto
	@third_party/protoc/bin/protoc \
		-I proto/ \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I $(GOPATH)/src/github.com/gogo/protobuf/ \
		-I $(GOPATH)/src/github.com/mwitkow/go-proto-validators/ \
		--gogo_out=plugins=grpc:pb \
		--govalidators_out=gogoimport=true:pb \
		--grpc-gateway_out=pb \
		--swagger_out=docs \
		proto/*.proto


run:
	@statik -f -src=docs -dest=cmd/server && cd cmd/server && go run main.go

build:
	@statik -f -src=docs -dest=cmd/server && cd cmd/server && go build -mod=mod -o bim .

test: ## Run test for whole project
	@go test -v $(TEST_PACKAGE) -cover -coverprofile=coverage.out

lint: ## Run linter
	@golangci-lint run ./...

install-oracle-instantclient: ## Install Oracle Instant Client
	@bash third_party/oracle/oracle-instantclient.sh

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
