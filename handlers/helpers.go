package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"user-service/models"
)

func decodeCreateUserRequest(w http.ResponseWriter, r *http.Request, user *models.User) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(user)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}
