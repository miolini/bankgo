all: depends test build

build:
	go build -o bankgo_http ./httpapi/server
	go build -o bankgo_rpc ./rpc/server

docker_clean:
	docker rmi -f bankgo_http
	docker rmi -f bankgo_rpc

docker_build:
	docker build -t bankgo_http ./httpapi/server
	docker build -t bankgo_rpc ./rpc/server

docker: docker_build
	docker-compose up

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
