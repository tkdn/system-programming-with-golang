package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

// see: https://github.com/tkdn/system-programming-with-golang/blob/c4cabe08f00ec3dc8f04b26e35524a171dd4ee99/ch3_io.reader/conn.go#L12
// User-Agent: Go-http-client/1.1
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
