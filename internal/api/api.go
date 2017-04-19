package api

import (
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/http"
	"../repository"
	"../models"
)


func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := repository.GetUsers()
	usersjson, _ := json.Marshal(users)
	w.Write(usersjson)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	p := models.UserParams{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p)
	user := models.User{
		Username:     p.Username,
		MoneyBalance: p.MoneyBalance,
	}
	repository.AddUser(user)

	w.WriteHeader(http.StatusCreated)
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", listUsersHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	return r
}
