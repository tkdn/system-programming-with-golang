package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	os.Remove(clientPath)
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	// var serverAddr net.Addr = unixServerAddr
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sending to Server")
	_, err = conn.WriteTo([]byte("Hello from Client"), unixServerAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from server")
	buff := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buff)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buff[:length]))
}
