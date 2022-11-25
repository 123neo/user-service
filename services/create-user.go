package services

import (
	"user-service/models"
	"user-service/repository"

	"github.com/gofrs/uuid"
)

type Service struct {
	repo *repository.Repo
	user models.User
}

func NewService(repo *repository.Repo, user models.User) *Service {
	return &Service{repo: repo, user: user}
}

func (s *Service) CreateUser(user models.User) (string, error) {
	uuid := uuid.NewV4()
	id := uuid.String()
	user.ID = id

	if err := s.repo.CreateUser(user); err != nil {
		return "", err
	}
	return id, nil

}
