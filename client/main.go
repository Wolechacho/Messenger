package main

import (
	"context"
	"log"
	"time"

	pb "messenger/messenger"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50055"
)

func main() {

	clientconn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error occured dialing to the server ..StackTrace : %v", err)
	}

	//close the client connection
	defer func() {
		clientconn.Close()
	}()

	msgClient := pb.NewMessengerClient(clientconn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	//close the cancellation
	defer func() {
		cancel()
	}()

	r, er := msgClient.SendMessage(ctx, &pb.InstructionRequest{To: "Grpc Server", Body: "Hello for Golang grpc client", Title: "Saying Hi"})
	if er != nil {
		log.Fatalf("could not get response %v", er)
	}
	log.Printf("Response: %s", r.Response)
}
