package context

import (
	"net"
)

type Request struct {
}

type Response struct {
	Connection *net.Conn
}

type Context struct {
	Request  *Request
	Response *Response
}

func New(conn *net.Conn) *Context {
	res := &Response{
		Connection: conn,
	}
	req := &Request{}

	return &Context{
		Request:  req,
		Response: res,
	}
}
