package validation

import (
	"strconv"
)

func IsValidMobileNumber(number string) bool {
	_, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	if len(number) != 10 {
		return false
	}
	return true
}

func AgeValidation(age string) bool {
	userAge, err := strconv.Atoi(age)
	if err != nil {
		return false
	}
	if userAge < 1 && userAge > 100 {
		return false
	}
	return true
}
