package handlers

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/models"
	"user-service/services"
)

func CreateHandlerFunc(app *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		err := decodeCreateUserRequest(w, r, &user)

		var payload CreateUserResponse

		if err != nil {
			app.log.Println("Error in decoding JSON : ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		service := services.NewService(app.repo, user)
		repsonse, err := service.CreateUser()

		payload = CreateUserResponse{
			Data: user,
		}

		err = encodeResponse(w, http.StatusAccepted, payload)

		if err != nil {
			log.Println("Some error occured: ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}
	}
}
