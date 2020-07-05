package main

import (
	"log"

	"github.com/penghap/shippy/service-user/handler"

	"github.com/micro/go-micro/v2"

	"github.com/penghap/shippy/service-user/repository"

	pb "github.com/penghap/shippy/service-user/proto/user"

	"github.com/penghap/shippy/service-user/database"
)

const (
	srvName = "go.micro.srv.user"
)

func main() {
	db, err := database.CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := repository.UserRepository{db}

	srv := micro.NewService(
		micro.Name(srvName),
	)

	srv.Init()

	h := &handler.Service{}
	h.SetRepo(&repo)

	pb.RegisterUserServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
