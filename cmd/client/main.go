package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jlocash/grpc-stream-example/internal/proto"
	"google.golang.org/grpc"
)

func initializeStreamExampleClient() (proto.StreamExampleClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return proto.NewStreamExampleClient(conn), conn, nil
}

func main() {
	client, conn, err := initializeStreamExampleClient()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	req := &proto.SubscribeRequest{
		Uuid: uuid.New().String(),
	}

	streamClient, err := client.Subscribe(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := streamClient.Recv()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(msg.GetData())
	}
}
