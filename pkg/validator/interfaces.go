package validator

type (
	EmailValidator interface {
		IsValid(email string) bool
	}
)
