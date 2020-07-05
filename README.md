# proto 

```shell script
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro/v2
```

# compile the proto file *.proto
```shell script
protoc --proto_path=. --go_out=. --micro_out=. *.proto
```