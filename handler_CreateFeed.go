package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.AppUser) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error decoding request body "+err.Error())
		return
	}
 
	feed, feed_err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if feed_err != nil {
		respondWithError(w, 400, "Error creating feed "+feed_err.Error())
		return
	}

	respondWithJSON(w, 201, DBFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, "Unable to access Feeds "+err.Error())
		return
	}

	feedList := []Feed{}
	for _, feed := range feeds {
		feedList = append(feedList, DBFeedToFeed(feed))
	}

	respondWithJSON(w, 200, feedList)
}
