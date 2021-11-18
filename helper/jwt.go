package helper

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zenklot/backend-zenknote/model/web"
)

func ValidateJWT(inputToken string, key string) (*web.TokenClaims, error) {
	var claims *web.TokenClaims
	var err error
	defer func() {
		if r := recover(); r != nil {
			claims = nil
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()

	asli := strings.Contains(inputToken, ".")
	if !asli {
		decToken, err := jwt.DecodeSegment(inputToken)
		if err != nil {
			return nil, err
		}
		inputToken = string(decToken)
	}

	token, err := jwt.ParseWithClaims(inputToken, &web.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	claims, ok := token.Claims.(*web.TokenClaims)
	if !ok && !token.Valid {
		return nil, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, err
		} else {
			return nil, err
		}
	}
	if claims == nil {
		return nil, err
	}
	return claims, nil
}
