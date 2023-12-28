package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/504BassSlapper/rss-agregator/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return

	}
	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		UserID:    user.ID,
		Url:       params.Url,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not create feed", err))
		return
	}
	respondWithJson(w, 201, databaseFeedToFeed(feed))

}
