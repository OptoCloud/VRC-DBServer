package middlewares

import (
	"log"
	"net/http"
	"time"
	"vrcdb/core"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		log.Printf("[%s (%s)] %s %s %v\n", r.Header.Get(core.HttpHeaderIpAddressKey), r.Header.Get(core.HttpHeaderIpCountryKey), r.Method, r.URL.Path, time.Since(start))
	})
}
