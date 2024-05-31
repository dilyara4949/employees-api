package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	JWTTokenSecret string
}

func NewEnv() (*Env, error) {
	env := Env{}

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}
	env.JWTTokenSecret = os.Getenv("JWTTokenSecret")
	return &env, nil
}
