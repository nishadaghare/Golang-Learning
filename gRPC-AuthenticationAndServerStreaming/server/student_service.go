// server/student_service.go
package main

import (
	"context"

	pb "grpc-student/proto"
)

type StudentService struct {
	pb.UnimplementedStudentServiceServer
}

func (s *StudentService) GetStudent(
	ctx context.Context,
	req *pb.StudentRequest,
) (*pb.StudentResponse, error) {

	return &pb.StudentResponse{
		Id:   req.Id,
		Name: "Rahul",
	}, nil
}

func (s *StudentService) ListStudents(
	req *pb.Empty,
	stream pb.StudentService_ListStudentsServer,
) error {

	students := []pb.StudentResponse{
		{Id: 1, Name: "Amit"},
		{Id: 2, Name: "Neha"},
		{Id: 3, Name: "Ravi"},
	}

	for _, student := range students {
		stream.Send(&student)
	}

	return nil
}
