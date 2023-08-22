package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/time", timeHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		_, err := fmt.Fprintf(writer, "ParseForm() err: %v", err)
		checkError(err)
		return
	}

	_, err := fmt.Fprintf(writer, "POST request succesful")
	checkError(err)

	name := request.FormValue("name")
	address := request.FormValue("address")

	_, err = fmt.Fprintf(writer, "Name - %s\n", name)
	checkError(err)
	_, err = fmt.Fprintf(writer, "Address - %s\n", address)
	checkError(err)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method is not supported", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(writer, "hello!")
	if err != nil {
		return
	}
}

func timeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/time" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method is not supported", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(writer, time.Now().String())
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
