package main

import (
	"context"
	"log"
	"time"

	pb "gopulse/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewJobServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SubmitJob(ctx, &pb.JobRequest{
		Payload: "hello",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Job ID:", res.JobId)
}