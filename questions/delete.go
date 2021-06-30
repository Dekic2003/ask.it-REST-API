package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.WriteError(w,"Unable to read body",err,http.StatusBadRequest)
		return
	}
	var question DeleteQuestion
	json.Unmarshal(req,&question)
	_, err = db.Connection.Query("DELETE FROM question WHERE id = ?", question.Id)
	if err != nil {
		utils.WriteError(w,"Unable to delete",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

