package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v': %v (%d bytes)\n", string(buffer), buffer, length)
}
func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader
	// 最初の 8 バイトを飛ばす
	file.Seek(0, 0)
	chunks = append(chunks, io.NewSectionReader(file, 0, 4))
	fmt.Printf("%v\n", io.NewSectionReader(file, 0, 4))
	// 次のチャンクの先頭に移動
	// 現在位置は長さを読み終わった箇所なので
	// チャンク名 (4 バイト ) + データ長 + CRC(4 バイト ) 先に移動
	stream := io.MultiReader(io.NewSectionReader(file, 0, 4))
	b := io.NewSectionReader(file, 0, 4)
	fmt.Println("---")
	io.Copy(os.Stdout, stream)
	fmt.Printf("%\n", stream)
	fmt.Println("---")
	io.Copy(os.Stdout, b)
	fmt.Println(b)
	fmt.Println("---")

	readBuf := make([]byte, 2)
	_, err := file.Read(readBuf)
	if err != nil {
		fmt.Println(err)
	}
	s := string(readBuf)
	fmt.Printf("-%v\n", s)

	fmt.Println("---")
	fmt.Printf("%b\n", s) // 2進表示
	fmt.Printf("%d\n", s) // 10進表示
	fmt.Printf("%x\n", s) // 16進小文字表示
	fmt.Printf("%X\n", s)

	return chunks
}
func main() {
	file, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		fmt.Printf("%v\n", chunk)
		dumpChunk(chunk)
	}
}
