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

func (apiConfig *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, "Could not parse parameters")
		return
	}

	feeds_follow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintln("could not create feed follow: ", err))
		return
	}
	helper.RespondWithJson(w, 201, models.DatabaseFeedFollowToFeedFollow(feeds_follow))

}
