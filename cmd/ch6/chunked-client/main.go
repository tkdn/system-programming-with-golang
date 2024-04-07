package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	err = req.Write(conn)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	res, err := http.ReadResponse(reader, req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	if len(res.TransferEncoding) < 1 || res.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encodig")
	}
	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		// 16進数のサイズをパースし、サイズがゼロならクローズ
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		// サイズ分バッファを確保して読み込み
		line := make([]byte, int(size))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf("  %d bytes: %s\n", size, string(line))
	}
}
