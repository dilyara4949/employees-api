package internal

import (
	"os"
)

type Env struct {
	JWTTokenSecret string
}

func NewEnv() *Env {
	env := Env{}

	env.JWTTokenSecret = os.Getenv("JWTTokenSecret")
	return &env
}
