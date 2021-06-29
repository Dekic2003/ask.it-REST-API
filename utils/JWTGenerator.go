package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

func JWTGenerator(name string,id int) string {
	signingKey:=[]byte(os.Getenv("TOKEN_KEY"))

	token:=jwt.New(jwt.SigningMethodHS256)

	claims:=token.Claims.(jwt.MapClaims)
	claims["user"]=name
	claims["id"]=id

	tokenString,err :=token.SignedString(signingKey)
	if err!=nil{
		fmt.Println("Error occurred in signing token",err)
	}
	return tokenString
}
