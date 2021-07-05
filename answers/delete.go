package answers

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
	_, err := db.Connection.Query("DELETE FROM notifications WHERE answer_author_id = (SELECT author_id FROM answers WHERE id=? )",id)
	if err != nil {
		utils.WriteError(w,"Unable to delete notifications",err,http.StatusInternalServerError)
		return
	}
	_, err = db.Connection.Query("DELETE FROM answer_reactions WHERE answer_reactions.answer_id = ?",id)
	if err != nil {
		utils.WriteError(w,"Unable to delete question reactions",err,http.StatusInternalServerError)
		return
	}
	_, err = db.Connection.Query("DELETE FROM answers WHERE id = ?",id)
	if err != nil {
		utils.WriteError(w,"Unable to delete",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

