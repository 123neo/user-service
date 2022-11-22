package handlers

import "user-service/models"

// CreateUserRequest holds the request parameters for the Create User method.
type CreateUserRequest struct {
	user models.User
}

// CreateUserResponse holds the response values for the Create User method.
type CreateUserResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}
