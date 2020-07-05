protoc:
	protoc -I. --proto_path=. --micro_out=. --go_out=. proto/vessel/vessel.proto

build:
	protoc -I. --proto_path=. --micro_out=. --go_out=. proto/vessel/vessel.proto

	GOOS=linux GOARCH=amd64 go build
	docker build -t service-vessel .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns service-vessel
