package client

import (
	"context"
	"flag"
	"github.com/pratikfuse/grpc-chat-server/pb"
	"log"
	"os"
)

var Client pb.ChatClient
var Ctx = context.Background()

func Init(clientConnection pb.ChatClient){
	connectFlags := flag.NewFlagSet("connect", flag.ExitOnError)
	username := connectFlags.String("u", "", "u")

	Client = clientConnection
	switch os.Args[1]{
	case "login":
		fallthrough
	case "connect":
		err := connectFlags.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		Chat(*username)
	}

}
