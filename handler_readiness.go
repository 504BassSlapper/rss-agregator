package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, request *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
