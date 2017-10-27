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
	Message string
	Guess string
	userGuess int
	DisUserG int

}

func guess(w http.ResponseWriter, r *http.Request){

	userConvo:=""
	//Random
	rand.Seed(time.Now().UTC().UnixNano())
	//Target variable set random
	rando:=rand.Intn(20-1)
	var cookie, err = r.Cookie("rando")
	
	//Ians Cookie Method checks if value is set//
	if err == nil{
		//if we could read it ,try to convert its value to an int
		rando, _ = strconv.Atoi(cookie.Value)
	}
	
	//Sets the userGuess var to input of form guess + converts to int
	userGuess,_ := strconv.Atoi(r.FormValue("Guess"))
	
	var dispG = userGuess
	
	//Checks if user guess is equal to target
	if userGuess == rando{
		userConvo="Correct Guess - Cookie value has been reset - Try Again!"
		rando=rand.Intn(20-1)
		
	}else if userGuess < rando{
		userConvo="Try Again your guess  was  too low!"
	}else{
		userConvo="Try Again your guess was too high!"
	}
	
	//Cookie
	cookie = &http.Cookie{
		Name: "rando",
		Value: strconv.Itoa(rando),
		Expires: time.Now().Add(72 * time.Hour),
	}
	
	//set the cookie
	http.SetCookie(w,cookie)
	
	
	//Generate Template
	t, _ := template.ParseFiles("template/guess.html")
	t.Execute(w, &guessStr{Message: "Pick a number between 0-20:",Guess:userConvo,DisUserG:dispG})

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