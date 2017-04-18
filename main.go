package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)

type User struct {
	Id           uint32 `json:"id"`
	Username     string `json:"username"`
	MoneyBalance uint32 `json:"balance"`
}

var userStore = []User{}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := json.Marshal(userStore)
	w.Write(users)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := json.Marshal(userStore)
	w.Write(users)
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", listUsersHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	return r
}
func main(){
	var userBob = User{0, "Bob",100}
	userStore = append(userStore,userBob)
	fmt.Println("Server starting")
	http.ListenAndServe(":8000", Handlers())
}