package main

import (
	"github.com/pratikfuse/grpc-chat-server/client"
	"github.com/pratikfuse/grpc-chat-server/pb"
	"google.golang.org/grpc"
	"log"
)

func main(){

	connection, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to the server", err)
	}
	defer connection.Close()

	clientConnection := pb.NewChatClient(connection)
	client.Init(clientConnection)

}