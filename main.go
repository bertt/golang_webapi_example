package main

import (
	"fmt"
	"net/http"

	"github.com/bertt/golang_webapi_example/api"
	"database/sql"
)

func main() {
	// test with sqlite
	sql.Open("sqlite3", "~/my.db")
	//db.Close()
	fmt.Println("Server starting")
	http.ListenAndServe(":8000", api.Handlers())
}
