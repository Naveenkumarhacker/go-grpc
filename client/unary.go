package main

import (
	"context"
	pb "go-grpc/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, close := context.WithTimeout(context.Background(), time.Second)
	defer close()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Cound not greet :%v", err)
	}
	log.Printf("%s", res.Message)
}
