package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lestrrat-go/server-starter/listener"
)

// このファイルを -o ./server としてビルド後に
// go run github.com/lestrrat-go/server-starter/cmd/start_server --port 8080 --pid-file app.pid ./server
// と実行、kill -HUP $(cat app.pid) のようにシグナルを送ると新しいWokerを起動し古いWorkerをグレイスフルシャットダウンする
//
// 以下のようなログ
//
// starting new worker 60763
//
// received HUP (num_old_workers=TODO)
// spawning a new worker (num_old_workers=TODO)
// starting new worker 60778
// new worker is now running, sending TERM to old workers:60763
// sleep 0 secs
// killing old workers
// old worker 60763 died, status:0
func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	<-signals
	server.Shutdown(context.Background())
}
