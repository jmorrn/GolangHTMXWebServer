package main

import (
	"fmt" 
	"html/template"
	"net/http" 
	"log"
)


func main(){
	fmt.Println("Hello World")
	
	handlerOne := func(writer http.ResponseWriter, request *http.Request){
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(writer, nil)
	}

	http.HandleFunc("/", handlerOne)

	log.Fatal(http.ListenAndServe(":3333", nil))
}