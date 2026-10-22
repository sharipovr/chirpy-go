package main

import (
	"encoding/json"
	"net/http"
)

type chirp struct {
	Body string `json:"body"`
}

type validResponse struct {
	Valid bool `json:"valid"`
}

func handlerChirpValidate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	chirpBody := chirp{}
	err := decoder.Decode(&chirpBody)

	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}
	if len(chirpBody.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	respondWithJSON(w, 200, validResponse{Valid: true})
}
