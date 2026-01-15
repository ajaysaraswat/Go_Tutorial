package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to go module")
	gretter()
	r := mux.NewRouter()
	r.HandleFunc("/",servHome).Methods("Get")

	log.Fatal(http.ListenAndServe(":5000",r))
}

func gretter(){
	fmt.Println("hey! there mod users")
}

func servHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to my module</h1>"))
}