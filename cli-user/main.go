package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"

	pb "github.com/penghap/shippy/service-user/proto/user"
)

const (
	userServiceName = "go.micro.srv.user"
)

func main() {
	// Define our flags
	srv := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			&cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	// Start as service
	option := func(c *cli.Context) error {
		name := c.String("name")
		email := c.String("email")
		password := c.String("password")
		company := c.String("company")

		// call service
		client := pb.NewUserService(userServiceName, srv.Client())
		ctx := context.Background()

		user := &pb.User{
			Name:     name,
			Email:    email,
			Company:  company,
			Password: password,
		}

		log.Println("User:", user)
		rsp, err := client.Create(ctx, user)
		if err != nil {
			log.Println("error creating user: ", err.Error())
			return err
		}

		// print the response
		log.Println("Response: ", rsp.User)
		return nil
	}
	srv.Init(
		micro.Action(option),
	)
}
