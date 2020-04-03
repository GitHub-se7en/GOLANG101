package service

import (
	"GOLANG101/JWT04/models"
	"github.com/dgrijalva/jwt-go"
)

//第一个作废了，太难了
//找了一圈，发现算法用错了，导致key一直是不生效的

var key = []byte("MyKey")

type CustomClaimsTwo struct {
	Todo *models.Todo
	jwt.StandardClaims
}

func Decode(token string) (*CustomClaimsTwo, error) {
	tokenType, e := jwt.ParseWithClaims(token, &CustomClaimsTwo{}, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})
	if claimsTwo, ok := tokenType.Claims.(*CustomClaimsTwo); ok && tokenType.Valid {
		return claimsTwo, nil
	} else {
		return nil, e
	}
}

func Encode(todo *models.Todo) (string, error) {
	two := CustomClaimsTwo{
		Todo:           todo,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt:1520000090345,这个过期在不要随便用，应该是当前的毫秒数，我是debug之后才发现这个错误的
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, two)
	return token.SignedString(key)
}
