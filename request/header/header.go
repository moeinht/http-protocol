package header

import (
	"errors"
	"strings"
)

type Header struct {
	Name  string
	Value string
}

var (
	ErrInvalidHeader = errors.New("Invalid Syntax")
)

func Parse(line string) (Header, error) {
	firstColon := strings.Index(line, ": ")

	header := Header{}

	if firstColon == -1 {
		return header, ErrInvalidHeader
	}

	name := line[:firstColon]
	value := line[firstColon+2:]

	if len(name) == 0 || len(value) == 0 {
		return header, ErrInvalidHeader
	}

	header.Name = name
	header.Value = value

	return header, nil
}
