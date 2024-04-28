package main

import (
	"fmt"
	"syscall"
)

func main() {
	kq, err := syscall.Kqueue()
	if err != nil {
		panic(err)
	}
	// 監視対象のfdをを取得
	fd, err := syscall.Open("./.temp", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	// 監視するイベントの構造体を定義
	ev1 := syscall.Kevent_t{
		Ident:  uint64(fd),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE,
		Data:   0,
		Udata:  nil,
	}
	// イベントを待ち受け(ブロッキング)
	for {
		events := make([]syscall.Kevent_t, 10)
		nev, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		if err != nil {
			panic(err)
		}
		// イベントの確認
		for i := 0; i < nev; i++ {
			fmt.Printf("Event [%d] -> %+v\n", i, events[i])
		}
	}
}
