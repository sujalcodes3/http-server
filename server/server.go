package server

import (
	"bytes"
	"fmt"
	"net"

	"github.com/sujalcodes3/http-server/context"
	"github.com/sujalcodes3/http-server/parser"
	"github.com/sujalcodes3/http-server/router"
	"github.com/sujalcodes3/http-server/util"
)

type Server struct {
	Port   int16
	Router *router.Router
}

func New(port int16) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	util.HandleError(err, "creating listener")

	for {
		conn, err := listener.Accept()
		util.HandleError(err, "accepting connection")

		go handleConnection(conn)
	}
}

func (s *Server) RegisterRouter() {
	s.Router = router.New()
}

func (s *Server) NewRoute(method, path string, handler *router.HandlerFunc) {
	s.Router.Register(method, path, handler)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client Connect : %v\n", clientAddr)

	var message []byte = make([]byte, 1024)

	for {
		len, err := conn.Read(message)
		util.HandleError(err, fmt.Sprintf("reading message from %s", clientAddr))
		if bytes.Equal([]byte("exit"), message) {
			break
		}

		p := parser.New(message)
		parserResult := p.ParseMessage()

		fmt.Printf("[%s] [%d] : %s\n", clientAddr, len, message)
		parserResult.PrintParseResult()

		context := context.New(&conn)

	}
}
