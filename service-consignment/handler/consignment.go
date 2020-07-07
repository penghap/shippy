package handler

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"

	log "github.com/micro/go-micro/v2/logger"

	pb "github.com/penghap/shippy/service-consignment/proto/consignment"
	"github.com/penghap/shippy/service-consignment/repository"
	vesselProto "github.com/penghap/shippy/service-vessel/proto/vessel"
)

type Service struct {
	Session      *mgo.Session
	VesselClient vesselProto.VesselService
}

func (s *Service) GetRepo() repository.Repository {
	return &repository.ConsignmentRepository{s.Session.Clone()}
}

func (s *Service) CreateConsignment(ctx context.Context, in *pb.Consignment, out *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.VesselClient.FindAvailable(ctx, &vesselProto.Specification{
		Capacity:  int32(len(in.Containers)),
		MaxWeight: in.Weight,
	})

	if err != nil {
		return err
	}

	log.Info("Found vessel: %s \n", vesselResponse.Vessel)

	//vessel service
	in.VesselId = vesselResponse.Vessel.Id

	if err = repo.Create(in); err != nil {
		return err
	}

	out.Created = true
	out.Consignment = in
	return nil
}

func (s *Service) GetConsignments(ctx context.Context, in *pb.Consignment, out *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	out.Consignments = consignments
	return nil
}
