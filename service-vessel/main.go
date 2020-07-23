package main

import (
	"os"
	"time"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/penghap/shippy/service-vessel/database"
	"github.com/penghap/shippy/service-vessel/handler"
	pb "github.com/penghap/shippy/service-vessel/proto/vessel"
)

const (
	dbHost     = "localhost:27017"
	srvName    = "shippy.service.vessel"
	srvTopic   = "shippy.service.vessel"
	srvVersion = "latest"
)

func main() {
	// connect DB
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

	//// register
	//registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = []string{"http://127.0.0.1:2379"}
	//})

	// New Service
	service := micro.NewService(
		micro.Name(srvName),
		micro.Version(srvVersion),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		//print log after start
		micro.AfterStart(func() error {
			log.Infof("🚀 service listen on ...")
			return nil
		}),
		//micro.Registry(registerDrive),
	)

	// Initialise service
	service.Init()

	// Register Handler
	h := &handler.Service{session}
	pb.RegisterVesselServiceHandler(service.Server(), h)

	// Register Struct as Subscriber
	//micro.RegisterSubscriber(srvTopic, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
