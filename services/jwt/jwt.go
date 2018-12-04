package jwt

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	opentracing "github.com/opentracing/opentracing-go"
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
func New(ctx context.Context, userID string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewToken")
	defer span.Finish()

	tokenExp := time.Now().Unix() + int64(expDuration)
	claims := UserClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    issuerString,
			ExpiresAt: tokenExp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	span.LogKV("event", "NewWithClaims")
	ss, err := token.SignedString(secretString)
	span.LogKV("event", "SignedString")
	return ss, err
}
