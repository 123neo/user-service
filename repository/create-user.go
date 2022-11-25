package repository

import (
	"database/sql"
	"user-service/models"
)

type Repo struct {
	Db *sql.DB
}

func NewRepo(Db *sql.DB) *Repo {
	return &Repo{
		Db: Db,
	}
}

type Repository interface {
	CreateUser(user models.User) error
}

func (repo *Repo) CreateUser(user models.User) {

}
