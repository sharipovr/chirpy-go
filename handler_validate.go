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
type cleanedChirp struct {
	CleanedBody string `json:"cleaned_body"`
}

func handlerChirpValidate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	chirpData := chirp{}
	err := decoder.Decode(&chirpData)

	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}
	if len(chirpData.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	cleanedText := helperBadWordsClean(chirpData.Body)

	respondWithJSON(w, 200, cleanedChirp{CleanedBody: cleanedText})
}
