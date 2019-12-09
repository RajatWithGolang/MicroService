package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RajatWithGolang/Microservice/services"
	"github.com/RajatWithGolang/Microservice/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	//getting the userId from the request
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonError, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write([]byte(jsonError))
		return
	}
	user, apiError2 := services.UserService.GetUser(userID)
	if apiError2 != nil {
		jsonError2, _ := json.Marshal(apiError2)
		res.WriteHeader(apiError2.StatusCode)
		res.Write([]byte(jsonError2))
		return
	}
	//return user to the client
	jsonData, err := json.Marshal(user)
	res.Write(jsonData)
}
