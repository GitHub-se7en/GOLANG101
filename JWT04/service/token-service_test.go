package service

import (
	"GOLANG101/JWT04/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestDecode(t *testing.T) {

}
func TestEncode(t *testing.T) {
	todo := &models.Todo{
		ID:          2,
		Title:       "token",
		Description: "my token",
	}
	token, err := Encode(todo)
	if err != nil {
		t.Fatal("加密token出错", err)
	}
	t.Log(token)
}
func TestAAA(t *testing.T) {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}
