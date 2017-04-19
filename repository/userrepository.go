package repository

import(
	"../models"
)

var users = []models.User{}
var userIdCounter uint32 = 0

func GetUsers() []models.User {
	return users
}

func AddUser(user models.User){
	user.Id = userIdCounter
	userIdCounter++
	users = append(users, user)
}


