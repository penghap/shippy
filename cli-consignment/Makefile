build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t cli-consignment .
run:
	docker run -e MICRO_REGISTRY=mdns cli-consignment