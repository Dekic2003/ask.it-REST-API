package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"main/utils"
	"net/http"
	"os"
	"strings"
)

func ValidateToken(next http.HandlerFunc)http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var signingKey= []byte(os.Getenv("TOKEN_KEY"))
		if r.Header.Get("authorization") == ""{
			utils.WriteError(w,"Missing token",errors.New("missing token"),http.StatusUnauthorized)
			return
		}
		if r.Header.Get("authorization")!=""{
			token, err :=jwt.Parse(strings.Split(r.Header.Get("authorization")," ")[1], func(token *jwt.Token) (interface{}, error) {
				return signingKey,nil
			})
			if err!=nil{
				utils.WriteError(w,"Invalid token",err,http.StatusUnauthorized)
				return
			}
			if token.Valid{
				next(w,r)
			}
		}
	})
}