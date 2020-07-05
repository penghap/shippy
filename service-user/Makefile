protoc:
	protoc -I. --proto_path=. --micro_out=. --go_out=. proto/user/user.proto

build:
	protoc -I. --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t service-user .

run:
	docker run -p 50053:50051 service-user