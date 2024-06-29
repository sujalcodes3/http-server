package main

import (
	"bytes"
	"fmt"
	"net"

	"github.com/sujalcodes3/http-server/handler"
	"github.com/sujalcodes3/http-server/parser"
	"github.com/sujalcodes3/http-server/util"
)

type User struct {
	Email string
	Name  string
	Age   uint
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

        request := handler.NewRequest(parserResult, conn)
        request.HandleRequest()
    }
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	util.HandleError(err, "creating listener")

	fmt.Println("Server listening at 127.0.0.1:8080")

	for {
		conn, err := listener.Accept()
		util.HandleError(err, "accepting connection")

		go handleConnection(conn)
	}
}
