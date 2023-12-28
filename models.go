package main

import (
	"time"

	"github.com/504BassSlapper/rss-agregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Apikey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"string"`
	UserId    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(databaseFeed database.Feed) Feed {
	return Feed{
		ID:        databaseFeed.ID,
		Name:      databaseFeed.Name,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
		Url:       databaseFeed.Url,
		UserId:    databaseFeed.UserID,
	}
}

func databaseUserToModelUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

func databaseUserToModelUserWithApiKey(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Apikey:    dbUser.ApiKey,
	}
}
