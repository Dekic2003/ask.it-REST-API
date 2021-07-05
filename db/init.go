package db

import (
	"database/sql"
	"fmt"
	"os"
)

var Connection *sql.DB


func Init()  {
	var err error
	dbName :=os.Getenv("DB_NAME")
	dbHost :=os.Getenv("DB_HOST")
	dbPass :=os.Getenv("DB_PASSWORD")
	dbUser :=os.Getenv("DB_USERNAME")
	Connection, err = sql.Open("mysql",fmt.Sprintf("%s:%s@%s/%s?reconnect=true?parseTime=true",dbUser,dbPass,dbHost,dbName))
	if err != nil{
		fmt.Println("Error occured",err)
		panic(err)
	}
	err = Connection.Ping()
	if err != nil{
		fmt.Println("Error pinging DB",err)
		panic(err)
	}
	fmt.Println("DB connected")
}
