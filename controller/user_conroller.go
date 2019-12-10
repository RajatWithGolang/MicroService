package controller

import (
	"net/http"
	"strconv"

	"github.com/RajatWithGolang/Microservice/services"
	"github.com/RajatWithGolang/Microservice/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	//getting the userId from the
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	user, apiError2 := services.UserService.GetUser(userID)
	if apiError2 != nil {
		c.JSON(apiError2.StatusCode, apiError2)
		return
	}
	//return user to the client
	c.JSON(http.StatusOK, user)
}
