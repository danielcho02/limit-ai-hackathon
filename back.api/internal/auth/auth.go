package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// returns the bcrypt hashed password
//
// password는 72byte를 초과하면 에러 발생(bcrypt 내부 로직)
func EncryptPassword(password string) (string, error) {
	// GeneratedFromPassword 내부 NewFromPassword에서 salt 생성
	pwHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwHash), nil
}

// ComparePassword: 평문 비번과 DB의 해시를 비교
func ComparePassword(hashedPassword string, plainPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}