package middlewares

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AddHeadersMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, r)
		})
	}
}
