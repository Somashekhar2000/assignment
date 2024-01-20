package service

import (
	"assignment/graph/model"
	"assignment/logger"
	mymodel "assignment/models"
	"assignment/validation"
	"errors"
)

func (s Service) CreateUser(userDetail model.User) (*model.User, error) {
	ok := validation.IsValidMobileNumber(userDetail.Mobile)
	if !ok {
		logger.ErrorLogger.Println("Inavlid mobile number for user user id : ", userDetail.ID)
		return nil, errors.New("invalid mobile number")
	}
	ok = validation.AgeValidation(userDetail.Age)
	if !ok {
		return nil, errors.New("invalid user age")
	}
	usersDetails := mymodel.User{
		FirstName:        userDetail.FirstName,
		LastName:         userDetail.LastName,
		Age:              userDetail.Age,
		Gender:           userDetail.Gender,
		Dob:              userDetail.Dob,
		Mobile:           userDetail.Mobile,
		Email:            userDetail.Email,
		PermanentAddress: userDetail.PermanentAddress,
	}
	user, err := s.Repo.CreateUser(usersDetails)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:               int(user.ID),
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Age:              user.Age,
		Gender:           user.Gender,
		Dob:              user.Dob,
		Mobile:           user.Mobile,
		Email:            user.Email,
		PermanentAddress: user.PermanentAddress,
	}, nil

}

func (s *Service) GetAllUser() ([]*model.User, error) {
	AllUser, err := s.Repo.FetchUser()
	if err != nil {
		return nil, err
	}
	var users []*model.User

	for _, value := range AllUser {
		user := model.User{
			ID:               int(value.ID),
			FirstName:        value.FirstName,
			LastName:         value.LastName,
			Age:              value.Age,
			Gender:           value.Gender,
			Dob:              value.Dob,
			Mobile:           value.Mobile,
			Email:            value.Email,
			PermanentAddress: value.PermanentAddress,
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *Service) GetUserByID(id int) (*model.User, error) {
	// intValue, _ := strconv.Atoi(id)
	user, err := s.Repo.FetchUserByID(id)
	if err != nil {
		return nil, err
	}
	userDetail := model.User{
		ID:               id,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Age:              user.Age,
		Gender:           user.Gender,
		Dob:              user.Dob,
		Mobile:           user.Mobile,
		Email:            user.Email,
		PermanentAddress: user.PermanentAddress,
	}
	return &userDetail, nil
}

func (s *Service) UpdateUser(userDetail model.UpdateUser) (*model.User, error) {
	ok := validation.IsValidMobileNumber(userDetail.Mobile)
	if !ok {
		return nil, errors.New("invalid mobile number")
	}
	userData, err := s.Repo.FetchUserByID(userDetail.ID)
	if err != nil {
		return nil, err
	}

	userData.Mobile = userDetail.Mobile
	userData.Email = userDetail.Email

	err = s.Repo.UpdateUser(userData)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:               int(userData.ID),
		FirstName:        userData.FirstName,
		LastName:         userData.LastName,
		Age:              userData.Age,
		Gender:           userData.Gender,
		Dob:              userData.Dob,
		Mobile:           userData.Mobile,
		Email:            userData.Email,
		PermanentAddress: userData.PermanentAddress,
	}, nil
}
