package jwt

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	issuerString = "jwt"
	expDuration  = 60
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

//Parse -
func Parse(ctx context.Context, token string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Parse")
	defer span.Finish()

	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretString, nil
	})

	if t != nil {
		if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
			return claims.UserID, nil
		}
	}
	return "", err
}

//IsExpired - error category
func IsExpired(err error) bool {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			return true
		}
	}
	return false
}

//IsMalformed - error category
func IsMalformed(err error) bool {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorMalformed) != 0 {
			return true
		}
	}
	return false
}
