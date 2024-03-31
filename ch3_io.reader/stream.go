package ch3ioreader

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// io.MultiReader を使い、複数の io.Reader 入力をつなげるよう動作
func MultiRead() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	body := bytes.NewBufferString("FooBar200\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	r := io.MultiReader(header, body, footer)
	io.Copy(os.Stdout, r)
}

// io.TeeReader で io.Writer に読み込まれたデータを書き出す
func TeeRead() {
	var buff bytes.Buffer
	r := bytes.NewBufferString("Example of io.TeeReader\n")
	tr := io.TeeReader(r, &buff)
	// データ読み捨て
	_, _ = io.ReadAll(tr)
	fmt.Println(buff.String())
}
