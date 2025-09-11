package validator

import (
	"regexp"
)

type emailValidator struct {
	regex *regexp.Regexp
}

var _ EmailValidator = (*emailValidator)(nil)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func NewEmailValidator() EmailValidator {
	return &emailValidator{regex: emailRegex}
}

func (v *emailValidator) IsValid(email string) bool {
	return v.regex.MatchString(email)
}
