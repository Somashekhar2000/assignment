package repository

import (
	mymodel "assignment/models"
	"errors"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

type Repository interface {
	CreateUser(userData mymodel.User) (*mymodel.User, error)
	FetchUser() ([]mymodel.User, error)
	FetchUserByID(id int) (mymodel.User, error)
	UpdateUser(userData mymodel.User) error
}

func NewRepo(db *gorm.DB) (Repository, error) {
	if db == nil {
		return nil, errors.New("db is empty")
	}
	return &Repo{
		db: db,
	}, nil
}
