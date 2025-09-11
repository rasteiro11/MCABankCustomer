package security

import "golang.org/x/crypto/bcrypt"

type bcryptHasher struct{}

var _ PasswordHasher = (*bcryptHasher)(nil)

func NewPasswordHasher() PasswordHasher {
	return &bcryptHasher{}
}

func (h *bcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *bcryptHasher) Verify(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
