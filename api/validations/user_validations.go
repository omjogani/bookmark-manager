package validations

import (
	"regexp"

	model "github.com/omjogani/bookmark-manager/models"
)

func IsUserValid(user model.User) string {
	emailRegEx := regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-]+)(\.[a-zA-Z]{2,5}){1,2}$`)
	passwordRegEx := regexp.MustCompile(`^[a-zA-Z0-9]{6,}$`)

	// Email Validation
	if !emailRegEx.MatchString(user.Email) {
		return "Invalid Email!"
	}

	// Password Validation (MinLength = 6, Contains alpha-numeric)
	if !passwordRegEx.MatchString(user.Password) {
		return "Invalid Password Format, MinLength = 6 & must Contain Alpha-Numeric Value"
	}

	// Username Validation
	if user.UserName == "" {
		return "Username must not be Empty!"
	} else if len(user.UserName) < 2 {
		return "Invalid User Name!"
	}

	return "OK"
}
