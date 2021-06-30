package user

import (
	"encoding/json"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func EditUserInfo(w http.ResponseWriter, r *http.Request)  {
	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.WriteError(w,"Unable to read body",err,http.StatusInternalServerError)
		return
	}
	var user ChangeUserInfo

	json.Unmarshal(req,&user)

	_,err=db.Connection.Exec("UPDATE users SET name=?,surname=?,email=? WHERE id=?",user.Name,user.Surname,user.Email,user.Id)
	if err != nil{
		utils.WriteError(w,"Unabel to change info",err,http.StatusUnprocessableEntity)
		return
	}
	utils.WriteSuccess(w,"User Info Changed",true)
}

