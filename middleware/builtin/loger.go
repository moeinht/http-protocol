package builtin

import (
	"fmt"
	"time"

	"github.com/moeinht/http-protocol/handler"
	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
)

func Logger(next handler.IHandler) handler.IHandler {
	return func(req request.Request) (response.Response, error) {
		start := time.Now()

		fmt.Printf("→ %s %s\n",
			req.RequestLine.Method,
			req.RequestLine.Path,
		)

		res, err := next(req)

		fmt.Printf("← %s %s %d %s\n",
			req.RequestLine.Method,
			req.RequestLine.Path,
			res.StatusLine.Status,
			time.Since(start),
		)

		return res, err
	}
}
