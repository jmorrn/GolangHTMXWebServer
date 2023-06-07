package main

import (
	"fmt" 
	"html/template"
	"net/http" 
	"log"
)

type Movie struct {
	Title string
	Director string
}

func main(){
	fmt.Println("Hello World")
	
	handlerOne := func(writer http.ResponseWriter, request *http.Request){
		tmpl := template.Must(template.ParseFiles("index.html"))
		movies := map[string][]Movie{
			"Movies": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
				{Title: "The Avengers", Director: "Marvel"},
			},
		}
		tmpl.Execute(writer, movies)
	}

	http.HandleFunc("/", handlerOne)

	log.Fatal(http.ListenAndServe(":3333", nil))
}