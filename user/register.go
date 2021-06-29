package user

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"main/db"
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
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword(hashedPass,password)
	if err != nil {
		panic(err)
	}
	res,err:=db.Connection.Exec("INSERT INTO users(name, surname, email, password) VALUES (?,?,?,?)",user.Name,user.Surname,user.Email,hashedPass)

}

