package repository

import (
	pb "github.com/penghap/shippy/service-vessel/proto/vessel"

	"gopkg.in/mgo.v2"
)

const (
	dbName         = "shippy"
	userCollection = "user"
)

type Repository interface {
	Create(*pb.Vessel) error
	FindAvailable() (*pb.Vessel, error)
	Close()
}

type VesselRepository struct {
	Session *mgo.Session
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) FindAvailable() (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(nil).One(&vessel)
	return vessel, err
}

// Close
func (repo *VesselRepository) Close() {
	repo.Session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.Session.DB(dbName).C(userCollection)
}
