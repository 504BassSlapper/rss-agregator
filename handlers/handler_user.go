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
	_ "github.com/lib/pq"
)

func (apiConfig *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing json:", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Could not create user: %v", err))
		return
	}
	helper.RespondWithJson(w, 201, models.DatabaseUserToModelUser(user))
}

// get user by apiKey
func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	helper.RespondWithJson(w, 200, models.DatabaseUserToModelUserWithApiKey(user))
}
