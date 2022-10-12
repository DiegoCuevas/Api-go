package main

import (
	"encoding/json" // modulo la codificacion y la decodificacion de Json
	"log"            // modulo para debuggear
	"net/http"       // modulo escribir las peticiones http
	"github.com/gorilla/mux" // modulo para las rutas
)


func main()  {
	router := mux.NewRouter()
	
	// Rutas
	router.HandleFunc("/book", GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}