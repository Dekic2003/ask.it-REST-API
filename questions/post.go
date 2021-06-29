package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"net/http"
)

type NewQuestion struct {
	AuthorId int `json:"author_id"`
	Question string `json:"question"`
}
func Post(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var question NewQuestion
	json.Unmarshal(req,&question)
	db.Connection.Query("INSERT INTO question(author_id,question) VALUES (?,?) ",question.AuthorId,question.Question)

}

