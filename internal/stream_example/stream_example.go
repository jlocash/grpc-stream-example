package stream_example

import (
	"errors"
	"github.com/jlocash/grpc-stream-example/internal/proto"
	"time"
)

type StreamExample struct {
	clients map[string]chan string
}

func New() *StreamExample {
	return &StreamExample{
		clients: make(map[string]chan string),
	}
}

func (se *StreamExample) Subscribe(req *proto.SubscribeRequest, stream proto.StreamExample_SubscribeServer) error {
	uuid := req.GetUuid()
	if uuid == "" {
		return errors.New("uuid must be non-empty")
	}

	se.clients[uuid] = make(chan string)
	defer func() {
		close(se.clients[uuid])
		delete(se.clients, uuid)
	}()

	// simulate an event-driven operation
	go se.pingClient(uuid)

	for {
		msg := <-se.clients[uuid]
		nr := &proto.NotificationReply{
			Data: msg,
		}

		if err := stream.Send(nr); err != nil {
			return err
		}
	}
}

func (se *StreamExample) pingClient(uuid string) {
	for {
		time.Sleep(3 * time.Second)
		se.clients[uuid] <- "Hey man, whats up?"
	}
}
