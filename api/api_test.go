package api

import (
	"testing"
	"strings"
	"net/http/httptest"
	"io"
	"fmt"
	"net/http"
)
var (
	server   *httptest.Server
	reader   io.Reader
	usersUrl string
)

func init() {
	server = httptest.NewServer(Handlers())

	usersUrl = fmt.Sprintf("%s/users", server.URL)
}


func TestListUsers(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestCreateUser(t *testing.T) {
	userJson := `{"username": "dennis", "balance": 200}`

	reader = strings.NewReader(userJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
