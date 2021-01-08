package server

import (
	"context"
	"github.com/pratikfuse/grpc-chat-server/pb"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	stream pb.Chat_ChatServer
}

type App struct {
	GrpcServer *grpc.Server
	errChan chan error
	Clients []*Client
	register chan *Client
	unregister chan *Client
}

func (s *App) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	panic("implement me")
}



func (s *App) Send(client *Client){

	client.stream.
}


func (s *App) Receive(client *Client){
	for {
		message, err := client.stream.Recv()
		if err != nil {
			s.errChan <- err
		}

		log.Printf("%s:%s", message.Username, message.ChatMessage)
	}
}

func (s *App) Chat(server pb.Chat_ChatServer) error {
	client := &Client{stream: server}
	s.Clients = append(s.Clients, client)
	go s.Send(client)
	go s.Receive(client)

	if err := <- s.errChan; err != nil {
		return err
	}
	return nil
}
