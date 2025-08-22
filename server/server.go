package main

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"net"

	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedHelloServerServer
}

func main() {

	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer() // engine
	proto.RegisterHelloServerServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("receive request from client", req.Message)
	fmt.Println("hello from server")
	return &proto.HelloResponse{}, errors.New("")
}