package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at port 8080.")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Printf("Accept: %v\n", conn.RemoteAddr())
			req, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
			res := http.Response{
				StatusCode: http.StatusOK,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
			}
			res.Write(conn)
			conn.Close()
		}()
	}
}
