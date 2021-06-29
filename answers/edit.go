package answers
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
	var answer EditedAnswer
	json.Unmarshal(req,&answer)
	db.Connection.Query("UPDATE answer SET answer = ? WHERE id = ?",answer.Answer,answer.Id)

}

