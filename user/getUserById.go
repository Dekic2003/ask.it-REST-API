package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"main/utils"
	"net/http"
)



func GetUserById(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]

	var user returnUserById
	err := db.Connection.QueryRow("SELECT id, name, surname, email FROM users WHERE id=?",id).Scan(&user.Id,&user.Name,&user.Surname,&user.Email)
	if err!=nil {
		utils.WriteError(w,"Unable to fetch user",err,http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w,user,true)
}
