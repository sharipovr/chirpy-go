package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {

	dbchirps, err := cfg.db.GetChirps(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	chirps := make([]Chirp, 0, len(dbchirps))

	for _, dbchirp := range dbchirps {
		chirps = append(chirps, Chirp{
			ID:        dbchirp.ID,
			CreatedAt: dbchirp.CreatedAt,
			UpdatedAt: dbchirp.UpdatedAt,
			Body:      dbchirp.Body,
			UserID:    dbchirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, chirps)
}
