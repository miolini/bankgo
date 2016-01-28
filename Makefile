all: depends doc

doc:
	aglio -i httpapi.md -o api.html

depends:
	npm install -g aglio

proto:
  protoc --go_out=plugins=grpc:rpc/proto rpc/proto/balancestorage.proto