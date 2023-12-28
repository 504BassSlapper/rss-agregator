package handlers

import (
	"net/http"

	"github.com/504BassSlapper/rss-agregator/auth"
	"github.com/504BassSlapper/rss-agregator/helper"
	"github.com/504BassSlapper/rss-agregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddleWareAuth(handler authedHandler) (handlerFunc http.HandlerFunc) {

	handlerFunc = func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(&r.Header)
		if err != nil {
			helper.RespondWithError(w, 403, err.Error())
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			helper.RespondWithError(w, 404, err.Error())
		}
		handler(w, r, user)
	}

	return
}
