package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/504BassSlapper/rss-agregator/helper"
	"github.com/504BassSlapper/rss-agregator/internal/database"
	"github.com/504BassSlapper/rss-agregator/models"
	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
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
		helper.RespondWithError(w, 400, fmt.Sprintf("could not create feed", err))
		return
	}
	helper.RespondWithJson(w, 201, models.DatabaseFeedToFeed(feed))

}

func (apiConfig *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		helper.RespondWithError(w, 404, fmt.Sprintf("feeds not found in database"))
	}
	fmt.Println(feeds)

	helper.RespondWithJson(w, 200, feeds)

}
