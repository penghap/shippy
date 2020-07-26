package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
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

	// register
	registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://localhost:2379"}
	})

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
		micro.Registry(registerDrive),
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
