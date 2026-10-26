package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type email struct {
	Email string `json:"email"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userEmail := email{}
	err := decoder.Decode(&userEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}
	dbuser, err := cfg.db.CreateUser(r.Context(), userEmail.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	user := User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Email:     dbuser.Email,
	}

	respondWithJSON(w, http.StatusCreated, user)
}
