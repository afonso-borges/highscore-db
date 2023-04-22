package handlers

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Ping string `json:"ping"`
}

func Ping(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{"pong"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
