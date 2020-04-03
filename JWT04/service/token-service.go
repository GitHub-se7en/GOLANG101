package service

//
//import (
//	"GOLANG101/JWT04/models"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//)
//
//var mykeyy = []byte("mySuperSecretKeyLol")
//
//type CustomClaims struct {
//	Foo string
//	jwt.StandardClaims
//}
//
//func Decode(token string) (*CustomClaims, error) {
//
//	//解析token
//	tokenType, err := jwt.ParseWithClaims(string(mykeyy), &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
//		return mykeyy, nil
//	})
//
//	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
//		return claims, nil
//	} else {
//		return nil, err
//	}
//
//}
//func Encode(todo *models.Todo) (string, error) {
//	mySigningKey := []byte("AllYourBase")
//	type MyCustomClaims struct {
//		Foo string `json:"foo"`
//		jwt.StandardClaims
//	}
//	// Create the Claims
//	claims := MyCustomClaims{
//		"bar",
//		jwt.StandardClaims{
//			ExpiresAt: 15000,
//			Issuer:    "test",
//		},
//	}
//
//	//
//	//claims := CustomClaims{
//	//	Foo: "bar",
//	//	StandardClaims: jwt.StandardClaims{
//	//		ExpiresAt: 15000,
//	//		Issuer:    "test",
//	//	},
//	//}
//	//token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
//	//ss, err := token.SignedString(mySigningKey)
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	ss, err := token.SignedString(mySigningKey)
//	fmt.Println(ss, err)
//	return "", nil
//}
