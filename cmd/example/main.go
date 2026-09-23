package main

import (
	"fmt"
	"log"

	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
	"github.com/moeinht/http-protocol/router"
	"github.com/moeinht/http-protocol/server"
)

func helloword(req request.Request) (response.Response, error) {
	res := response.NewResponse(200)

	if err := res.H(response.HType{
		"message": "hello",
		"status":  "ok",
	}); err != nil {
		return response.Response{}, err
	}

	return *res, nil
}

func users(req request.Request) (response.Response, error) {
	res := response.NewResponse(200)

	if err := res.SetBody([]byte(fmt.Sprintf("userId:%v and postId:%v", req.Params["userId"], req.Params["postId"]))); err != nil {
		return response.Response{}, err
	}

	return *res, nil
}

func createUser(req request.Request) (response.Response, error) {
	res := response.NewResponse(201)

	if err := res.SetBody([]byte("User Created")); err != nil {
		return response.Response{}, err
	}

	return *res, nil
}

func main() {
	r := router.NewRouter()

	r.GET("/hello", helloword)
	r.GET("/users/:userId/posts/:postId", users)
	r.POST("/users", createUser)

	listener, err := server.Listener(8000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server running on :8000")

	if err := server.Server(listener, r); err != nil {
		log.Fatal(err)
	}
}
