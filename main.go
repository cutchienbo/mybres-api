package main

import (
	"net"
	"net/http"
	"os"
	"sample/app"
)

func main() {
	route := app.InitRoot()

	http.Handle("/", route)

	port := os.Getenv("PORT")

	ln, _ := net.Listen("tcp", "localhost:" + port)
	
	_ = http.Serve(ln, route)
}
