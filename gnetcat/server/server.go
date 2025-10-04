package server

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/ductran999/go-socket-poc/logger"
)

type GNetCatServer interface {
	// Server register port to listen on
	Open() error

	// Accept connect from client
	Serve() error

	// Stop listening
	Close() error
}

type server struct {
	srv net.Listener
}

func NewServer() GNetCatServer {
	return &server{}
}

func (s *server) Open() (err error) {
	s.srv, err = net.Listen("tcp", "localhost:8080")
	return
}

func (s *server) Serve() error {
	for {
		conn, err := s.srv.Accept()
		if err != nil {
			return fmt.Errorf("accept error: %w", err)
		}

		// Handle client in separate goroutine
		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 4096)

			for {
				n, err := c.Read(buf)
				if err != nil && errors.Is(err, io.EOF) {
					logger.Info("client", c.RemoteAddr().String(), "close connection")
					return
				}

				if err != nil {
					logger.Error("read error:", err.Error())
					return
				}

				msg := string(buf[:n])
				ans := fmt.Sprintf("Hello '%s'", msg)
				_, err = c.Write([]byte(ans))
				if err != nil {
					logger.Error("write error:", err.Error())
					return
				}
			}
		}(conn)
	}
}

func (s *server) Close() error {
	return s.srv.Close()
}
