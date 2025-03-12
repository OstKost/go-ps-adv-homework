package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Logger возвращает middleware для логирования HTTP запросов.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		logRequest(wrapper, r, start)
	})
}

// logRequest логирует информацию о запросе.
func logRequest(w *WrapperWriter, r *http.Request, start time.Time) {
	requestLogger := log.WithFields(log.Fields{
		"statusCode": w.StatusCode,
		"method":     r.Method,
		"url":        r.URL.String(),
		"duration":   time.Since(start),
	})

	switch w.StatusCode {
	case http.StatusOK:
		requestLogger.Info("OK")
	case http.StatusNotFound:
		requestLogger.Warn("Not found")
	case http.StatusInternalServerError:
		requestLogger.Error("Internal server error")
	case http.StatusUnauthorized:
		requestLogger.Warn("Unauthorized")
	case http.StatusForbidden:
		requestLogger.Warn("Request forbidden")
	default:
		requestLogger.Info("Request processed")
	}
}
