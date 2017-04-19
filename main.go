package main

import (
	"net/http"
	"fmt"
	"github.com/bertt/golang_webapi_example/api"
)


func main(){
	var userBob = api.User{0, "Bob",100}
	api.UserStore = append(api.UserStore,userBob)
	fmt.Println("Server starting")
	http.ListenAndServe(":8000", api.Handlers())
}