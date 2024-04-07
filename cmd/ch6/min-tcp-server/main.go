package main

import (
	"fmt"
	"net"
)

// * Received HTTP/0.9 when not allowed
// * Closing connection
// curl: (1) Received HTTP/0.9 when not allowed
// となるので以下のように確認
// curl -v --http0.9 localhost:8080
func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("handle error: %v\n", err)
		}
		go func() {
			conn.Write([]byte("FooBar"))
			conn.Close()
		}()
	}
}
