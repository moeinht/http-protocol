package requestline

import (
	"errors"
	"strings"
)

var (
	ErrInvalidRequestLine = errors.New("invalid request line")
	ErrInvalidMethod      = errors.New("invalid method")
	ErrInvalidPath        = errors.New("invalid path")
)

type RequestLine struct {
	Method string
	Path   string
}

func Tokenizer(line string) ([]string, error) {

	spaceIndex := strings.Index(line, " ")
	fields := make([]string, 0, 2)
	if spaceIndex == -1 {
		return fields, ErrInvalidRequestLine
	}

	method := line[:spaceIndex]
	path := line[spaceIndex+1:]

	if method == "" || path == "" {
		return fields, ErrInvalidRequestLine
	}

	if strings.Contains(path, " ") {
		return fields, ErrInvalidRequestLine
	}

	fields = append(fields, method)
	fields = append(fields, path)

	return fields, nil
}

func Parse(line string) (RequestLine, error) {
	fileds, err := Tokenizer(line)

	if err != nil {
		return RequestLine{}, err
	}

	method := fileds[0]
	path := fileds[1]

	if !validMethod(method) {
		return RequestLine{}, ErrInvalidMethod
	}

	if path == "" {
		return RequestLine{}, ErrInvalidPath
	}

	return RequestLine{
		Method: method,
		Path:   path,
	}, nil

}

func validMethod(method string) bool {
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
		return true
	default:
		return false
	}
}
