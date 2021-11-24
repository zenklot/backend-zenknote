package helper

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Config(key string) string {
	// err := godotenv.Load(".env")
	// ErrPanic(err)
	// fmt.Println("in production", err)
	return os.Getenv(key)
}

func ErrPanic(err error) error {
	if err != nil {
		log.Fatal("Error : ", err)
		return err
	}
	return nil
}

func StringToSlice(str string) []string {
	var strs []string
	strs = append(strs, str)
	return strs
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
