package main

import "github.com/sujalcodes3/http-server/server"

func main() {
    app := server.New(8080)
    app.RegisterRouter() 
    

    app.Start()
}
