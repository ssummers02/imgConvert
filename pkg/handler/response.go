package handler

import (
	"encoding/json"
	"net/http"
)

// DefaultResponse is a JSON response in case of success.
type DefaultResponse struct {
	IsOK bool        `json:"is_ok"`
	Data interface{} `json:"data"`
}

// DefaultError is a JSON response in case of failure.
type DefaultError struct {
	Text string `json:"text"`
}

type LinkResponse struct {
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
}

// SendErr sends a response to the client in case of success.
func SendErr(w http.ResponseWriter, code int, text string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(
		DefaultError{Text: text},
	)
}

// SendOK sends a response to the client in case of success.
func SendOK(w http.ResponseWriter, code int, p interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	// These two do not allow body
	_ = json.NewEncoder(w).Encode(
		p,
	)
}
