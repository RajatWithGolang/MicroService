package domain

import (
	"fmt"
	"net/http"

	"github.com/RajatWithGolang/Microservice/utils"
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDAO struct{}

var (
	users = map[int64]*User{
		123: {Id: 123, Name: "Rajat", Email: "rajat@test.com"},
	}
	UserDAO userDaoInterface
)

func init() {
	UserDAO = &userDAO{}
}

func (ud *userDAO) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with Id %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
