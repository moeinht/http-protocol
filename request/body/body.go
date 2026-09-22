package body

import (
	"errors"
	"github.com/moeinht/http-protocol/request/header"
	"strconv"
)

var (
	ErrInvalidContentLength = errors.New("invalid content length")
)

func FindContentLength(headers []header.Header) (int64, error) {
	for _, value := range headers {
		if value.Name == "Content-Length" {
			contentSize, err := strconv.ParseInt(value.Value, 10, 64)
			if err != nil {
				return 0, err
			}

			if contentSize < 0 {
				return 0, ErrInvalidContentLength
			}

			return contentSize, nil
		}
	}

	return 0, nil
}
