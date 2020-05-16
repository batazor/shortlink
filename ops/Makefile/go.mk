# GO TASKS =============================================================================================================

generate: ## Code generation
	@echo "proto generation link entity"
	@protoc -I/usr/local/include -I. \
	--gotemplate_out=all=true,template_dir=pkg/api/graphql/template:pkg/api/graphql \
	--go_out=plugins=grpc:. \
	pkg/link/link.proto

	@protoc -I/usr/local/include -I. \
    	--gotemplate_out=all=true,template_dir=internal/store/query/template:internal/store/query \
    	pkg/link/link.proto

	@echo "proto generation gRPC-web"
	@protoc -I/usr/local/include -I. \
	-I=pkg/api/grpc-web \
	-I=third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
	--go_out=plugins=grpc:. \
	--swagger_out=logtostderr=true,allow_delete_body=true:. \
	--grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
	pkg/api/grpc-web/api.proto
	@mv pkg/api/grpc-web/api.swagger.json docs/api.swagger.json

	@echo "Generate from .go code"
	@go generate internal/store/postgres/postgres.go
	@go generate internal/di/wire.go

	@echo "Generate go static"
	@pkger -o cmd/shortlink

	@make fmt

.PHONY: fmt
fmt: ## Format source using gofmt
	@echo Apply go fmt
	@gofmt -l -s -w cmd pkg internal

gosec: ## Golang security checker
	@gosec ./...

golint: ## Linter for golang
	@golangci-lint run ./...

test: ## Run all test
	@echo "run test"
	@sh ./ops/scripts/coverage.sh

bench: ## Run benchmark tests
	go test -bench ./...