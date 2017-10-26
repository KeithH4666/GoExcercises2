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
	//Random
	rand.Seed(time.Now().UTC().UnixNano())
	//Target variable set random
	target:=rand.Intn(20-1)
	var cookie, err = r.Cookie("target")
	
	//Ians Cookie//
	if err == nil{
		//if we could read it ,try to convert its value to an int
		target, _ = strconv.Atoi(cookie.Value)
		
	}
	cookie = &http.Cookie{
		Name: "target",
		Value: strconv.Itoa(target),
		Expires: time.Now().Add(72 * time.Hour),
	}
	
	//set the cookie
	http.SetCookie(w,cookie)
	
	
	
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