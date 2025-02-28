package api

import (
	"fmt"
	"net/http"

	"github.com/fritzyl/receipt-processor-challenge/api/routes"
)

func Serve(port string) {
	fmt.Println("Listening on Port", port)
	http.ListenAndServe(port, buildServer())
}

func buildServer() *http.ServeMux {
	var server *http.ServeMux = http.NewServeMux()
	routes.Register(server)
	return server
}
