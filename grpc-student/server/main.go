package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	studentpb "grpc-student/proto"
)

type server struct {
	studentpb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(
	ctx context.Context,
	req *studentpb.StudentRequest,
) (*studentpb.StudentResponse, error) {

	log.Println("Received student_id:", req.StudentId)

	return &studentpb.StudentResponse{
		Id:   req.StudentId,
		Name: "Nishada Ghare",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on :50051")
	grpcServer.Serve(lis)
}
