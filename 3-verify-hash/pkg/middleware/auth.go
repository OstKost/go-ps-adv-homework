package middleware

import (
	"fmt"
	"go-ps-adv-homework/pkg/jwt"
	"net/http"
	"os"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authorization, "Bearer ")
		err := jwt.NewJWT(os.Getenv("SECRET")).VerifyToken(token)
		if err != nil {
			fmt.Println(err)
		}
		next.ServeHTTP(w, r)
	})
}
