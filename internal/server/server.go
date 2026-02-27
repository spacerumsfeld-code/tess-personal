package server

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
	)

	var handler http.Handler = mux
	handler = corsMiddleware(handler)
	return handler
}
