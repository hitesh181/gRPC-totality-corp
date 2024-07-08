package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "grpc-user-service/proto"
)

func main() {
    // Establish a connection to the server.
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

   
    c := pb.NewUserServiceClient(conn)

    
    getUserResponse, err := c.GetUser(context.Background(), &pb.GetUserRequest{Id: 2})
    if err != nil {
        log.Fatalf("could not get user: %v", err)
    }
    log.Printf("User: %v", getUserResponse.User)

    
    listUsersResponse, err := c.ListUsers(context.Background(), &pb.ListUsersRequest{Ids: []int32{1,2,3,4}})
    if err != nil {
        log.Fatalf("could not list users: %v", err)
    }
    log.Printf("Users: %v", listUsersResponse.Users)

  
    searchUsersResponse, err := c.SearchUsers(context.Background(), &pb.SearchUsersRequest{City: "CA"})
    if err != nil {
        log.Fatalf("could not search users: %v", err)
    }
    log.Printf("Users: %v", searchUsersResponse.Users)
}
