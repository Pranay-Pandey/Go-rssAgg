package main

import (
	"fmt"
	"net/http"

)

func (apiCfg *apiConfig)handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.Getusers(r.Context())
	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to get users: %v", err))
		return
	}

	convertedUsers := make([]User, len(users))
	for i, user := range users {
		convertedUsers[i] = DBUserToUser(user)
	}

	respondWithJSON(w, 200, convertedUsers)
}