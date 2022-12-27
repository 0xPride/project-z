package post

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type message struct {
	Message string
}

type res struct {
	Message string
	Status  int32
}

func PostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var m message
	var response res

	// check if json in header
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "json is the only supported format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if m.Message == "" {
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}

	stm, err := db.Prepare("INSERT INTO nwita(content) VALUES(?)")
	if err != nil {
		http.Error(w, "could create query", http.StatusInternalServerError)
	}
	stm.Close()
	_, err = stm.Exec(m.Message)
	if err != nil {
		http.Error(w, "could not save data", http.StatusInternalServerError)
	}

	response.Status = 200
	response.Message = m.Message

	b, err := json.Marshal(&response)
	if err != nil {
		response.Status = 400
		response.Message = "could not response with inserted data as json"
	}
	w.Write(b)
}
