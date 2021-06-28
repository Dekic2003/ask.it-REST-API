package questions

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/db"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	dbresponse,err :=db.Connection.Query("SELECT * FROM users")
	fmt.Println(dbresponse)
	fmt.Println(err)
}

