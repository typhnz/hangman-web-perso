package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/play", Play)
	http.HandleFunc("/present", Present)
	fileServer := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	var fileName = "home.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file.", err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template.", err)
		return
	}
}

func Play(w http.ResponseWriter, r *http.Request) {
	var fileName = "play.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file.", err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template.", err)
		return
	}
}

func Present(w http.ResponseWriter, r *http.Request) {
	var fileName = "present.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file.", err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template.", err)
		return
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {

	// generate new data struct
	type data struct {
		attemptsRemaining int
		guessAttempt      string
		success           bool
	}

	// Read the guess value
	guess := r.FormValue("guess")
	fmt.Println(guess)
}
