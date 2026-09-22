package handler

import (
	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
)

type IHandler func(request request.Request) (response.Response, error)

func Handler(request request.Request) (response.Response, error) {
	handlerResponse := response.NewResponse(200)

	return *handlerResponse, nil
}
