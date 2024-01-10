package service

import (
	"assignment/repository"
	"errors"
)

type Service struct {
	Repo repository.Repository
}
type Services interface {
}

func NewService(r repository.Repository) (Services, error) {
	if r == nil {
		return nil, errors.New("service nil")
	}
	return &Service{}, nil
}
