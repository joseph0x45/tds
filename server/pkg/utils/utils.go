package utils

import (
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func IsNotEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return !regex.MatchString(email)
}

func IsEmpty(s *string) bool {
	return *s == ""
}

func IsOneOf(x string, ys ...string) bool {
	ok := false
	for _, y := range ys {
		if y == x {
			ok = true
			break
		}
	}
	return ok
}

func Hash(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Error while hashing password: %w", err)
	}
	return string(hash), nil
}

func HashMatches(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateRandomDigit() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
