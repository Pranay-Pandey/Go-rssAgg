package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFollow(w http.ResponseWriter, r *http.Request, user database.AppUser){
	type parameters struct{
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	
	if err != nil {
		respondWithError(w, 400, "Error decoding request body, " + err.Error())
		return
	}

	feedFollow, feedFollowErr := apiCfg.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID	:		uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		FeedID    : params.FeedID,
		UserID    : user.ID,
	})

	if feedFollowErr != nil {
		respondWithError(w, 400, "Error Creating Follow on provided Feed , " + feedFollowErr.Error())
		return 
	}

	respondWithJSON(w, 201, DBFollowToFollow(feedFollow))
}	

func (apiCfg *apiConfig) handlerGetFollows(w http.ResponseWriter, r *http.Request, user database.AppUser){
	follows, err := apiCfg.DB.GetFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, "Error Creating Follow on provided Feed , " + err.Error())
		return 
	}

	respondWithJSON(w, 200, DBFollowsToFollows(follows))
}