package main

import (
	"encoding/json" // modulo la codificacion y la decodificacion de Json
	"log"            // modulo para debuggear
	"net/http"       // modulo escribir las peticiones http
	"github.com/gorilla/mux" // modulo para las rutas
)

type Book struct {
	ID 			string 		`json: "id,omitempty"`
	Name 		string 		`json: "omit,omitempty"`
	Author 	*Author 	`json:"author,omitempty"`
}

type Author struct {
	FirstName 	string 		`json: "firstname,omitempty"`
	LastName 		string 		`json: "lastname,omitempty"`
}

var books []Book

// w: es para devolver una respuesta al navegador
// req: cuando el navegador envia data para usarlo en el servidor

func GetAllBooks(w http.ResponseWriter, req *http.Request )  {
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, req *http.Request )  {
	params := mux.Vars(req)
	for _, item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func CreateBook(w http.ResponseWriter, req *http.Request )  {
	params := mux.Vars(req)
	var newbook Book
	_ = json.NewDecoder(req.Body).Decode(&newbook)
	newbook.ID = params["id"]
	books = append(books, newbook)
	json.NewEncoder(w).Encode(books)
}


func main()  {
	router := mux.NewRouter()

	books = append(books, Book{ID:"1", Name: "Book1", Author: &Author{FirstName: "Diego", LastName: "Cuevas"}})
	books = append(books, Book{ ID:"2", Name: "Book2", Author: &Author{FirstName: "Diego", LastName: "Cuevas"}})

	// Rutas
	router.HandleFunc("/book", GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}