package requests

import (
	"fmt"
	"github.com/kish1n/shortlink/internal/data"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func DBHandler(w http.ResponseWriter, r *http.Request) {
	db, err := data.InitDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("DBHandler: %v", err)
		return
	}
	defer db.Close()

	var currentTime string
	err = db.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("DBHandler: %v", err)
		return
	}

	fmt.Fprintf(w, "Current time: %s", currentTime)
}
