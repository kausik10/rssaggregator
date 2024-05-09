package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/kausik10/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}
	respondJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feed_follow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondError(w, 400, fmt.Sprintf("Error getting feed follows: %v", err))
		return
	}

	respondJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follow))
}
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")

	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		respondError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondError(w, 400, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	// For simplicity just return empty JSON
	respondJSON(w, 200, struct{}{})
}
