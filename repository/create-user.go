package repository

import (
	"database/sql"
	"log"
	"user-service/models"
)

type Repo struct {
	Db  *sql.DB
	log *log.Logger
}

func NewRepo(Db *sql.DB, log *log.Logger) Repository {
	return &Repo{
		Db:  Db,
		log: log,
	}
}

type Repository interface {
	CreateUser(user models.User) error
}

func (repo *Repo) CreateUser(user models.User) error {
	sqlStatement := `
	INSERT INTO users (user_id, first_name, last_name, email, contact)
	VALUES ($1, $2, $3, $4, $5)`

	repo.log.Println("User: ", user)

	_, err := repo.Db.Exec(sqlStatement, user.ID, user.FirstName, user.LastName, user.Email, user.Contact)
	if err != nil {
		return err
	}

	return nil
}
