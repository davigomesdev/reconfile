package providers

import "golang.org/x/crypto/bcrypt"

type HashProvider struct{}

func NewHashProvider() *HashProvider {
	return &HashProvider{}
}

func (p *HashProvider) GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p *HashProvider) CompareHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
