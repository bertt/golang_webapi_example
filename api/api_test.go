package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	// "github.com/stretchr/testify/assert"
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

func RouterTest(t *testing.T){
	//var h = Handlers()
	//assert.True(t, h.!=nil)
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
