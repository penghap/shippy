package repository

import (
	pb "github.com/penghap/shippy/service-consignment/proto/consignment"

	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	Session *mgo.Session
}

// Create consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll consignment
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// Close
func (repo *ConsignmentRepository) Close() {
	repo.Session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.Session.DB(dbName).C(consignmentCollection)
}
