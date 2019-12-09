package services

import (
	"github.com/RajatWithGolang/Microservice/domain"
	"github.com/RajatWithGolang/Microservice/utils"
)

type userService struct{}

var UserService userService

func (us *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDAO.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
