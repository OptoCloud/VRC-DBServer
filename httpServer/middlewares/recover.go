package middlewares

import (
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			/*
				if err := recover(); err != nil {
					log.Printf("panic: %+v", err)
					helpers.WriteGenericError(w, "general", http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			*/
		}()

		next.ServeHTTP(w, r)
	})
}
