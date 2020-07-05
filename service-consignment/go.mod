module github.com/penghap/shippy/service-consignment

go 1.14

replace (
	github.com/penghap/shippy/service-vessel => ../service-vessel
	google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/micro/protoc-gen-micro/v2 v2.3.0 // indirect
	github.com/penghap/shippy/service-vessel v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/protobuf v1.25.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
