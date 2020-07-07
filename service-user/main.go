package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/penghap/shippy/service-user/database"
	"github.com/penghap/shippy/service-user/handler"
	pb "github.com/penghap/shippy/service-user/proto/user"
	"github.com/penghap/shippy/service-user/repository"
)

const (
	srvName = "shippy.service.user"
	//srvTopic   = "shippy.service.user"
	srvVersion = "latest"
)

func main() {
	db, err := database.CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := repository.UserRepository{db}

	// New Service
	service := micro.NewService(
		micro.Name(srvName),
		micro.Version(srvVersion),
	)

	// Initialise service
	service.Init()

	// Register Handler
	h := &handler.Service{}
	h.SetRepo(&repo)
	pb.RegisterUserServiceHandler(service.Server(), h)

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("", service.Server(), new(subscriber.ServiceUser))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
