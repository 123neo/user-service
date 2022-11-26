package models

type User struct {
	ID        string `json: "userId,omitEmpty"`
	FirstName string `json: "firstName,omitEmpty"`
	LastName  string `json: "lastName,omitEmpty"`
	Email     string `json: "email,omitEmpty"`
	Contact   string `json: "contact,omitEmpty"`
}
