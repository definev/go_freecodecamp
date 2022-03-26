package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() failed: Error = %v", err.Error())
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %v", name)
	fmt.Fprintf(w, "Address = %v", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	staticServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", staticServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Running on localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf(err.Error())
	}
}
