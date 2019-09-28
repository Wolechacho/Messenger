package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "messenger/messenger"

	"google.golang.org/grpc"
)

const (
	port = ":50055"
)

type server struct{}

func (srv *server) SendMessage(ctx context.Context, in *pb.InstructionRequest) (*pb.InstructionResponse, error) {
	return &pb.InstructionResponse{
		Response: fmt.Sprintf("Hello i received message with the title : %s and the content : %s", in.Title, in.Body),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error occured creating port %s\n .. StackTrace %v\n", port, err)
	}

	//create the grpc server
	s := grpc.NewServer()

	//register the rpc functions to the grpc server using the IDK
	pb.RegisterMessengerServer(s, &server{})

	//make the grpc server listen on the above port created
	fmt.Println("Server Listening on port : ", port)

	er := s.Serve(listener)
	if er != nil {
		log.Fatalf("Error occured serving grpc server on port %s\n .. StackTrace %v\n", port, er)
	}
}
