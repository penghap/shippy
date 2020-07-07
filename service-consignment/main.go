// service-consignment/main.go
package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"

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

	// Create service
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name(srvName),
	)

	vesselClient := vesselProto.NewVesselService(vesselServiceName, srv.Client())

	// New Service
	service := micro.NewService(
		micro.Name(srvName),
		micro.Version(srvVersion),
	)

	// Initialise service
	service.Init()

	// Register Handler
	h := &handler.Service{session, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("Consignment service run failed: %v", err)
	}
}
