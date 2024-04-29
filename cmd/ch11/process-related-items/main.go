package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/shirou/gopsutil/process"
)

func main() {
	// 実行ファイル自身
	path, _ := os.Executable()
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("実行ファイルパス: %s\n", path)

	// プロセスID
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	// プロセスグループ・セッショングループ
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stdout, "グループID: %d セッションID: %d\n", syscall.Getpgrp(), sid)

	// ユーザ・ユーザグループ、実効ユーザ・ユーザグループ
	fmt.Printf("ユーザーID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)
	fmt.Printf("実効ユーザID: %d\n", os.Geteuid())
	fmt.Printf("実効ユーザグループID: %d\n", os.Getegid())

	// 作業ディレクトリ
	wd, _ := os.Getwd()
	fmt.Println(wd)

	// コマンドライン引数
	i := flag.Int("int", 1, "argument is int")
	s := flag.String("string", "str", "argument is string")
	b := flag.Bool("bool", true, "argument is bool")
	flag.Parse()
	fmt.Printf("command line args: %d, %v, %t\n", *i, *s, *b)

	// 環境変数,珍しいもの
	fmt.Println(os.ExpandEnv("${HOME}/gobin")) // 環境変数を展開してくれる

	// プロセス実行時の実行ファイル、引数などを取得
	p, _ := process.NewProcess(int32(os.Getppid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	fmt.Printf("parent pid: %d name: '%s' cmd: '%s'\n", p.Pid, name, cmd)
}
