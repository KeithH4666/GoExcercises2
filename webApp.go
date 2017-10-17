package main

import (
	"net/http"
	"log"
)

/*
func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","text/html");// allows browser to render html tags

	fmt.Fprintln(w, "<h1>Guessing Game</h1>")
}
*/

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}