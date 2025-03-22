package jwt_test

import (
	"go-ps-adv-homework/pkg/jwt"
	"testing"
)

func TestJWTSign(t *testing.T) {
	const email = "test@test.test"
	jwtService := jwt.NewJWT("8F2AB5438C3D975AB2D106BAECE75477")
	token, err := jwtService.SignToken(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, tokenData := jwtService.ParseToken(token)
	if !isValid {
		t.Fatal("Expected token is valid, but got false")
	}
	if tokenData.Email != email {
		t.Fatalf("Expected %s, got %s", email, tokenData.Email)
	}
}
