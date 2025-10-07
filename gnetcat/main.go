package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ductran999/go-socket-poc/gnetcat/client"
	"github.com/ductran999/go-socket-poc/gnetcat/server"
	"github.com/ductran999/go-socket-poc/logger"
)

func main() {
	fmt.Println("Welcome to Go Net Cat")
	fmt.Println("1. Start Server")
	fmt.Println("2. Start Client")
	fmt.Print("=> Enter Your option:")

	var option int
	_, err := fmt.Scanln(&option)
	if err != nil {
		logger.Fatal("invalid input:", err.Error())
	}

	switch option {
	case 1:
		startServer()
	case 2:
		startClient()
	default:
		logger.Error("Invalid option. Please enter 1 or 2.")
	}
}

func startServer() {
	srv := server.NewServer()
	if err := srv.Open(); err != nil {
		logger.Fatal("failed to open server", err.Error())
	}
	logger.Info("server listening on port 8080")

	go func() {
		if err := srv.Serve(); err != nil && !errors.Is(err, net.ErrClosed) {
			logger.Error("server got error:", err.Error())
		}
	}()

	gracefulShutdown(srv.Close)
}

func startClient() {
	c := client.NewClient()
	if err := c.Dial(); err != nil {
		logger.Fatal("failed to dial to host", err.Error())
	}
	logger.Info("server accept your call")

	go func() {
		if err := c.Send(); err != nil &&
			!errors.Is(err, net.ErrClosed) &&
			!errors.Is(err, io.EOF) {
			logger.Error(err.Error())
		}
	}()

	gracefulShutdown(c.Close)
}

func gracefulShutdown(close func() error) {
	// wait for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.Info("Shutting down...")

	err := close()
	if err != nil {
		logger.Error("failed to shut down:", err.Error())
	} else {
		logger.Info("shutdown cleanly")
	}
}
