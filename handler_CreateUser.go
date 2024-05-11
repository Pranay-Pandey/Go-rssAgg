package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	respondWithJSON(w, 200, DBUserToUser(user))
}


func (apiCfg *apiConfig) handleGetUserByAPIKEY(w http.ResponseWriter, r *http.Request, user database.AppUser) {
	respondWithJSON(w, 200, DBUserToUser(user))
}

