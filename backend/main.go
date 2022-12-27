package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/0xpride/project-z/post"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// handle post request
		post.PostHandler(w, r, db)
	} else if r.Method == "GET" {

	} else {
		// 404 not found
	}
}

func main() {
	http.HandleFunc("/", handler)
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("cant open database.db file", err.Error())
	}
	defer db.Close()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server running on http://localhost:8080")
}
