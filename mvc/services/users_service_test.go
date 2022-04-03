package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/vadym-98/golang-microservices/mvc/domain"
	"github.com/vadym-98/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"testing"
)

var (
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func init() {
	log.Println("init from test file")
	domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	errorMessage := "user with id: '0' was not found"
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    errorMessage,
			StatusCode: http.StatusNotFound,
		}
	}

	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, errorMessage, err.Message)
}
func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id: 123}, nil
	}

	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
