package router

import "github.com/sujalcodes3/http-server/context"

type Router struct {
    Routes map[Route]*HandlerFunc 
}

type Route struct {
    Method string
    Path string
}

type HandlerFunc func(req *context.Request, res *context.Response);

func New() *Router {
    routes := make(map[Route]*HandlerFunc)
    return &Router {
        Routes: routes,
    }
}

func (r * Router) Register(method, path string, handler *HandlerFunc) {
    route := Route {
        Method: method,
        Path: path,
    }

    r.Routes[route] = handler
}
