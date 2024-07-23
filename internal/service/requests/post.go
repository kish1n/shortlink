package requests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/kish1n/shortlink/internal/data"
	"log"
	"net/http"
)

type AddLinkRequest struct {
	Original string `json:"original"`
}

type AddLinkResponse struct {
	Shortened string `json:"shortened"`
}

func AddLink(db *sql.DB, original string) (string, error) {
	var shortened string

	err := db.QueryRow("SELECT shortened FROM link WHERE original = $1", original).Scan(&shortened)
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("error checking existing link: %v", err)
	}

	if shortened != "" {
		return shortened, nil
	}

	shortened = GenShortURL()

	_, err = db.Exec("INSERT INTO link (original, shortened) VALUES ($1, $2)", original, shortened)
	if err != nil {
		return "", fmt.Errorf("error inserting new link: %v", err)
	}

	return shortened, nil
}

func AddLinkHandler(w http.ResponseWriter, r *http.Request) {
	var req AddLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db, err := data.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Printf("AddLinkHandler: %v", err)
		return
	}
	defer db.Close()

	shortened, err := AddLink(db, req.Original)
	if err != nil {
		http.Error(w, "Error adding link", http.StatusInternalServerError)
		log.Printf("AddLinkHandler: %v", err)
		return
	}

	resp := AddLinkResponse{Shortened: shortened}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
