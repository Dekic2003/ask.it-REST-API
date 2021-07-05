package questions

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"main/utils"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]
	_, err := db.Connection.Query("DELETE FROM notifications WHERE question_id = ?",id)
	if err != nil {
		utils.WriteError(w,"Unable to delete notifications",err,http.StatusInternalServerError)
		return
	}
	_, err = db.Connection.Query("DELETE FROM question_reactions WHERE question_id = ?",id)
	if err != nil {
		utils.WriteError(w,"Unable to delete question reactions",err,http.StatusInternalServerError)
		return
	}
	_, err = db.Connection.Query("DELETE FROM questions WHERE id = ?",id)
	if err != nil {
		utils.WriteError(w,"Unable to delete",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

