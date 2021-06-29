package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var question DeleteQuestion
	json.Unmarshal(req,&question)
	db.Connection.Query("DELETE FROM question WHERE id = ?",question.Id)

}

