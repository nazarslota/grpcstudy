package main

import (
	"net"

	pb "github.com/nazarslota/grpc-study/greet/proto"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:5051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
