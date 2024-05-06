package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, 200, struct{}{})
}
