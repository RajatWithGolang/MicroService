package services

import (
	"github.com/RajatWithGolang/Microservice/domain"
	"github.com/RajatWithGolang/Microservice/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
