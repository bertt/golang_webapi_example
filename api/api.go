package api

import (
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
Id           uint32 `json:"id"`
Username     string `json:"username"`
MoneyBalance uint32 `json:"balance"`
}

type UserParams struct {
	Username     string `json:"username"`
	MoneyBalance uint32 `json:"balance"`
}

var UserStore = []User{}
var userIdCounter uint32 = 0

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := json.Marshal(UserStore)
	w.Write(users)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	p := UserParams{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p)
	u := User{
		Id:           userIdCounter,
		Username:     p.Username,
		MoneyBalance: p.MoneyBalance,
	}

	UserStore = append(UserStore, u)

	userIdCounter += 1
	w.WriteHeader(http.StatusCreated)
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", listUsersHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	return r
}
