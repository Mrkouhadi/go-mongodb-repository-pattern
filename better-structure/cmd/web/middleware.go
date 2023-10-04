package main

import "net/http"

// LoadSession loads and save session on every request
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
