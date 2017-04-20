package repository

import (
	"github.com/bertt/golang_webapi_example/models"
)

var users = []models.User{}
var userIDCounter uint32

// GetUsers get the users from repos
func (r *UserRepository) GetUsers() ([]models.User) {
	return users
}

// AddUser adds a user to repos
func (r *UserRepository) AddUser(user models.User){
	user.Id = userIDCounter
	userIDCounter++
	users = append(users, user)
}

type UserRepository struct{}