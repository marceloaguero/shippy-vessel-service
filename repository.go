package main

import (
	pb "github.com/marceloaguero/shippy-vessel-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	Create(*pb.Vessel) error
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

// Create a new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

// Close closes the database session after each query
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}
