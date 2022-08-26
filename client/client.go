package main

import (
	"fmt"
	pb "github.com/AXYGEN0141/hezzlTestTask/protoUsers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5432", grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx := context.Background()

	createRequest := &pb.CreateUserRequest{
		User: &pb.User{
			Id:   1,
			Name: "Danila",
		},
	}

	listRequest := &pb.ListUserRequest{}

	createResp, err := client.CreateUser(ctx, createRequest)
	listResp, err := client.ListUser(ctx, listRequest)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(createResp, listResp)
}
