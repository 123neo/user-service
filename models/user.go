package models

type User struct {
	name    string `json: fullname`
	email   string `json: emailId`
	contact string `json: mobile`
}
