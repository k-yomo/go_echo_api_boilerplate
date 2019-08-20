package jwt_generator

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/k-yomo/go_echo_api_boilerplate/config"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"
	"github.com/pkg/errors"
	"time"
)

// GenerateJwt generates JWT token
func GenerateJwt(id uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = clock.Now().Unix()
	claims["exp"] = clock.Now().Add(time.Hour * 6).Unix()

	encodedToken, err := token.SignedString([]byte(config.JwtSigningKey))
	if err != nil {
		return "", errors.Wrap(err, "generate jwt_generator token failed")
	}
	return encodedToken, nil
}
