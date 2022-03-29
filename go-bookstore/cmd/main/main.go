package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/definev/go_freecodecamp/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Server serving at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
