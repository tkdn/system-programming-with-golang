package ch3ioreader

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func readChunk(file *os.File) []io.Reader {
	var chunks []io.Reader
	// 先頭8byteをとばす
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンク先頭へ移動
		// 現在地点は長さを読み終えた場所、チャンク名(4byte) + データ長 + CRC(4byte)先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buff := make([]byte, 4)
	chunk.Read(buff)
	fmt.Printf("chunk %q (%d bytes)\n", string(buff), length)
}

func ReadPng(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("can't open file: %v", err)
	}
	defer file.Close()
	chunks := readChunk(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
