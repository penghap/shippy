package handler

import (
	log "github.com/micro/go-micro/v2/logger"
	pb "github.com/penghap/shippy/service-vessel/proto/vessel"
	"github.com/penghap/shippy/service-vessel/repository"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type Service struct {
	Session *mgo.Session
}

func (s *Service) GetRepo() repository.Repository {
	return &repository.VesselRepository{s.Session.Clone()}
}

func (s *Service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	log.Info("FindAvailable in", in)
	var vessel *pb.Vessel
	vessel, err := repo.FindAvailable()

	if err != nil {
		return err
	}

	out.Vessel = vessel
	return nil
}

func (s *Service) Create(ctx context.Context, in *pb.Vessel, out *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(in); err != nil {
		return err
	}

	out.Vessel = in
	out.Created = true
	return nil
}
