package main

import (
	"context"
	pb "grpc-user-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []pb.User
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.GetUserResponse{User: &user}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var userList []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				userList = append(userList, &user)
			}
		}
	}
	return &pb.ListUsersResponse{Users: userList}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var userList []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == user.Married) {
			userList = append(userList, &user)
		}
	}
	return &pb.SearchUsersResponse{Users: userList}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "Peter", City: "CA", Phone: 8978678978, Height: 5.9, Married: false},
			{Id: 3, Fname: "John", City: "NY", Phone: 9878675645, Height: 6.2, Married: false},
			{Id: 4, Fname: "Angela", City: "UK", Phone: 78676578234, Height: 5.8, Married: true},
			
			
			// Add more users as needed
		},
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
