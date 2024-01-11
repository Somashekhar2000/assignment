package mymodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName        string `json:"FirstName" validate:"required"`
	LastName         string `json:"LastName" validate:"required"`
	Age              int    `json:"Age" validate:"required"`
	Gender           string `json:"Gender" validate:"required"`
	Dob              string `json:"DOB" validate:"required"`
	Mobile           int    `json:"Mobile" validate:"required"`
	Email            string `json:"Email" validate:"required"`
	PermanentAddress string `json:"PermanentAddress" validate:"required"`
}
