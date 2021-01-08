package main

import (
	"github.com/pratikfuse/grpc-chat-server/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	log.Println("Starting server")
	sigIntChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(sigIntChan, os.Interrupt, syscall.SIGTERM)
	srv := server.GetApp()
	go srv.Start(errChan)
	select {
	case err := <-errChan:
		log.Fatal(err)
	case <-sigIntChan:
		srv.Graceful()
	}
}