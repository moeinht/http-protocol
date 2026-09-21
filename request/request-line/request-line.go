package requestline

import (
	"errors"
	"strings"
)

type RequestLine struct {
	Method  string
	Path    string
	Version string
}

var (
	ErrInvalidRequestLine = errors.New("invalid request line")
)

func Tokenizer(line string) ([]string, error) {
	spaceIndex := strings.Index(line, " ")
	lastSpace := strings.LastIndex(line, " ")

	fields := make([]string, 0, 3)

	if spaceIndex == -1 || lastSpace == -1 {
		return fields, ErrInvalidRequestLine
	}

	method := line[:spaceIndex]
	path := line[spaceIndex+1 : lastSpace]
	version := line[lastSpace+1:]

	if method == "" || path == "" || version == "" {
		return fields, ErrInvalidRequestLine
	}

	if strings.Contains(path, " ") {
		return fields, ErrInvalidRequestLine
	}

	fields = append(fields, method)
	fields = append(fields, path)
	fields = append(fields, version)

	return fields, nil
}

func Parse(line string) (RequestLine, error) {
	fields, err := Tokenizer(line)
	if err != nil {
		return RequestLine{}, err
	}

	method := fields[0]
	path := fields[1]
	version := fields[2]

	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
	default:
		return RequestLine{}, ErrInvalidRequestLine
	}

	if path == "" {
		return RequestLine{}, ErrInvalidRequestLine
	}

	if version != "HTTP/1.1" {
		return RequestLine{}, ErrInvalidRequestLine
	}

	return RequestLine{
		Method:  method,
		Path:    path,
		Version: version,
	}, nil
}
