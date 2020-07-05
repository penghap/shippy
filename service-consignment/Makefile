protobuf:
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get -u github.com/micro/protoc-gen-micro/v2

protoc:
	protoc -I. --proto_path=. --micro_out=. --go_out=. proto/consignment/consignment.proto

build:
	protoc -I. --proto_path=. --micro_out=. --go_out=. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t service-consignment .

run:
	docker run -p 50051:50051 service-consignment