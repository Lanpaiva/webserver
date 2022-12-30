package main

import (
	"net/http"
	"text/template"
)

type Estudo struct {
	Nome   string
	Numero int
}

type Estudos []Estudo

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Estudos{
			{"Go", 0},
			{"Node", 1},
			{"Typescript", 2},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8990", nil)
}
