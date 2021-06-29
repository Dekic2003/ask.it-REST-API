package questions

import (
"encoding/json"
_ "github.com/go-sql-driver/mysql"
"io/ioutil"
"main/db"
"net/http"
)


func Edit(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var question EditedQuestion
	json.Unmarshal(req,&question)
	db.Connection.Query("UPDATE question SET question = ? WHERE id = ?",question.Question,question.Id)

}

