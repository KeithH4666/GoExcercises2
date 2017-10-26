package main

import (
	"net/http"
	"log"
	"text/template"

)


type guessStr struct{

	Guess string

}

func guess(w http.ResponseWriter, r *http.Request){

	t, _ := template.ParseFiles("template/time.html")
	t.Execute(w, guessStr{Guess: "Pick a number between 1-20:"})

}

func main() {
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/guess",guess)


    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}