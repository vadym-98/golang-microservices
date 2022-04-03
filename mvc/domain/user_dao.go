package domain

import (
	"fmt"
	"github.com/vadym-98/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Fede",
			LastName:  "Leon",
			Email:     "myemail@gmail.com",
		},
	}
	UserDao usersDaoInterface
)

func init() {
	log.Println("init from user_dao file")
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing the database")

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id: '%v' was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
