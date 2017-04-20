package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bertt/golang_webapi_example/repository"

	"github.com/bertt/golang_webapi_example/models"
	"github.com/gorilla/mux"
	"fmt"
)

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	var repos = repository.UserRepository{}
	users := repos.GetUsers()
	usersjson, _ := json.Marshal(users)
	w.Write(usersjson)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create user")
	p := models.UserParams{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p)
	user := models.User{
		Username:     p.Username,
		MoneyBalance: p.MoneyBalance,
	}
	var repos = repository.UserRepository{}
	repos.AddUser(user)

	w.WriteHeader(http.StatusCreated)
}

// getXfs() testing the X-Forwarded-For header
func getXfs(w http.ResponseWriter, r *http.Request) {
	externalUri := r.Header.Get("X-Forwarded-For")
	bytes := []byte("xfs: " + externalUri)
	w.Write(bytes)
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", listUsersHandler).Methods("GET")
	r.HandleFunc("/xfs", getXfs).Methods("GET")

	r.HandleFunc("/users", createUserHandler).Methods("POST")
	return r
}
