module github.com/penghap/shippy/cli-consignment

go 1.14

replace (
	github.com/penghap/shippy/service-consignment => ../service-consignment
	github.com/penghap/shippy/service-vessel => ../service-vessel
	google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
)

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/penghap/shippy/service-consignment v0.0.0-00010101000000-000000000000
	golang.org/x/tools v0.0.0-20200702044944-0cc1aa72b347 // indirect
)
