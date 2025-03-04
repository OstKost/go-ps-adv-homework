package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		requestLogger := log.WithFields(log.Fields{
			"statusCode": wrapper.StatusCode,
			"method":     r.Method,
			"url":        r.URL.String(),
			"duration":   time.Since(start),
		})
		switch wrapper.StatusCode {
		case http.StatusOK:
			requestLogger.Info("OK")
		case http.StatusNotFound:
			requestLogger.Warn("not found")
		case http.StatusInternalServerError:
			requestLogger.Error("internal server error")
		case http.StatusUnauthorized:
			requestLogger.Warn("unauthorized")
		case http.StatusForbidden:
			requestLogger.Warn("request forbidden")
		}
	})
}
