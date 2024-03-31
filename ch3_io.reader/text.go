package ch3ioreader

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行目 Foo
2行目 Bar
3行目 2000`

func ReadText() {
	r := bufio.NewReader(strings.NewReader(source))
	for {
		l, err := r.ReadString('\n')
		fmt.Printf("%#v\n", l)
		if err == io.EOF {
			break
		}
	}
}

// bufio.Scanner を使ってテキスト解析を行う場合は
// 分割文字(ここでは改行記号)が削除される、実行結果でよくわかる
func ScanText() {
	s := bufio.NewScanner(strings.NewReader(source))
	for s.Scan() {
		fmt.Printf("%#v\n", s.Text())
	}
}

// scanner.Split(bufio.ScanWords) とすることで単語分割が可能(スペース文字を分割文字)
func ScanTextByWord() {
	s := bufio.NewScanner(strings.NewReader(source))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		fmt.Printf("%#v\n", s.Text())
	}
}
