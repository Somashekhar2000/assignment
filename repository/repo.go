package repository

import (
	model "assignment/models"
	"errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

type Repository interface {
	CreateUser(userData *model.User) (model.User, error)
}

func NewRepo(db *gorm.DB) (Repository, error) {
	if db == nil {
		return nil, errors.New("db is empty")
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateUser(userData *model.User) (model.User, error) {
	userDetail := r.db.Create(&userData)
	if userDetail.Error != nil {
		log.Info().Err(userDetail.Error).Send()
		return model.User{}, userDetail.Error
	}
	return *userData, nil
}
