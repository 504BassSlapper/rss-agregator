package main

import (
	"net/http"

	"github.com/504BassSlapper/rss-agregator/auth"
	"github.com/504BassSlapper/rss-agregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middleWareAuth(handler authedHandler) (handlerFunc http.HandlerFunc) {

	handlerFunc = func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(&r.Header)
		if err != nil {
			respondWithError(w, 403, err.Error())
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, err.Error())
		}
		handler(w, r, user)
	}

	return
}
