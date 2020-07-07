package main

import (
	"os"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"gopkg.in/mgo.v2"

	"github.com/penghap/shippy/service-vessel/database"
	"github.com/penghap/shippy/service-vessel/handler"
	pb "github.com/penghap/shippy/service-vessel/proto/vessel"
)

const (
	dbHost  = "localhost:27017"
	srvName = "shippy.service.vessel"
	//srvTopic   = "shippy.service.vessel"
	srvVersion = "latest"
)

func connectDB() *mgo.Session {
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

	return session
}

func main() {
	// connect DB
	session := connectDB()

	// New Service
	service := micro.NewService(
		micro.Name(srvName),
		micro.Version(srvVersion),
	)

	// Initialise service
	service.Init()

	// Register Handler
	h := &handler.Service{session}
	pb.RegisterVesselServiceHandler(service.Server(), h)

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("shippy.service.service-vessel", service.Server(), new(subscriber.ServiceVessel))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
