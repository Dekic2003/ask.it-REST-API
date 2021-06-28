package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"main/db"
	"main/questions"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err !=nil{
		panic(err)
	}
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", questions.Get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}