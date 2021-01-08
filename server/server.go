package server

import (
	"github.com/pratikfuse/grpc-chat-server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func GetApp() *App{
	srv := grpc.NewServer()
	app := &App{
		GrpcServer: srv,
	}
	pb.RegisterChatServer(srv, app)
	return app
}

func (s *App) Start(errorChan chan <- error)   {
	log.Println("Started server at :9000")
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		errorChan <- err
	}else {
		defer listener.Close()
	}
	if err := s.GrpcServer.Serve(listener); err != nil {
		errorChan <- err
	}

	errorChan <- err
}

func (s *App) Graceful() {
	log.Println("Gracefully stopping chat server")
	if s.GrpcServer != nil {
		s.GrpcServer.GracefulStop()
		os.Exit(0)
	}
}
