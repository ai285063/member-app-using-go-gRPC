package main

import (
	"context"
	"log"
	"net"

	memberApp "github.com/ai285063/member_app_gRPC/proto/memberApp"
	grpc "google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	memberApp.UnimplementedCRUDServer
}

func (s *server) GetUsers(ctx context.Context, req *memberApp.GetUsersRequest) (*memberApp.GetUsersResponse, error) {
	var users []*memberApp.User
	if err := MysqlDB.Table("users").Find(&users).Error; err != nil {
		panic(err)
	}

	// manually cast *gorm to *memberApp.User
	usersArr := make([]*memberApp.User, 0)

	for _, u := range users {
		usersArr = append(usersArr, &memberApp.User{
			Id:       u.Id,
			Account:  u.Account,
			Email:    u.Email,
			Password: u.Password,
		})
	}
	res := &memberApp.GetUsersResponse{
		Users: usersArr,
	}

	return res, nil
}

func main() {
	ConnectMysql()
	ConnectRedis()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	memberApp.RegisterCRUDServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
