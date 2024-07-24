package handlers

import (
	"encoding/json"
	"github.com/kish1n/shortlink/internal/service/helpers"
	"github.com/kish1n/shortlink/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

type LinkRequest struct {
	Original string `json:"original"`
}

func GetShort(w http.ResponseWriter, r *http.Request) {

	logger := helpers.Log(r)
	request, err := requests.NewLinkRequest(r)

	if err != nil {
		logger.WithError(err).Debug("bad request handlers/short.go 21")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	original := request.Original
	db := helpers.DB(r)

	logger.Infof("db con")

	res, err := db.Link().FilterByOriginal(original)
	if err == nil {
		logger.Infof("here's already a link res %s", res)
		response := map[string]string{
			"shortened": res.Shortened,
			"original":  res.Original,
		}
		ape.Render(w, response)
		return
	}

	var req LinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.WithError(err).Debug("handlers/short.go 26")
		return
	}

	insertedPair, err := db.Link().Insert(res)

	if err != nil {
		logger.WithError(err).Error("failed to query db handlers/short.go 41")
		ape.RenderErr(w)
		return
	}

	response := map[string]string{
		"shortened": insertedPair.Shortened,
		"original":  insertedPair.Original,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.WithError(err).Error("failed to encode response")
		ape.RenderErr(w)
	}

	ape.Render(w, response)
}
