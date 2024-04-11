package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Sending to Server")
	if _, err = conn.Write([]byte("Hello from Client, Hello from Client")); err != nil {
		panic(err)
	}
	fmt.Println("Receiving from Server")
	buff := make([]byte, 32)
	length, err := conn.Read(buff)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Recived: %s\n", string(buff[:length]))
}
