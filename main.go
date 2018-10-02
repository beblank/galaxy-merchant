package main

import (
	"html/template"
	"log"
	"metal/calculator"
	"metal/converter"
	"metal/validator"
	"net/http"
	"strconv"
)

type Data struct {
	Result string
}

var result string

func galaxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
	}
	input := r.FormValue("input")
	metal := r.FormValue("submit")
	roman := converter.ConvertGalaxy(input)
	validated := validator.IsValid(roman)
	calculated := calculator.Calculate(roman, metal)
	if validated {
		result = input + " is " + strconv.FormatFloat(calculated, 'f', 1, 64) + " Credits"
	} else {
		result = "input not valid"
	}
	if input == "No Input" {
		result = input
	}
	t, _ := template.ParseFiles("views/index.html")
	d := Data{Result: result}
	t.Execute(w, d)
}

func main() {
	http.HandleFunc("/", galaxy)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
