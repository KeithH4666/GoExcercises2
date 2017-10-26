package main

import (
	"net/http"
	"log"
	"text/template"
	"time"
	"math/rand"
	"strconv"

)


type guessStr struct{

	Guess string

}


func guess(w http.ResponseWriter, r *http.Request){
	
	rand.Seed(time.Now().UnixNano())//Seeds time
	
	expiration := time.Now().Add(365 * 24 * time.Hour)  ////Set cookie to expire in 1 year
	cookie := http.Cookie{Name: "CookVal", Value:strconv.Itoa(rand.Intn(20)) , Expires: expiration}//Sets values to cookie + converts num between 1-20 to dtring
	http.SetCookie(w, &cookie)//set cookie
	
	//Generate Template
	t, _ := template.ParseFiles("template/time.html")
	t.Execute(w, guessStr{Guess: "Pick a number between 1-20:"})

}

func main() {
	
	//Handles static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/guess",guess)

	//Displays user messgae to console 
    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}