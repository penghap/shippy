module github.com/penghap/shippy/cli-user

go 1.14

replace (
	github.com/penghap/shippy/service-user => ../service-user
	google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
)

require (
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/penghap/shippy/service-user v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
)
