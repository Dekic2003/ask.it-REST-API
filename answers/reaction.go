package answers

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
		panic(err)
	}
	var reaction AnswerReaction
	json.Unmarshal(req,&reaction)
	_,err =db.Connection.Exec("INSERT INTO answer_reactions(answer_id, author_id, rating) VALUES (?,?,?) ON DUPLICATE KEY UPDATE rating=? ",reaction.AnswerId,reaction.AuthorId,reaction.Reaction,reaction.Reaction)
	if err!=nil{
		utils.WriteError(w,"Unable to react to answer",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

