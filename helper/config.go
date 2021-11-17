package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	ErrPanic(err)
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
