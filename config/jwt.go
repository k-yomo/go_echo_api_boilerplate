package config

import (
	"os"
)

// JwtSigningKey is a key used for signing jwt
var JwtSigningKey = os.Getenv("JWT_SIGNING_KEY")

func init() {
	if JwtSigningKey == "" {
		panic("load JWT_SIGNING_KEY from env variable failed")
	}
}
