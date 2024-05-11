package main

import (
	"net/http"

	"github.com/Pranay-Pandey/rssagg/internal/auth"
	"github.com/Pranay-Pandey/rssagg/internal/database"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.AppUser)

func (apiCfg apiConfig)authMiddleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKEY(r.Header)
	
		if err != nil {
			respondWithError(w, 401, "Not valid api key, " + err.Error() )
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 401, "Invalid apiKey")
			return
		}

		handler(w, r, user)
	}
}