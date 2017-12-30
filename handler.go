package main

import (
	"context"

	pb "github.com/marceloaguero/shippy-vessel-service/proto"
	mgo "gopkg.in/mgo.v2"
)

/* service struct implements the defined protobuf interface.
// Server API for VesselService service

type VesselServiceServer interface {
	FindAvailable(context.Context, *Specification) (*Response, error)
	Create(context.Context, *Vessel) (*Response, error)
}
*/

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) Create(ctx context.Context, req *pb.Specification) (*pb.Response, error) {
	repo := s.GetRepo()
	defer repo.Close()

	// Save our vessel
	if err := repo.Create(req); err != nil {
		return nil, err
	}

	return &pb.Response{Vessel: req, Created: true}, nil
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification) (*pb.Response, error) {
	repo := s.GetRepo()
	defer repo.Close()

	// Find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Vessel: vessel}, nil
}
