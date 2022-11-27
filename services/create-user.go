package services

import (
	"context"
	"log"
	"user-service/models"
	"user-service/repository"

	"github.com/gofrs/uuid"
)

type service struct {
	repo repository.Repository
	log  *log.Logger
}

func NewService(repo repository.Repository, log *log.Logger) Service {
	return &service{repo: repo, log: log}
}

type Service interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
}

func (s *service) CreateUser(ctx context.Context, user models.User) (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		s.log.Println("Error in generating uuid..")
	}
	id := uuid.String()
	user.ID = id

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return "", err
	}
	return id, nil

}
