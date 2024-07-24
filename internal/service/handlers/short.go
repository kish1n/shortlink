package handlers

import (
	"encoding/json"
	"github.com/google/jsonapi"
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

	res, err := db.Link().FilterByOriginal(original)

	if err != nil {
		logger.Infof("not found %s", res)
		res.Shortened = requests.GenShortURL()
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
		return
	}

	logger.Infof("here's already a link res %s", res)
	response := map[string]string{
		"shortened": res.Shortened,
		"original":  res.Original,
	}
	ape.Render(w, response)
	return
}

func GetOriginal(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.ShortenedLinkRequest(r)

	if err != nil {
		logger.WithError(err).Debug("bad request handlers/short.go 71")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	shortened := request
	db := helpers.DB(r)

	res, err := db.Link().FilterByShortened(shortened)

	if res.Original == "" {
		logger.WithError(err).Debug("Not Found 404")
		ape.Render(w, &jsonapi.ErrorObject{
			Status: "404",
			Title:  "Not Found 404",
			Detail: "Nonexistent link",
		})
		return
	}

	if err != nil {
		logger.WithError(err).Debug("Server error")
		ape.Render(w, &jsonapi.ErrorObject{
			Status: "500",
			Title:  "Server error 500",
			Detail: "Unpredictable behavior",
		})
		return
	}

	logger.Infof("here's already a link res %s", res)
	response := map[string]string{
		"shortened": res.Shortened,
		"original":  res.Original,
	}

	ape.Render(w, response)
	return
}
