package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bertt/golang_webapi_example/repository"

	"github.com/bertt/golang_webapi_example/models"
	"github.com/gorilla/mux"
	"fmt"
	"time"
)

var IndentJSON = true

func StartTimer(name string) func() {
	var t = time.Now()
	fmt.Println("Request started: ", name)
	return func() {
		var elapsed = time.Now().Sub(t)
		fmt.Println("Request duration: ", elapsed)
	}

}


func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	var stop = StartTimer("List users")
	defer stop()
	var repos = repository.UserRepository{}
	users := repos.GetUsers()
	res, _ := JSONMarshal(users)
	_, err := w.Write(res)
	if(err!=nil){
		panic("hola")
	}
}


//JSONMarshal converts the data and converts special characters such as &
func JSONMarshal(data interface{}) ([]byte, error) {
	var b []byte
	var err error

	if IndentJSON {
		b, err = json.MarshalIndent(data, "", "   ")
	} else {
		b, err = json.Marshal(data)
	}
	return b, err
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var stop = StartTimer("Create user")
	defer stop()
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
