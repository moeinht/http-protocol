package parser

import (
	readLine "http-protocol/parser/read-line"
	requestline "http-protocol/parser/request-line"
	"io"
)

type Parser struct {
	lineReader *readLine.LineReader
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		lineReader: readLine.NewLineReader(reader),
	}
}

func (p *Parser) ParseRequestLine() (requestline.RequestLine, error) {
	line, err := p.lineReader.Read()
	if err != nil {
		return requestline.RequestLine{}, err
	}

	return requestline.Parse(line)
}
