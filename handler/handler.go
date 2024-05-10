package handler

import (
	"net"

	"github.com/sujalcodes3/http-server/parser"
	"github.com/sujalcodes3/http-server/util"
)

type Request struct {
    requestData *parser.ParseResult
    conn net.Conn
}

type Response struct {}

func (r * Request) HandleRequest() {
    conn := r.conn          

    switch r.requestData.GetMethod() {
    case parser.GET:
    case parser.PUT:
    case parser.POST:
    case parser.DELETE:
    case parser.PATCH:
    }

    defer conn.Close()

    _, err := conn.Write([]byte("hello World"))
    
    util.HandleError(err, "writing hello world to the connection")
}

func NewRequest(requestData * parser.ParseResult, conn net.Conn) *Request {
    request := &Request{requestData: requestData, conn: conn}

    return request
}
