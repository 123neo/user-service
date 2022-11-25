package config

import (
	"database/sql"
	"log"
	"user-service/repository"
)

type Config struct {
	Db   *sql.DB
	Repo *repository.Repo
	Log  *log.Logger
}

func NewConfig(conn *sql.DB, repo *repository.Repo, log *log.Logger) *Config {
	return &Config{
		Db:   conn,
		Repo: repo,
		Log:  log,
	}
}
