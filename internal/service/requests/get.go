package requests

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kish1n/shortlink/internal/data"
	"log"
	"net/http"
)

func GetLink(db *sql.DB, shortened string) (string, error) {
	var original string

	err := db.QueryRow("SELECT original FROM links WHERE shortened = $1", shortened).Scan(&original)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Сокращенная ссылка не найдена
		}
		return "", fmt.Errorf("error checking existing link: %v", err)
	}

	return original, nil
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortened := chi.URLParam(r, "shortened")

	db, err := data.InitDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Printf("RedirectHandler: %v", err)
		return
	}
	defer db.Close()

	original, err := GetLink(db, shortened)
	if err != nil {
		http.Error(w, "Error getting link", http.StatusInternalServerError)
		log.Printf("GetLink: %v", err)
		return
	}

	if original == "" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
	log.Printf("Redirecting to: %s", original)
}
