package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const wrongTestUserId = 0
const correctTestUserId = 123

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, appErr := UserDao.GetUser(wrongTestUserId)

	// Validation
	assert.Nil(t, user, fmt.Sprintf("we were not expecting user with id: %v", wrongTestUserId))
	assert.NotNil(t, appErr, fmt.Sprintf("we were expecting an error when user id is %v", wrongTestUserId))
	assert.EqualValues(t, http.StatusNotFound, appErr.StatusCode)
	assert.EqualValues(t, "not_found", appErr.Code)
	assert.EqualValues(t, fmt.Sprintf("user with id: '%v' was not found", wrongTestUserId), appErr.Message)

	// solution without assert
	//if user != nil {
	//	t.Error(fmt.Sprintf("we were not expecting user with id: %v", wrongTestUserId))
	//}
	//
	//if appErr == nil {
	//	t.Error(fmt.Sprintf("we were expecting an error when user id is %v", wrongTestUserId))
	//}
	//
	//if appErr.StatusCode != http.StatusNotFound {
	//	t.Error("we were expecting 404 when user is not found")
	//}
}

func TestGetUserNoError(t *testing.T) {
	user, appErr := UserDao.GetUser(correctTestUserId)

	assert.Nil(t, appErr)
	assert.NotNil(t, user)
	assert.EqualValues(t, correctTestUserId, user.Id)
	assert.EqualValues(t, "Fede", user.FirstName)
	assert.EqualValues(t, "Leon", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}
