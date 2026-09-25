package middleware

import "github.com/moeinht/http-protocol/handler"

type Chain struct {
	middlewares []Middleware
}

func New() *Chain {
	return &Chain{}
}

func (c *Chain) Use(mw Middleware) *Chain {
	c.middlewares = append(c.middlewares, mw)
	return c
}

func (c *Chain) Then(handler handler.IHandler) handler.IHandler {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	return handler
}
