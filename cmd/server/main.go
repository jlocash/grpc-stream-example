package main

import (
	"log"
	"net"

	"github.com/jlocash/grpc-stream-example/internal/proto"
	streamexample "github.com/jlocash/grpc-stream-example/internal/stream_example"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	se := streamexample.New()
	proto.RegisterStreamExampleServer(srv, se)
	srv.Serve(lis)
}
