package post

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Nwita struct {
	gorm.Model
	Content string
}

func PostHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// var m message
	// var response res
	var nwita Nwita

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "json is the only supported format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&nwita); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if nwita.Content == "" {
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}

	res := db.Create(&Nwita{Content: nwita.Content})

	if res.Error != nil {
		http.Error(w, "could not create record in db", http.StatusInternalServerError)
	}

	b, err := json.Marshal(&nwita)
	if err != nil {
		http.Error(w, "could not responsd", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
