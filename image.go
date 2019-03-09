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

	var length int32
	err0 := binary.Read(file, binary.BigEndian, &length)
	if err0 == io.EOF {
		panic(err0)
	}
	fmt.Printf("%d\n", int64(length))

	fmt.Println("---")
	io.Copy(os.Stdout, stream)
	fmt.Printf("%\n", stream)
	fmt.Println("---")

	return chunks
}

func readChunks3(r io.ReaderAt) {
	fmt.Println("====================")
	b := io.NewSectionReader(r, int64(0), 1000)
	aa := int64(1<<63 - 1)
	fmt.Printf("-%d\n", aa)

	io.Copy(os.Stdout, b)
	fmt.Println(b)
	fmt.Println("---")
	readBuf := make([]byte, int64(40))
	_, err := b.Read(readBuf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(readBuf)) // =&gt; re
	fmt.Println("====================")
	fmt.Printf("%x\n", string(readBuf)) // 16進小文字表示
	fmt.Printf("%X\n", string(readBuf))
	fmt.Println("====================")
	s := string(readBuf)
	fmt.Printf("-%v\n", s)

	fmt.Println("---")
	fmt.Printf("%b\n", s) // 2進表示
	fmt.Printf("%d\n", s) // 10進表示
	fmt.Printf("%x\n", s) // 16進小文字表示
	fmt.Printf("%X\n", s)
	fmt.Println("====================")
}

func readChunks2(file *os.File) {
	// 最初の 8 バイトを飛ばす
	file.Seek(0, 0)

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
}

func main() {
	file, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	//     readChunks2(file)
	readChunks3(file)
	for _, chunk := range chunks {
		fmt.Printf("%v\n", chunk)
		dumpChunk(chunk)
	}
}
