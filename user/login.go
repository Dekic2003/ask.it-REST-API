package user

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var user LoginUser
	var userDB DbUser

	json.Unmarshal(req,&user)

	password:=[]byte(user.Password)

	err = db.Connection.QueryRow("SELECT * FROM users where email=? ",user.Email).Scan(&userDB.Id,&userDB.Name,&userDB.Surname,&userDB.Email,&userDB.Password,&userDB.CreatedAt,&userDB.UpdatedAt)
	if err !=nil{
		utils.WriteError(w,"Unable to fetch user",err,http.StatusInternalServerError)
		return
	}
	passwordDB:=[]byte(userDB.Password)
	err =bcrypt.CompareHashAndPassword(passwordDB,password);
	if err!=nil{
		utils.WriteError(w,"Password incorrect",err,http.StatusUnprocessableEntity)
		return
	}

		var userR returnUser
		userR.AccessToken=utils.JWTGenerator(userDB.Name,userDB.Id)
		userR.Name=userDB.Name
		userR.Id=userDB.Id
		userR.Surname=userDB.Surname
		userR.Email=userDB.Email
		utils.WriteSuccess(w,userR)
}

