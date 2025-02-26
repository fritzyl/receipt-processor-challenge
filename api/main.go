package main

import (
	"net/http"

	"github.com/fritzyl/receipt-processor-challenge/routes"
)

func main() {
	http.ListenAndServe(":8080", buildServer())
}

func buildServer() *http.ServeMux {
	var server *http.ServeMux = http.NewServeMux()
	routes.Register(server)
	return server
}
