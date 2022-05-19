package controller

import "net/http"

type Con interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}