package router

import (
	"github.com/moeinht/http-protocol/handler"
	"github.com/moeinht/http-protocol/params"
	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
)

type Route struct {
	Method  string
	Path    string
	handler handler.IHandler
	Params  params.Params
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

func (r *Router) PUT(path string, routeHandler handler.IHandler) {
	route := Route{
		Method:  "PUT",
		Path:    path,
		handler: routeHandler,
	}

	r.Routes = append(r.Routes, route)
}

func (r *Router) PATCH(path string, routeHandler handler.IHandler) {
	route := Route{
		Method:  "PATCH",
		Path:    path,
		handler: routeHandler,
	}

	r.Routes = append(r.Routes, route)
}
func (r *Router) DELETE(path string, routeHandler handler.IHandler) {
	route := Route{
		Method:  "DELETE",
		Path:    path,
		handler: routeHandler,
	}

	r.Routes = append(r.Routes, route)
}

func (r *Router) Handler(req request.Request) (response.Response, error) {
	method := req.RequestLine.Method
	path := req.RequestLine.Path

	foundPath := false
	for _, route := range r.Routes {
		if isMatch, params := params.Match(route.Path, path); isMatch {
			foundPath = true
			if route.Method == method {
				req.Params = params
				return route.handler(req)
			}
		}
	}
	if foundPath {
		res := response.NewResponse(405)
		return *res, nil
	}

	res := response.NewResponse(404)

	return *res, nil
}

func NewRouter() *Router {
	return &Router{}
}
