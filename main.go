package main

import (
	c "countWebService/count"
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO Welcome to count string app. go to /count to start counting a word. And type in form-data to send the input.")
	http.Error(w, "Bad request", http.StatusBadRequest)
}

func countController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		input := r.FormValue("input")
		vowels, consonant := c.Count(input)
		fmt.Fprintf(w, vowels+"\n"+consonant)
		return
	}
	http.Error(w, "Bad request try POST method to start count a string", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", homepage)

	http.HandleFunc("/count", countController)

	port := "8000"

	fmt.Println("server running in port " + port)
	http.ListenAndServe(":"+port, nil)
}
