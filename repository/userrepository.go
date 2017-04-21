package repository

import (
	"github.com/bertt/golang_webapi_example/models"
	"gopkg.in/mgo.v2"
)

var dbserver = "localhost"

// GetUsers get the users from repos
func (r *UserRepository) GetUsers() ([]models.User) {
	session, err := mgo.Dial(dbserver)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var users []models.User
	c := session.DB("test").C("users")
	c.Find(nil).All(&users)

	return users
}

// AddUser adds a user to repos
func (r *UserRepository) AddUser(user models.User){
	session, err := mgo.Dial(dbserver)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("users")
	err = c.Insert(&user)
}

type UserRepository struct{}

type UserInterface interface {
	AddUser(user models.User)
	GetUsers() ([]models.User)
}