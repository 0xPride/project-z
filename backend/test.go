package backend

import (
	"database/sql"
	"log"
)

func _main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("could not open datanbase ", err.Error())
	}
	db.
}
