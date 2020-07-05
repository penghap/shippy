// service-vessel/main.go
package main

import (
	"errors"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	"golang.org/x/net/context"

	"github.com/penghap/shippy/service-vessel/database"
	"github.com/penghap/shippy/service-vessel/handler"
	pb "github.com/penghap/shippy/service-vessel/proto/vessel"
)

const (
	dbHost  = "localhost:27017"
	srvName = "go.micro.srv.vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(in *pb.Specification) (*pb.Vessel, error) {

	for _, vessel := range repo.vessels {
		if in.Capacity <= vessel.Capacity && in.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that input")
}

type Service struct {
	repo Repository
}

func (s *Service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	vessel, err := s.repo.FindAvailable(in)
	if err != nil {
		return err
	}

	out.Vessel = vessel
	return nil
}

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

	srv := micro.NewService(
		micro.Name(srvName),
	)

	srv.Init()

	h := &handler.Service{session}

	pb.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("Vessel service run failed: %v", err)
	}
}
