package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	issuerString = "jwt"
	expDuration  = 3600
)

var secretString = []byte("jwt-secret")

// UserClaims is the type contain userID and standardClaims
type UserClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

// New - create a token
func New(userID string) (string, error) {
	tokenExp := time.Now().Unix() + int64(expDuration)
	claims := UserClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    issuerString,
			ExpiresAt: tokenExp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretString)
	return ss, err
}
