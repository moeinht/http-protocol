package main

import (
	"fmt"
	"log"

	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
	"github.com/moeinht/http-protocol/router"
	"github.com/moeinht/http-protocol/server"
)

type HelloHandler struct{}

func (HelloHandler) Handler(req request.Request) (response.Response, error) {
	res := response.NewResponse(200)

	if err := res.SetBody([]byte("Hello World")); err != nil {
		return response.Response{}, err
	}

	return *res, nil
}

type UserHandler struct{}

func (UserHandler) Handler(req request.Request) (response.Response, error) {
	res := response.NewResponse(200)

	if err := res.SetBody([]byte("Users")); err != nil {
		return response.Response{}, err
	}

	return *res, nil
}

type CreateUserHandler struct{}

func (CreateUserHandler) Handler(req request.Request) (response.Response, error) {
	res := response.NewResponse(201)

	if err := res.SetBody([]byte("User Created")); err != nil {
		return response.Response{}, err
	}

	return *res, nil
}

func main() {
	r := router.NewRouter()

	r.GET("/hello", HelloHandler{})
	r.GET("/users", UserHandler{})
	r.POST("/users", CreateUserHandler{})

	listener, err := server.Listener(8000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server running on :8000")

	if err := server.Server(listener, r); err != nil {
		log.Fatal(err)
	}
}
