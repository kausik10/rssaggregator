package main

import (
	"fmt"
	"net/http"

	"github.com/kausik10/rssaggregator/internal/auth"
	"github.com/kausik10/rssaggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuthHandler(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondError(w, 403, fmt.Sprintf("Error getting api key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondError(w, 400, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
