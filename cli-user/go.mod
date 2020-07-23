module github.com/penghap/shippy/cli-user

go 1.14

replace (
	google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
	github.com/penghap/shippy/service-user => ../service-user
)

require (
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/penghap/shippy/service-user v0.0.0-20200705150531-4a357099c0a8
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
)
