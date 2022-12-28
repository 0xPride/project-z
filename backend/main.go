package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/0xpride/project-z/get"
	"github.com/0xpride/project-z/post"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// handle post request
		post.PostHandler(w, r, db)
	} else if r.Method == "GET" {
		get.GetHandler(w, r, db)
	} else {
		// 404 not found
	}
}

func main() {

	d, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("could not open db duo to: ", err.Error())
	}
	d.AutoMigrate(&post.Nwita{})

	db = d
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	db, err := sql.Open("sqlite3", "./database.db")
// 	if err != nil {
// 		log.Fatal("could not open db ", err.Error())
// 	}
// 	sta, err := db.Prepare("INSERT INTO nwita(content) VALUES(?)")
// 	_, err = sta.Exec("wassup madafaka")
// 	sta.Close()
// 	if err != nil {
// 		log.Fatal("cant insert to db ", err.Error())
// 	}
// 	res, err := db.Query("SELECT * FROM nwita")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var id int32
// 	var content string
// 	// res.Next()
// 	for res.Next() {
// 		res.Scan(&id, &content)
// 		log.Println(id, " ", content)
// 	}
// 	res.Close()
// 	db.Close()
// }
