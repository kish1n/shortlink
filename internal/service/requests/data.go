package requests

import (
	"github.com/kish1n/shortlink/internal/data"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func DBHandler(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("DBHandler: %v", err)
		return
	}
	defer db.Close()
}
