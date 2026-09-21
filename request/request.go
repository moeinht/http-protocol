package request

import (
	"http-protocol/request/header"
	"http-protocol/request/identifier"
	requestline "http-protocol/request/request-line"
	"io"
)

type Request struct {
	RequestLine requestline.RequestLine
	Headers     []header.Header
	Body        []byte
}

type Parser struct {
	lineReader *identifier.LineReader
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		lineReader: identifier.NewLineReader(reader),
	}
}

func (p *Parser) ParseRequestLine() (requestline.RequestLine, error) {
	line, err := p.lineReader.Read()
	if err != nil {
		return requestline.RequestLine{}, err
	}

	return requestline.Parse(line)
}

func (p *Parser) ParseHeaders() ([]header.Header, error) {
	headers := []header.Header{}

	for {
		line, err := p.lineReader.Read()
		if err != nil {
			return nil, err
		}

		// Empty line means headers are finished.
		if line == "" {
			return headers, nil
		}

		h, err := header.Parse(line)
		if err != nil {
			return nil, err
		}

		headers = append(headers, h)
	}
}

func (p *Parser) Parse() (Request, error) {
	requestLine, err := p.ParseRequestLine()
	if err != nil {
		return Request{}, err
	}

	headers, err := p.ParseHeaders()
	if err != nil {
		return Request{}, err
	}

	return Request{
		RequestLine: requestLine,
		Headers:     headers,
	}, nil
}
