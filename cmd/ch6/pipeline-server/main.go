package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
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
		go processSession(conn)
	}
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	// セッション内のリクエストを処理するチャンネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// レスポンスを直列化してソケットに書き出すgoroutine
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)

	// レスポンスを受取りセッションキューに入れる
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエスト読み込み
		req, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse
		go handleRequest(req, sessionResponse)
	}
}

func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for sessionResponse := range sessionResponses {
		res := <-sessionResponse
		res.Write(conn)
		close(sessionResponse)
	}
}

func handleRequest(req *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	content := "Hello World\n"
	// セッションの維持のためKeep-Alive
	res := &http.Response{
		StatusCode:    http.StatusOK,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          io.NopCloser(strings.NewReader(content)),
	}
	// チャネルに書き込み, ブロックしているwriteToConnの処理を再開
	resultReceiver <- res
}
