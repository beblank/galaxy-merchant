package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Result string
}

var result string

func galaxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
	}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", galaxy)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
