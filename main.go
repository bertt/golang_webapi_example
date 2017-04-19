package main

import (
	"net/http"
	"fmt"
	"./api"
)


func main(){
	fmt.Println("Server starting")
	http.ListenAndServe(":8000", api.Handlers())
}