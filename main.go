package main

import (
	"fmt"
	"net/http"

	"github.com/bertt/golang_webapi_example/api"
)

func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":8000", api.Handlers())
}
