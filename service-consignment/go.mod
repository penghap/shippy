module github.com/penghap/shippy/service-consignment

go 1.14

replace (
	github.com/penghap/shippy/service-vessel => ../service-vessel
	// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
	// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
	google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.9.1
	github.com/penghap/shippy/service-vessel v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/protobuf v1.25.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
