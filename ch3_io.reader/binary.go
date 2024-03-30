package ch3ioreader

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
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
	if bytes.Equal(buff, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buff bytes.Buffer
	binary.Write(&buff, binary.BigEndian, int32(len(byteData)))
	buff.WriteString("tEXt")
	buff.Write(byteData)
	// CRC を計算し追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buff, binary.BigEndian, crc.Sum32())
	return &buff
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

func AddTextChunk(src string, dist string) {
	file, err := os.Open(src)
	if err != nil {
		fmt.Printf("can't open file: %v", err)
	}
	defer file.Close()
	newFile, err := os.Create(dist)
	if err != nil {
		fmt.Printf("can't create file: %v", err)
	}
	defer newFile.Close()
	chunks := readChunk(file)
	// シグニチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要な IHDR チャンク書き込み
	io.Copy(newFile, chunks[0])
	// テキストチャンク追加
	io.Copy(newFile, textChunk("FooBar2000"))
	for _, chunk := range chunks {
		io.Copy(newFile, chunk)
	}
}
