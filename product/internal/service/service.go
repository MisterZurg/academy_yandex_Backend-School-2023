package service

import (
	"yandex-team.ru/bstask/internal/api"
	"yandex-team.ru/bstask/internal/repository"
)

type Repository interface {
	// *repository.Repository
	api.ServerInterface
}

type Service struct {
	*repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos}
}

//func (s *Service) Create
