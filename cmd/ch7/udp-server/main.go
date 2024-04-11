package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8080")
	conn, err := net.ListenPacket("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buff := make([]byte, 32)
	for {
		length, remoteAddress, err := conn.ReadFrom(buff)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buff[:length]))
		if _, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress); err != nil {
			panic(err)
		}
	}
}
