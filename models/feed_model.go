package models

import (
	"time"

	"github.com/504BassSlapper/rss-agregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"string"`
	UserId    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(databaseFeed database.Feed) Feed {
	return Feed{
		ID:        databaseFeed.ID,
		Name:      databaseFeed.Name,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
		Url:       databaseFeed.Url,
		UserId:    databaseFeed.UserID,
	}
}
