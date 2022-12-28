package get

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/0xpride/project-z/post"
	"gorm.io/gorm"
)

func GetHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var nwitat [10]post.Nwita

	queryRes := db.Limit(10).Find(&nwitat)
	for _, v := range nwitat {
		log.Println(v)
	}
	if queryRes.Error != nil {
		http.Error(w, "could not fetch data", http.StatusInternalServerError)
	}
	res, err := json.Marshal(&nwitat)
	if err != nil {
		http.Error(w, "could not send json", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
