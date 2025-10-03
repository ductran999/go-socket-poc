package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	// create listener
	srv, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer srv.Close()

	// start client connection
	client, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer client.Close()

	var wg sync.WaitGroup

	// server goroutine
	wg.Go(func() {
		conn, err := srv.Accept()
		if err != nil {
			log.Println("accept error:", err)
			return
		}
		defer conn.Close()

		n, err := conn.Write([]byte("duc"))
		if err != nil {
			fmt.Println("write error:", err)
		}
		fmt.Println("server wrote bytes:", n)
	})

	// client goroutine
	wg.Go(func() {
		buf := make([]byte, 1024)
		defer client.Close()
		for {
			n, err := client.Read(buf)
			if err != nil {
				log.Println("client read error:", err)
				return
			}
			fmt.Println("client received:", string(buf[:n]))
		}
	})

	wg.Wait()
}
