package handlers

import (
	"log"
	"net/http"
	"user-service/models"
)

type jsonResponse struct {
	Error   bool        `json: "error,omitEmpty"`
	Message string      `json: "message,omitEmpty"`
	Data    interface{} `json: "data,omitEmpty"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := decodeCreateUserRequest(w, r, &user)

	var payload jsonResponse

	if err != nil {
		log.Println("Error in decoding JSON : ", err)
		payload = jsonResponse{
			Error:   true,
			Message: err.Error(),
		}
	} else {
		payload = jsonResponse{
			Data: user,
		}
	}

	// fmt.Fprintf(w, "User: %+v", user)

	err = encodeResponse(w, http.StatusAccepted, payload)

	if err != nil {
		log.Println("Some error occured: ", err)
		payload := jsonResponse{
			Error:   true,
			Message: err.Error(),
		}
		err = encodeResponse(w, http.StatusAccepted, payload)
	}
}
