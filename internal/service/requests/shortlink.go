package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"time"
)

type LinkRequest struct {
	Original string `json:"original"`
}

type ShortenedRequest struct {
	Shortened string `json:"shortened"`
}

func NewLinkRequest(r *http.Request) (LinkRequest, error) {
	var request LinkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, err
}

func ShortenedLinkRequest(r *http.Request) (string, error) {
	shortened := chi.URLParam(r, "shortened")
	if shortened == "" {
		return "", errors.New("shortened parameter is required")
	}
	return shortened, nil
}

func GenShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
