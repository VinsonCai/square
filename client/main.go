package main

import (
	"context"
	"log"

	"square/proto" // 相对路径引用生成的 proto 包

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewSquareCaculatorClient(conn)
	resp, err := c.SayHello(context.Background(), &proto.SquareRequest{Factor: 5})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Printf("Result: %d", resp.Result)
}
