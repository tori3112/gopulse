package main

import (
	"context"
	"log"
	"net"

	pb "gopulse/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedJobServiceServer
}

func (s *server) SubmitJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	log.Println("Received job:", req.Payload)
	return &pb.JobResponse{JobId: "123"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterJobServiceServer(grpcServer, &server{})

	log.Println("Server running on :50051")
	grpcServer.Serve(lis)
}