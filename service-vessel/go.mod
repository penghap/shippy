module github.com/penghap/shippy/service-vessel

go 1.14

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/genproto v0.0.0-20200702021140-07506425bd67 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
