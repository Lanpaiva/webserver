package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Estudo struct {
	Nome   string
	Numero int
}

type Estudos []Estudo

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	http.HandleFunc("/", HttpHandle)
	http.ListenAndServe(":8990", nil)
}

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	{
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"toUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(w, Estudos{
			{"Go", 0},
			{"Node", 1},
			{"Typescript", 2},
		})
		if err != nil {
			panic(err)
		}
	}
}
