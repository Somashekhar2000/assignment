package service

import (
	"assignment/graph/model"
	"assignment/repository"
	"errors"
)

type Service struct {
	Repo repository.Repository
}
type Services interface {
	CreateUser(userDetail model.User) (*model.User, error)
	GetAllUser() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
	UpdateUser(userDetail model.UpdateUser) (*model.User, error)
}

func NewService(r repository.Repository) (Services, error) {
	if r == nil {
		return nil, errors.New("service nil")
	}
	return &Service{
		Repo: r,
	}, nil
}
