package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"main/answers"
	"main/db"
	"main/middleware"
	"main/questions"
	"main/user"
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
	r.HandleFunc("/question/{id}", questions.GetByQuestionId).Methods("GET")
	r.HandleFunc("/author/{id}", questions.GetByAuthorId).Methods("GET")
	r.HandleFunc("/",middleware.ValidateToken(questions.Post)).Methods("POST")
	r.HandleFunc("/", middleware.ValidateToken(questions.Edit)).Methods("PUT")
	r.HandleFunc("/", middleware.ValidateToken(questions.Delete)).Methods("DELETE")
	r.HandleFunc("/question/reaction",middleware.ValidateToken(questions.Reaction)).Methods("POST")
	r.HandleFunc("/register", user.Register).Methods("POST")
	r.HandleFunc("/login", user.Login).Methods("POST")
	r.HandleFunc("/answer/{id}", answers.Get).Methods("GET")
	r.HandleFunc("/answer", middleware.ValidateToken(answers.Post)).Methods("POST")
	r.HandleFunc("/answer/reaction",middleware.ValidateToken(answers.Reaction)).Methods("POST")
	r.HandleFunc("/answer", middleware.ValidateToken(answers.Edit)).Methods("PUT")
	r.HandleFunc("/answer", middleware.ValidateToken(answers.Delete)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
