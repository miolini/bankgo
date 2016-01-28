all: docker

build: depends test
	go build -o bankgo_http ./httpapi/server
	go build -o bankgo_rpc ./rpc/server

docker_clean:
	docker rmi -f bankgo/http || true
	docker rmi -f bankgo/rpc || true

docker_build:
	docker build --no-cache=true --rm -t bankgo/http ./httpapi/server
	docker build --no-cache=true --rm -t bankgo/rpc ./rpc/server

docker: docker_build
	docker-compose up

test:
	go test -v -race ./...

doc:
	aglio -i httpapi.md -o api.html

depends:
	go get -v ./...

doc_depends:
	npm install -g aglio

proto:
	protoc --go_out=plugins=grpc:rpc/proto rpc/proto/balancestorage.proto
