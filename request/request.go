package request

import (
	"io"

	"github.com/moeinht/http-protocol/request/body"
	"github.com/moeinht/http-protocol/request/header"
	"github.com/moeinht/http-protocol/request/identifier"
	requestline "github.com/moeinht/http-protocol/request/request-line"
)

type Request struct {
	RequestLine requestline.RequestLine
	Headers     []header.Header
	Body        []byte
}

type Parser struct {
	lineReader *identifier.BufferedReader
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		lineReader: identifier.NewBufferedReader(reader),
	}
}

func (p *Parser) ParseRequestLine() (requestline.RequestLine, error) {
	line, err := p.lineReader.ReadLine()
	if err != nil {
		return requestline.RequestLine{}, err
	}

	return requestline.Parse(line)
}

func (p *Parser) ParseHeaders() ([]header.Header, error) {
	headers := []header.Header{}

	for {
		line, err := p.lineReader.ReadLine()
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

func (p *Parser) ParseBody(headers []header.Header) ([]byte, error) {
	contentLength, err := body.FindContentLength(headers)

	if err != nil {
		return nil, nil
	}

	body, err := p.lineReader.ReadN(contentLength)

	if err != nil {
		return nil, nil
	}

	return body, nil
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

	body, err := p.ParseBody(headers)

	if err != nil {
		return Request{}, err
	}

	return Request{
		RequestLine: requestLine,
		Headers:     headers,
		Body:        body,
	}, nil
}
