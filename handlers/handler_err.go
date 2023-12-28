package handlers

import (
	"net/http"

	"github.com/504BassSlapper/rss-agregator/helper"
)

func HandleErr(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithError(w, 400, "Something went wrong ")
}
