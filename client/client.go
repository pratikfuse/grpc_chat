package client

import (
	"bufio"
	"fmt"
	"github.com/pratikfuse/grpc-chat-server/pb"
	"log"
	"os"
)

func Chat(username string) {
	stream, err := Client.Chat(Ctx)
	errChan := make(chan error)
	if err != nil {
		log.Fatal(err)
	}
	go SendMessage(stream.Send, username,errChan )
	go ReceiveMessage(stream.Recv, errChan)

	select {
	case done := <-stream.Context().Done():
		fmt.Println(done)
	case err := <-stream.Context().Done():
		fmt.Println(err)
	case err := <- errChan:
		log.Fatal(err)
	}

}

func SendMessage(send func(message *pb.ChatMessage) error, username string, errorChan chan <- error) {
	reader  := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s: ", username)
		text, err := reader.ReadString('\n')
		if err != nil {
			errorChan <- err
		}
		message := &pb.ChatMessage{
			Username:    username,
			ChatMessage: text,
		}
		err = send(message)
		if err != nil {
			errorChan <- err
		}
	}
}

func ReceiveMessage(recv func() (*pb.ChatMessage, error), errorChan chan <- error) {
	for {
		message, err := recv()
		if err != nil {
			errorChan <- err
		}
		log.Printf("%s: %s \n", message.Username, message.ChatMessage)
	}
}
