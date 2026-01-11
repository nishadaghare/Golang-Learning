// server/main.go
package main

import (
	"log"
	"net"

	pb "gRPC-AuthenticationAndServerStreaming/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	server := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
	)

	pb.RegisterStudentServiceServer(server, &StudentService{})

	log.Println("Server running on :50051")
	server.Serve(lis)
}
