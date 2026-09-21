package identifier

import (
	"bytes"
	"io"
	"strings"
)

type BufferedReader struct {
	reader io.Reader
	buffer []byte
}

func NewBufferedReader(reader io.Reader) *BufferedReader {
	return &BufferedReader{
		reader: reader,
		buffer: make([]byte, 0),
	}
}

func (br *BufferedReader) ReadLine() (string, error) {
	for {
		if index := bytes.IndexByte(br.buffer, '\n'); index != -1 {
			line := string(br.buffer[:index])

			br.buffer = br.buffer[index+1:]
			line = strings.TrimSuffix(line, "\r")
			return line, nil
		}

		readBuffer := make([]byte, 1024)

		n, err := br.reader.Read(readBuffer)

		if n > 0 {
			br.buffer = append(br.buffer, readBuffer[:n]...)
		}

		if err != nil {
			return "", err
		}
	}
}
func (br *BufferedReader) ReadN(byteSize int64) ([]byte, error) {
	if byteSize == 0 {
		return []byte{}, nil
	}

	bodyData := make([]byte, 0, byteSize)

	for int64(len(bodyData)) < byteSize {

		remaining := byteSize - int64(len(bodyData))

		if len(br.buffer) > 0 {
			bufferSize := int64(len(br.buffer))

			if bufferSize > remaining {
				bufferSize = remaining
			}

			bodyData = append(bodyData, br.buffer[:bufferSize]...)
			br.buffer = br.buffer[bufferSize:]

			continue
		}

		readBuffer := make([]byte, remaining)

		n, err := br.reader.Read(readBuffer)

		if n > 0 {
			bodyData = append(bodyData, readBuffer[:n]...)
		}

		if err != nil {
			if int64(len(bodyData)) == byteSize {
				return bodyData, nil
			}

			return nil, io.ErrUnexpectedEOF
		}
	}

	return bodyData, nil
}
