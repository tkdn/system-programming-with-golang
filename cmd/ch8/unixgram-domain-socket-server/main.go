package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	os.Remove(path)
	fmt.Println("Server is running at " + path)
	conn, err := net.ListenPacket("unixgram", path)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buff := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buff)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buff[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
