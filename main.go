package main

import (
	"fmt" 
	"time"
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

	handlerTwo := func(writer http.ResponseWriter, request *http.Request){
		time.Sleep(1 * time.Second)
		title := request.PostFormValue("title")
		director := request.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li>%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)

		tmpl.Execute(writer, nil)
	}

	http.HandleFunc("/", handlerOne)
	http.HandleFunc("/add-movie/", handlerTwo)

	log.Fatal(http.ListenAndServe(":3333", nil))
}