package router

import (
	"github.com/moeinht/http-protocol/handler"
	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
)

type Route struct {
	Method  string
	Path    string
	handler handler.IHandler
}

type Router struct {
	Routes []Route
}

func (r *Router) GET(path string, routeHandler handler.IHandler) {
	route := Route{
		Method:  "GET",
		Path:    path,
		handler: routeHandler,
	}

	r.Routes = append(r.Routes, route)
}

func (r *Router) POST(path string, routeHandler handler.IHandler) {
	route := Route{
		Method:  "POST",
		Path:    path,
		handler: routeHandler,
	}

	r.Routes = append(r.Routes, route)
}

func (r *Router) Handler(req request.Request) (response.Response, error) {
	method := req.RequestLine.Method
	path := req.RequestLine.Path

	for _, route := range r.Routes {
		if route.Method == method &&
			route.Path == path {

			return route.handler(req)
		}
	}

	res := response.NewResponse(404)

	return *res, nil
}

func NewRouter() *Router {
	return &Router{}
}
