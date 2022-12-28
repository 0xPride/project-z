package post

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type message struct {
	Message string
}

type res struct {
	Message string
}

type Nwita struct {
	gorm.Model
	Content string
}

func PostHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var m message
	var response res

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

	res := db.Create(&Nwita{Content: m.Message})

	if res.Error != nil {
		// http.Error(w, "could not create record in db", http.)
	}

	response.Message = m.Message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	b, err := json.Marshal(&response)
	if err != nil {
		response.Message = "could not response with inserted data as json"
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(b)
}
