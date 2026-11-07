package jwt

import (
	"auth-two/internal/config"
	"fmt"
)

var session_key = config.LoadEnv().SESSION_SECRET

type Response struct {
	Token string `json:"token"`
	Status string `json:"status"`
}

func CreateJwtToken() {
	fmt.Println("create token ->", session_key)
}

func VerifyToken() {
	fmt.Println("verify token ->", session_key)
}