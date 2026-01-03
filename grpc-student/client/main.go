package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	studentpb "grpc-student/proto"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := studentpb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetStudent(ctx, &studentpb.StudentRequest{
		StudentId: 101,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Student Name:", resp.Name)
}
