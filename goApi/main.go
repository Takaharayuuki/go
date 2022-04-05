package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Author  *Author `json:"author"`
}

type Author struct {
	FirstName  string  `json:"firstname"`
	LastName  string  `json:"lastname"`
}

var books []Book

func main() {
	// ルーティングのイニシャライズ
	r := mux.NewRouter()

	// モックデータの作成
	books = append(books, Book{ID: "1", Title: "Book one", Author: &Author{FirstName: "Philip", LastName: "Williams"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Johnson"}})

	// ルートのエンドポイント
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books.{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}