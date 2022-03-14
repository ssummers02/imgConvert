package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func sendFile(w http.ResponseWriter, r *http.Request, fileName string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, fileName)
}
