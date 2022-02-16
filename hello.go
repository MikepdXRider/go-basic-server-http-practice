package main;

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	);

// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	// *Author refers to a custom type.
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice book struct. slice is basically a variable length array.
var books []Book

// Get all Books
func getBooks(res http.ResponseWriter, req *http.Request) {
	fmt.Println("getBooks invoked through ...8000/api/v1/books");
	res.Header().Set("Content-Type", "application/json");
	json.NewEncoder(res).Encode(books)
}

// Get single book
func getBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("getBook invoked through ...8000/api/v1/books");
	params := mux.Vars(req); // get params

	// Loop through books and find with id
	for _, item:= range books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		} 
	}
	json.NewEncoder(res).Encode( &Book{})
}

// Create a new book
func createBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("createBook invoked through ...8000/api/v1/books")
	res.Header().Set("Content-Type", "application/json") // sets header 
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book) // Dunno what this does?
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - note safe
	books = append(books, book) // append our new book to books
	json.NewEncoder(res).Encode(book) // send back the book
}

// Update a book
func updateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("updateBook invoked through ...8000/api/v1/books")
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	json.NewEncoder(res).Encode(books)
}

// Delete a book
func deleteBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("deleteBook invoked through ...8000/api/v1/books")
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(books)
}

// In go we use func to declare a function. 
// No need to call a fn, it will run automatically. 
func main() {
	// init router
	router := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "11111", Title: "Mock Book One", Author: &Author{Firstname: "Mock", Lastname: "Author"}})
	books = append(books, Book{ID: "2", Isbn: "22222", Title: "Mock Book Two", Author: &Author{Firstname: "Mock2", Lastname: "Author"}})

	// Route Handlers / Endpoints
	// Should be imported
	router.HandleFunc("/api/v1/books", getBooks).Methods("GET");
	router.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET");
	router.HandleFunc("/api/v1/books", createBook).Methods("POST");
	router.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT");
	router.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE");

	// Start and listen on server
	// log if fails. 
	log.Fatal(http.ListenAndServe(":8000", router))

}