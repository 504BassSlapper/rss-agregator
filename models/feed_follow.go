package models

import (
	"time"

	"github.com/504BassSlapper/rss-agregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"string"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowToFeedFollow(databaseFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        databaseFeedFollow.ID,
		CreatedAt: databaseFeedFollow.CreatedAt,
		UpdatedAt: databaseFeedFollow.UpdatedAt,
		UserId:    databaseFeedFollow.UserID,
		FeedId:    databaseFeedFollow.FeedID,
	}
}

func DataBaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) (feedFollows []FeedFollow) {
	for _, dbFeedFollodbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(dbFeedFollodbFeedFollow))
	}
	return
}
