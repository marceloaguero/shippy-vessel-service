package main

import (
	"log"
	"net"
	"os"

	pb "github.com/marceloaguero/shippy-vessel-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultHost = "db:27017"
	port        = ":50052"
)

func CreateDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	// Database host from environment
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	repo := &VesselRepository{session.Copy()}
	CreateDummyData(repo)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterVesselServiceServer(s, &service{session})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
