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

func Register(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var user NewUser



	json.Unmarshal(req,&user)
	password :=[]byte(user.Password)

	hashedPass,err:=bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)
	if err !=nil{
		utils.WriteError(w,"Unable to hash password",err,http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword(hashedPass,password)
	if err != nil {
		utils.WriteError(w,"Password not matching hash",err,http.StatusInternalServerError)
		return
	}
	res,err:=db.Connection.Exec("INSERT INTO users(name, surname, email, password) VALUES (?,?,?,?)",user.Name,user.Surname,user.Email,hashedPass)
	userid,err:=res.LastInsertId()
	if err != nil{
		utils.WriteError(w,"Error occured while fetching id",err,http.StatusInternalServerError)
	}
	var userR returnUser
	userR.AccessToken=utils.JWTGenerator(user.Name, int(userid))
	userR.Name=user.Name
	userR.Id= int(userid)
	userR.Surname=user.Surname
	userR.Email=user.Email
	utils.WriteSuccess(w,userR,true)

}

