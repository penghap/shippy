// service-consignment/main.go
package main

import (
	"context"
	"os"
	"time"

	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"

	"github.com/penghap/shippy/service-consignment/database"
	"github.com/penghap/shippy/service-consignment/handler"
	pb "github.com/penghap/shippy/service-consignment/proto/consignment"
	vesselProto "github.com/penghap/shippy/service-vessel/proto/vessel"
)

const (
	srvName           = "shippy.service.consignment"
	srvVersion        = "latest"
	vesselServiceName = "shippy.service.vessel"
	dbHost            = "localhost:27017"
)

func main() {
	// Database host from env
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = dbHost
	}
	session, err := database.CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("Could not connect to mongo host %s - %v", host, err)
	}

	// register
	registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://localhost:2379"}
	})

	// Create service
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name(srvName),
		micro.Version(srvVersion),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		//print log after start
		micro.AfterStart(func() error {
			log.Infof("🚀 service listen on ...")
			return nil
		}),
		micro.Registry(registerDrive),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)

	// Initialise service
	srv.Init()

	hystrixGo.DefaultMaxConcurrent = 3 //change concurrrent to 3
	hystrixGo.DefaultTimeout = 200     //change timeout to 200 milliseconds

	vesselClient := vesselProto.NewVesselService(vesselServiceName, srv.Client())

	// Register Handler
	h := &handler.Service{session, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// create publisher
	event := micro.NewEvent(vesselServiceName, srv.Client())

	// publish message every second
	for now := range time.Tick(time.Second) {
		if err := event.Publish(context.Background(), &vesselProto.Response{Created: true}); err != nil {
			log.Info("publish err", err, now)
		}
		log.Info("now: ", now)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Consignment service run failed: %v", err)
	}
}
