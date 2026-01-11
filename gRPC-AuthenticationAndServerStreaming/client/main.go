// client/main.go
package main

import (
	"context"
	"log"
	"time"

	pb "gRPC-AuthenticationAndServerStreaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	// Metadata (AUTH)
	md := metadata.New(map[string]string{
		"authorization": "Bearer secret123",
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := client.GetStudent(ctx, &pb.StudentRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
	time.Sleep(time.Second)

	stream, err := client.ListStudents(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		log.Println(resp)
	}

}
