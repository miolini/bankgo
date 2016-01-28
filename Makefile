all: depends test build

build:
	go build -o bankgo_http ./httpapi/server
	go build -o bankgo_rpc ./rpc/server

test:
	go test -v ./...

doc:
	aglio -i httpapi.md -o api.html

depends:
	go get -v ./...

doc_depends:
	npm install -g aglio

proto:
	protoc --go_out=plugins=grpc:rpc/proto rpc/proto/balancestorage.proto
