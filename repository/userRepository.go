package repository

import (
	mymodel "assignment/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(userData mymodel.User) (*mymodel.User, error) {
	userDetail := r.db.Create(&userData)
	if userDetail.Error != nil {
		log.Info().Err(userDetail.Error).Send()
		return nil, userDetail.Error
	}
	return &userData, nil
}

func (r *Repo) FetchUser() ([]mymodel.User, error) {
	var user []mymodel.User
	err := r.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repo) FetchUserByID(id int) (mymodel.User, error) {
	var user mymodel.User
	tx := r.db.Where("id = ?", id)
	err := tx.Find(&user).Error
	if err != nil {
		return mymodel.User{}, nil
	}
	return user, nil
}

func (r *Repo) UpdateUser(userData mymodel.User) error {
	err := r.db.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}
