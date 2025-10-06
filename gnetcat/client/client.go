package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type GNetCatClient interface {
	// Dial establishes a connection to the configured remote address.
	Dial() error

	// Send transmits data over the active connection.
	Send() error

	// Close gracefully terminates the connection and releases any resources.
	Close() error
}

type gClient struct {
	conn net.Conn
}

func NewClient() GNetCatClient {
	return &gClient{}
}

func (c *gClient) Dial() (err error) {
	c.conn, err = net.Dial("tcp", "localhost:8080")
	return
}

func (c *gClient) Close() error {
	return c.conn.Close()
}

func (c *gClient) Send() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("input error: %w", err)
		}
		input = strings.TrimSpace(input)

		// Send input to server
		_, err = c.conn.Write([]byte(input))
		if err != nil {
			return fmt.Errorf("client write error: %w", err)
		}

		// Read server response
		buf := make([]byte, 1024)
		n, err := c.conn.Read(buf)
		if err != nil {
			return fmt.Errorf("client read error: %w", err)
		}
		fmt.Println("<<", string(buf[:n]))
	}
}
