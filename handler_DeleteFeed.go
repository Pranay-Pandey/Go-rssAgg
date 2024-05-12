package main

import (
	"net/http"

	"github.com/Pranay-Pandey/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCnf *apiConfig) handlerDeleteFeed(w http.ResponseWriter, r *http.Request, user database.AppUser) {
	feedString := chi.URLParam(r, "feedId")
	feedId, err := uuid.Parse(feedString)

	if err != nil {
		respondWithError(w, 400, "Cannot find the feed from Id, "+ err.Error())
		return
	}

	err = apiCnf.DB.DeleteFeed(r.Context(), database.DeleteFeedParams{
		ID : feedId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, "Cannot Delete the feed, "+ err.Error())
		return
	}

	respondWithJSON(w, 200, "Feed deleted successfully")
}