package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func Reaction(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.WriteError(w,"Unable to read body",err,http.StatusInternalServerError)
		return
	}
	var reaction QuestionReaction
	json.Unmarshal(req,&reaction)
	_,err =db.Connection.Exec("INSERT INTO question_reactions(question_id, author_id, rating) VALUES (?,?,?) ON DUPLICATE KEY UPDATE rating=? ",reaction.QuestionId,reaction.AuthorId,reaction.Reaction,reaction.Reaction)
	if err!=nil{
		utils.WriteError(w,"Unable to react to question",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

