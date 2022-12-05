package handlers

import (
	"net/http"
	"user-service/config"
	"user-service/models"
	"user-service/services"
)

func CreateHandlerFunc(app *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		// decoding the request payload

		var user models.User

		err := decodeCreateUserRequest(w, r, &user)

		requestPaylod := CreateUserRequest{
			user: user,
		}

		if err != nil {
			app.Log.Println("Error in decoding JSON : ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		// calling the required service

		service := services.NewService(app.Repo, app.Log)
		repsonse, err := service.CreateUser(ctx, requestPaylod.user)

		if err != nil {
			app.Log.Println("Some error occured: ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		// sending the response back to the client

		payload := CreateUserResponse{
			Data: repsonse,
		}

		err = encodeResponse(w, http.StatusAccepted, payload)

		if err != nil {
			app.Log.Println("Some error occured: ", err)
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}
	}
}
