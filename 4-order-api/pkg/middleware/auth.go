package middleware

import (
	"context"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhoneKey   key = "PhoneKey"
	ContextSessionKey key = "SessionKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			writeUnauthorized(w)
			return
		}
		token := strings.TrimPrefix(authorization, "Bearer ")
		isValid, jwtData := jwt.NewJWT(config.Auth.Secret).ParseToken(token)
		if !isValid {
			writeUnauthorized(w)
			return
		}
		ctx := context.Background()
		ctx = context.WithValue(ctx, ContextPhoneKey, jwtData.Phone)
		ctx = context.WithValue(ctx, ContextSessionKey, jwtData.Session)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
