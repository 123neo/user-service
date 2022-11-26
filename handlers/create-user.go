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

		if err != nil {
			app.Log.Println("Error in decoding JSON : ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		service := services.NewService(app.Repo, app.Log)
		repsonse, err := service.CreateUser(user)

		if err != nil {
			log.Println("Some error occured: ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		payload := CreateUserResponse{
			Data: repsonse,
		}

		err = encodeResponse(w, http.StatusAccepted, payload)

		if err != nil {
			log.Println("Some error occured: ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}
	}
}
