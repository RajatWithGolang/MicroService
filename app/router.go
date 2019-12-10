package app

import (
	"net/http"

	"github.com/RajatWithGolang/Microservice/controller"
)

func mapURL() {
	http.HandleFunc("/users/:user_id", controller.GetUser)
}
