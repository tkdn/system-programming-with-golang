package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn
	var err error
	reqs := make(chan *http.Request, len(sendMessages))

	conn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access: %d\n", current)
	defer conn.Close()

	// リクエストを先に送る
	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		req, err := http.NewRequest("GET", "http://localhost:8080?message="+sendMessages[i], nil)
		if lastMessage {
			req.Header.Add("Connection", "close")
		} else {
			req.Header.Add("Connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}
		if err = req.Write(conn); err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessages[i])
		reqs <- req
	}
	close(reqs)

	// レスポンスをまとめて受信する
	reader := bufio.NewReader(conn)
	for req := range reqs {
		res, err := http.ReadResponse(reader, req)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		if current == len(sendMessages) {
			break
		}
	}
}
