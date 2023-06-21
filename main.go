package main

import (
	"context"
	"log"
	"net"

	"github.com/pedramkouasri/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type InventoryClient struct {
	pb.UnimplementedInventoryServer
}

func (s InventoryClient) GetBookList(context.Context, *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func getSampleBooks() []*pb.Book {
	sampleBook := []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R. Tolkien",
			PageCount: 1234,
		},
	}

	return sampleBook
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterInventoryServer(s, &InventoryClient{})

	log.Println("listen on port 8081")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("listen on port 8080")

}
