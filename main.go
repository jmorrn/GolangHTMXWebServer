package main

import (
	"fmt" 
	"io"
	"net/http" 
	"log"
)


func main(){
	fmt.Println("Hello World")
	
	handlerOne := func(writer http.ResponseWriter, request *http.Request){
		io.WriteString(writer, "Hello World")
		io.WriteString(writer,request.Method)
	}

	http.HandleFunc("/", handlerOne)

	log.Fatal(http.ListenAndServe(":3333", nil))
}