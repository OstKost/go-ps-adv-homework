package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := w.Header().Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}

		header := w.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, PATCH")
			header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept, Authorization")
			header.Set("Access-Control-Max-Age", "86400")
			return
		}

		next.ServeHTTP(w, r)
	})
}
