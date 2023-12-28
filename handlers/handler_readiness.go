package handlers

import (
	"net/http"

	"github.com/504BassSlapper/rss-agregator/helper"
)

func HandlerReadiness(w http.ResponseWriter, request *http.Request) {
	helper.RespondWithJson(w, 200, struct{}{})
}
