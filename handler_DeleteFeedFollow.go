package main

import (
	"net/http"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type SucessMessage struct {
	Message string `json:"message"`
}

func (apiCnf *apiConfig) handlerDeleteFollow(w http.ResponseWriter, r *http.Request, user database.AppUser) {
	feedFollowId := chi.URLParam(r, "feedFollowId")
	id, err := uuid.Parse(feedFollowId)

	if err != nil {
		respondWithError(w, 400, "Invalid request payload: "+err.Error())
	}

	delErr := apiCnf.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		ID:     id,
		UserID: user.ID,
	})

	if delErr != nil {
		respondWithError(w, 400, "Could not unfollow feed: "+err.Error())
	}

	respondWithJSON(w, 200, SucessMessage{
		Message: "Feed Follow unsubscribed successfully",
	})
}
