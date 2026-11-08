package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sharipovr/chirpy-go/internal/auth"
	"github.com/sharipovr/chirpy-go/internal/database"
)

type chirp struct {
	Body string `json:"body"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirpCreate(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}
	authUUID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	decoder := json.NewDecoder(r.Body)
	chirpData := chirp{}
	err = decoder.Decode(&chirpData)

	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}
	if len(chirpData.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	cleanedText := helperBadWordsClean(chirpData.Body)

	// Now chirp is valid we can create it
	dbchirp, err := cfg.db.CreateChirp(r.Context(),
		database.CreateChirpParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Body:      cleanedText,
			UserID:    authUUID,
		})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	chirp := Chirp{
		ID:        dbchirp.ID,
		CreatedAt: dbchirp.CreatedAt,
		UpdatedAt: dbchirp.UpdatedAt,
		Body:      dbchirp.Body,
		UserID:    dbchirp.UserID,
	}
	respondWithJSON(w, http.StatusCreated, chirp)
}
