.PHONY: build run docs proto

build-metadata:
	go build -o bin/metadata/metadata cmd/metadata/main.go

build-data:
	go build -o bin/data/data cmd/data/main.go

build-sumo:
	go build -o bin/sumo-integration/sumo-integration cmd/sumo-integration/main.go

build-managment:
	go build -o bin/managment/managment cmd/managment/main.go

build-storage:
	go build -o bin/storage/storage cmd/storage/main.go

build: build-metadata build-data build-sumo build-managment build-storage

run-metadata:
	cd cmd/metadata && \
	go run main.go

run-data:
	cd cmd/data && \
	go run main.go

run-sumo:
	cd cmd/sumo-integration && \
	go run main.go

run-managment:
	cd cmd/managment && \
	go run main.go

run-storage:
	cd cmd/storage && \
	go run main.go

run: run-metadata run-data run-sumo run-managment run-storage

docs-metadata:
	swag init -g cmd/metadata/main.go -o internal/metadata/controllers/docs

docs-data:
	swag init -g cmd/data/main.go -o internal/data/controllers/docs

docs-sumo:
	swag init -g cmd/sumo-integration/main.go -o internal/sumo-integration/controllers/docs

docs-managment:
	swag init -g cmd/managment/main.go -o internal/managment/controllers/docs

docs-storage:
	swag init -g cmd/storage/main.go -o internal/storage/controllers/docs

docs: docs-metadata docs-data docs-sumo docs-managment docs-storage

PROTOC_OUT_DIR="./ts-lib"

protoc-metadata:
	@protoc -I internal/metadata/proto \
	-I $(shell go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	--openapiv2_out=cmd/metadata/assets/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=metadata \
	--grpc-gateway-ts_out=ts_import_roots=./proto,ts_import_root_aliases=base:$(PROTOC_OUT_DIR)/metadata \
	internal/metadata/proto/**/*.proto

protoc-data:
	@protoc -I internal/data/proto \
	-I $(shell go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	--openapiv2_out=cmd/data/assets/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=data \
	--grpc-gateway-ts_out=ts_import_roots=./proto,ts_import_root_aliases=base:$(PROTOC_OUT_DIR)/data \
	internal/data/proto/**/*.proto

protoc-sumo:
	@protoc -I internal/sumo-integration/proto \
	-I $(shell go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	--openapiv2_out=cmd/sumo-integration/assets/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=sumo-integration \
	--grpc-gateway-ts_out=ts_import_roots=./proto,ts_import_root_aliases=base:$(PROTOC_OUT_DIR)/sumo-integration \
	internal/sumo-integration/proto/**/*.proto

protoc-managment:
	@protoc -I internal/managment/proto \
	-I $(shell go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	--openapiv2_out=cmd/managment/assets/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=managment \
	--grpc-gateway-ts_out=ts_import_roots=./proto,ts_import_root_aliases=base:$(PROTOC_OUT_DIR)/managment \
	internal/managment/proto/**/*.proto

protoc-storage:
	@protoc -I internal/storage/proto \
	-I $(shell go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	--openapiv2_out=cmd/storage/assets/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=storage \
	--grpc-gateway-ts_out=ts_import_roots=./proto,ts_import_root_aliases=base:$(PROTOC_OUT_DIR)/storage \
	internal/storage/proto/**/*.proto

proto: protoc-metadata protoc-data protoc-sumo protoc-managment protoc-storage

release:
	goreleaser release --snapshot --clean

publish:
	goreleaser  --clean --skip=validate,announce