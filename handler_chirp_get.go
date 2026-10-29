package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpGetById(w http.ResponseWriter, r *http.Request) {

	chirpId, err := uuid.Parse(r.PathValue("ChirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing ChirpID")
		return
	}

	dbchirp, err := cfg.db.GetChirpById(r.Context(), chirpId)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp was not found")
		return
	}

	chirp := Chirp{
		ID:        dbchirp.ID,
		CreatedAt: dbchirp.CreatedAt,
		UpdatedAt: dbchirp.UpdatedAt,
		Body:      dbchirp.Body,
		UserID:    dbchirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirp)
}
