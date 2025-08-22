package main

import (
	"context"
	proto "grpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewHelloServerClient(conn)

	req := &proto.HelloRequest{Message: "hello from client"}

	client.ServerReply(context.TODO(), req)
}