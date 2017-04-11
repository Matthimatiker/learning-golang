package ReverseReader

import (
	"os"
	"io"
)

type ReverseReader struct {
	file *os.File
	pos  int64
}

func NewReverseReader(file *os.File) *ReverseReader {
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	return &ReverseReader{
		file: file,
		pos:  info.Size(),
	}
}

func (reader *ReverseReader) Read(p []byte) (int, error) {
	numberOfBytesToRead := len(p)
	if int(reader.pos) < numberOfBytesToRead {
		// Not enough bytes left, read less
		numberOfBytesToRead = int(reader.pos)
	}
	if numberOfBytesToRead == 0 {
		// No more bytes to read. Send io.EOF to indicate end of file.
		return 0, io.EOF
	}
	reader.pos -= int64(numberOfBytesToRead)
	// Read the chunk...
	buffer := make([]byte, numberOfBytesToRead, numberOfBytesToRead)
	_, err := reader.file.ReadAt(buffer, reader.pos)
	// ...and pass provide the bytes in reverse order.
	for i := 0; i < numberOfBytesToRead; i++ {
		p[i] = buffer[numberOfBytesToRead-i-1]
	}
	return numberOfBytesToRead, err
}
