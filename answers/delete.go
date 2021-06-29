package answers
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
	var answer DeleteAnswer
	json.Unmarshal(req,&answer)
	db.Connection.Query("DELETE FROM answer WHERE id = ?",answer.Id)

}

