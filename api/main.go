package api

import (
	"net/http"

	"github.com/fritzyl/receipt-processor-challenge/api/routes"
)

func Serve() {
	http.ListenAndServe(":8080", buildServer())
}

func buildServer() *http.ServeMux {
	var server *http.ServeMux = http.NewServeMux()
	routes.Register(server)
	return server
}
