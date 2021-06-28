package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Doslo")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getQuestions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
