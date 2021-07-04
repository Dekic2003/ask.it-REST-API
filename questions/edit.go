package questions

import (
"encoding/json"
_ "github.com/go-sql-driver/mysql"
"io/ioutil"
"main/db"
	"main/utils"
	"net/http"
)


func Edit(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.WriteError(w,"Unable to read body",err,http.StatusBadRequest)
		return
	}
	var question EditedQuestion
	json.Unmarshal(req,&question)
	_,err = db.Connection.Exec("UPDATE questions SET question = ? WHERE id = ?", question.Question, question.Id)
	if err != nil {
		utils.WriteError(w,"Unable to update",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)


}

