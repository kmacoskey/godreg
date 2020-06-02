package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=kmacoskey sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListNotes)
	http.HandleFunc("/note.html", handleViewNote)
	http.HandleFunc("/save", handleSaveNote)
	http.HandleFunc("/delete", handleDeleteNote)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
