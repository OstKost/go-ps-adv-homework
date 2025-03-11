package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

type JWTData struct {
	Phone   string `json:"phone"`
	Session string `json:"session"`
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) SignToken(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone":   data.Phone,
		"session": data.Session,
	})
	signedToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JWT) ParseToken(tokenString string) (bool, *JWTData) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(err)
		return false, nil
	}
	jwtData := &JWTData{
		Phone:   claims["phone"].(string),
		Session: claims["session"].(string),
	}
	return token.Valid, jwtData
}
