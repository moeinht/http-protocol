package identifier

import (
	"bytes"
	"io"
	"strings"
)

type LineReader struct {
	reader io.Reader
	buffer []byte
}

func NewLineReader(reader io.Reader) *LineReader {
	return &LineReader{
		reader: reader,
		buffer: make([]byte, 0),
	}
}

func (lr *LineReader) Read() (string, error) {
	for {
		if index := bytes.IndexByte(lr.buffer, '\n'); index != -1 {
			line := string(lr.buffer[:index])

			lr.buffer = lr.buffer[index+1:]
			line = strings.TrimSuffix(line, "\r")
			return line, nil
		}

		readBuffer := make([]byte, 1024)

		n, err := lr.reader.Read(readBuffer)

		if n > 0 {
			lr.buffer = append(lr.buffer, readBuffer[:n]...)
		}

		if err != nil {
			return "", err
		}
	}
}
