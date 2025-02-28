package main

import (
	"github.com/fritzyl/receipt-processor-challenge/api"
)

func main() {
	api.Serve(":8080")
}
