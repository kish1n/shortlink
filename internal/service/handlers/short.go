package handlers

import (
	"encoding/json"
	"github.com/kish1n/shortlink/internal/data"
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

	logger.Infof("req ini")

	original := request.Original
	db := helpers.DB(r)

	logger.Infof("db con")

	var pair data.CoupleLinks

	for {
		shortened := requests.GenShortURL()
		logger.Infof(shortened)
		existingLink, err := db.Link().FilterByShortened(shortened).Get()
		logger.Info(existingLink)
		if err != nil {
			logger.WithError(err).Error("failed to query db handlers/short.go 35")
			ape.RenderErr(w)
			return
		}
		if existingLink == nil {
			pair = data.CoupleLinks{
				Shortened: shortened,
				Original:  original,
			}
			break
		}
	}

	logger.Infof("checked")

	var req LinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.WithError(err).Debug("handlers/short.go 26")
		return
	}

	insertedPair, err := db.Link().Insert(pair)

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
}
