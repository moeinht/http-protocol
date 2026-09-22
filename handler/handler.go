package handler

import (
	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
)

type IHandler interface {
	Handler(request request.Request) (response.Response, error)
}

type MyHandler struct{}

func (MyHandler) Handler(request request.Request) (response.Response, error) {
	handlerResponse := response.NewResponse(200)

	return *handlerResponse, nil
}
