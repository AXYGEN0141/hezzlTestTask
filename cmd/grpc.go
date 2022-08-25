package main

import (
	pb "github.com/AXYGEN0141/hezzlTestTask/protoUsers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	srv := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterUserServiceServer(srv, Server{})

	err = srv.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}

func (s *Server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Need to implement CreateUser
	panic("CreateUser")
}

func (s *Server) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Need to implement DeleteUser
	panic("DeleteUser")
}

func (s *Server) ListUser(ctx context.Context, request *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	// Need to implement ListUsers
	panic("ListUsers")
}
