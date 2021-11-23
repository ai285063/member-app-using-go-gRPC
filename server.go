package main

import (
	"context"
	"log"
	"net"
	pb "github.com/ai285063/member_app_gRPC"
	grpc "google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	UnimplementedCRUDServer
}

func (s *server) GetUsers(ctx context.Context, in *GetUsersRequest) (*GetUsersResponse, error) {

	res := &GetUsersResponse{
		// Users: users,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterCRUDServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
