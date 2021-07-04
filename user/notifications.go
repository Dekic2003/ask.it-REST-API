package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"main/utils"
	"net/http"
)

func Notifications(w http.ResponseWriter, r *http.Request) {

	id:=r.Header.Get("id")
	var notifiations []Notification
	res,err := db.Connection.Query("SELECT id,question_id,question_author_id,answer_author_id,(SELECT email FROM users  WHERE id=answer_author_id) as author_answer_name,isRead FROM notifications WHERE question_author_id=? AND isRead=false",id)
	if err!=nil{
		utils.WriteError(w,"Unable to fetch notifiactions",err,http.StatusInternalServerError)
		return
	}
	for res.Next(){
		var notification Notification
		err=res.Scan(&notification.Id,&notification.QuestionId,&notification.QuestionAuthorId,&notification.AnswerAuthorId,&notification.AuthorAnswerName,&notification.IsRead)
		if err != nil{
			utils.WriteError(w,"Unable to scan notifiaction",err,http.StatusInternalServerError)
			return
		}
		notifiations=append(notifiations,notification)
	}

	utils.WriteSuccess(w,notifiations,true)
}
func NotificationMark(w http.ResponseWriter, r *http.Request) {
	params:=mux.Vars(r)
	id:=params["id"]
	_,err := db.Connection.Exec("UPDATE notifications SET isRead=true WHERE id=?",id)
	if err != nil {
		utils.WriteError(w, "Unable to update notification", err, http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"Updated notification",true)
}
func NotificationMarkAll(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]
	_,err := db.Connection.Exec("UPDATE notifications SET isRead=true WHERE question_author_id=?",id)
	if err != nil {
		utils.WriteError(w, "Unable to update notification", err, http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"Updated notification",true)
}

