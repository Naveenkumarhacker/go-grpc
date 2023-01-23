package main

import (
	pb "go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Naveenkumar", "alice", "bob"},
	}

	// Unary
	callSayHello(client)

	// server streaming
	CallSayHelloServerStream(client, names)

	// Client streaming
	CallSayHelloClientStream(client, names)

	// bidirectional streaming
	CallSayHelloBidirectionalStream(client, names)
}
