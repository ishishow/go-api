protoc:
	cd idl \
	&& protoc -I/usr/local/include -I. -I/usr/local/opt/protobuf/include -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --plugin=protoc-gen-grpc-gateway=$(GOPATH)/bin/protoc-gen-grpc-gateway --grpc-gateway_out=logtostderr=true:../pkg/interfaces/rpc ./*.proto \
	&& protoc -I/usr/local/include -I. -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis -I/usr/local/opt/protobuf/include --go_out=../pkg/interfaces/rpc --go-grpc_out=../pkg/interfaces/rpc ./*.proto

wire:
	wire ./pkg/interfaces/registry/

serve:
	go run ./cmd/api/main.go
gateway:
	go run ./cmd/gateway/main.go

exec_all:
	curl -X POST "http://localhost:8080/user/create" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"string\"}" \
	&& curl -X GET "http://localhost:8080/user/get" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" \
	&& curl -X POST "http://localhost:8080/game/finish" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" -H  "Content-Type: application/json" -d "{\"score\":10000}" \
	&& curl -X POST "http://localhost:8080/gacha/draw" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" -H  "Content-Type: application/json" -d "{\"times\":4}" \
	&& curl -X GET "http://localhost:8080/ranking/list?start=2" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" \
	&& curl -X GET "http://localhost:8080/collection/list" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" \

exec_use_envoy:
	curl -X POST "http://localhost:51051/user/create" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"string\"}" \
	&& curl -X GET "http://localhost:51051/user/get" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" \
	&& curl -X POST "http://localhost:51051/game/finish" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" -H  "Content-Type: application/json" -d "{\"score\":10000}" \
	&& curl -X POST "http://localhost:51051/gacha/draw" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" -H  "Content-Type: application/json" -d "{\"times\":4}" \
	&& curl -X GET "http://localhost:51051/ranking/list?start=2" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" \
	&& curl -X GET "http://localhost:51051/collection/list" -H  "accept: application/json" -H  "x-token: 85cc624c-c721-4e4f-9ae6-013df3ebae77" \


envoy_proto:
	cd idl \
	&& protoc -I/usr/local/include -I. -I/usr/local/opt/protobuf/include -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	--include_imports \
	--include_source_info \
	--descriptor_set_out=../envoy/proto.pb \
  	./*.proto