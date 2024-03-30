package ch3ioreader

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func Run() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		fmt.Printf("can't connect: %v", err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		fmt.Printf("can't read response: %v", err)
	}
	defer res.Body.Close()
	fmt.Printf("Header: %v\n\n\n", res.Header)
	io.Copy(os.Stdout, res.Body)
}
