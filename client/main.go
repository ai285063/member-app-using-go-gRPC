package main

import (
	"context"
	"log"

	"github.com/ai285063/member_app_gRPC/proto/memberApp"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := memberApp.NewCRUDClient(conn)

	doGetUsers(client)
	doCreateUser(client)

}

func doGetUsers(client memberApp.CRUDClient) {
	req := &memberApp.GetUsersRequest{}
	res, err := client.GetUsers(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling doGetUsers: %v \n", err)
	}
	log.Printf("Response from GetUsers: %v\n", res.Users)
}

func doCreateUser(client memberApp.CRUDClient) {

}
