package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"square/proto"
	pb "square/proto" // 相对路径引用生成的 proto 包

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSquareCaculatorServer
}

func (s *server) SayHello(ctx context.Context, in *proto.SquareRequest) (*proto.SquareReply, error) {
	fmt.Println("accept request")
	result := in.Factor * in.Factor
	return &proto.SquareReply{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSquareCaculatorServer(s, &server{})
	log.Println("Server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
