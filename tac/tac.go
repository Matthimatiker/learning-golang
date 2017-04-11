package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

type ReverseReader struct {
	file *os.File;
	pos int64
}

func NewReverseReader(file *os.File) *ReverseReader {
	info, err := file.Stat()
	if (err != nil) {
		panic(err)
	}
	return &ReverseReader{
		file: file,
		pos: info.Size(),
	}
}

func (reader *ReverseReader) Read(p []byte) (int, error) {
	numberOfBytesToRead := len(p)
	if (int(reader.pos) < numberOfBytesToRead) {
		// Not enough bytes left, read less
		numberOfBytesToRead = int(reader.pos)
	}
	if numberOfBytesToRead == 0 {
		// No more bytes to read. Send io.EOF to indicate end of file.
		return 0, io.EOF
	}
	reader.pos -= int64(numberOfBytesToRead)
	_, err := reader.file.ReadAt(p, reader.pos)
	return numberOfBytesToRead, err
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := NewReverseReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
