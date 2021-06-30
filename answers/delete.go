package answers
import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.WriteError(w,"Unable to read body",err,http.StatusBadRequest)
		return
	}
	var answer DeleteAnswer
	json.Unmarshal(req,&answer)
	_, err = db.Connection.Query("DELETE FROM answer WHERE id = ?",answer.Id)
	if err != nil {
		utils.WriteError(w,"Unable to delete",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

