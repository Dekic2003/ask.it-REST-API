package user

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func EditPassword(w http.ResponseWriter, r *http.Request)  {
	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.WriteError(w,"Unable to read body",err,http.StatusInternalServerError)
		return
	}
	var user ChangePasswordUser
	var userDB DbUser

	json.Unmarshal(req,&user)

	password:=[]byte(user.Password)
	newPassword:=[]byte(user.NewPassword)

	err = db.Connection.QueryRow("SELECT * FROM users where id=? ",user.Id).Scan(&userDB.Id,&userDB.Name,&userDB.Surname,&userDB.Email,&userDB.Password,&userDB.CreatedAt,&userDB.UpdatedAt)
	if err !=nil{
		utils.WriteError(w,"Unable to fetch user",err,http.StatusInternalServerError)
		return
	}
	passwordDB:=[]byte(userDB.Password)
	err =bcrypt.CompareHashAndPassword(passwordDB,password);
	if err!=nil{
		utils.WriteError(w,"Old Password incorrect",err,http.StatusUnprocessableEntity)
		return
	}
	hashedPass,err:=bcrypt.GenerateFromPassword(newPassword,bcrypt.DefaultCost)
	if err !=nil{
		utils.WriteError(w,"Unable to hash password",err,http.StatusInternalServerError)
		return
	}
	_,err=db.Connection.Exec("UPDATE users SET password=? WHERE id=?",hashedPass,user.Id)
	if err !=nil{
		utils.WriteError(w,"Unable to replace password",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"Password Changed",true)
}

