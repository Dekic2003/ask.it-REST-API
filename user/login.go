package user

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"main/db"
	"net/http"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var signingKey= []byte(os.Getenv("TOKEN_KEY"))
	var user LoginUser
	var userDB DbUser

	json.Unmarshal(req,&user)

	password:=[]byte(user.Password)

	results, err := db.Connection.Query("SELECT * FROM users where email=? ",user.Email)
	for results.Next(){
		err=results.Scan(&userDB.Id,&userDB.Name,&userDB.Surname,&userDB.Email,&userDB.Password,&userDB.CreatedAt,&userDB.UpdatedAt)
		if err!=nil{
			panic(err)
		}
	}
	passwordDB:=[]byte(userDB.Password)
	err =bcrypt.CompareHashAndPassword(passwordDB,password);
	if err!=nil{
		panic(err)
	}
		token:=jwt.New(jwt.SigningMethodHS256)

		claims:=token.Claims.(jwt.MapClaims)

		claims["authorized"]=true
		claims["user"]=userDB.Email

		tokenString,err :=token.SignedString(signingKey)
		if err!=nil{
			panic(err)
		}
		var userR returnUser
		userR.AccessToken=tokenString
		userR.Name=userDB.Name
		userR.Id=userDB.Id
		userR.Surname=userDB.Surname
		userR.Email=userDB.Email
		res,err := json.Marshal(userR)
		if err !=nil {
			panic(err)
		}
		w.Header().Set("Content-Type","application/json")
		w.Write(res)

}

