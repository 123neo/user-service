package handlers

import (
	"log"
	"net/http"
	"user-service/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := decodeCreateUserRequest(w, r, &user)

	if err != nil {
		log.Println("Error in decoding JSON : ", err)
	}

	log.Println(user)

}
