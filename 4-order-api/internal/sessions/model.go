package sessions

import (
	"github.com/google/uuid"
	"go-ps-adv-homework/internal/users"
	"gorm.io/gorm"
	"math/rand"
)

type Session struct {
	gorm.Model
	UserID  uint       `json:"userId"`
	User    users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Phone   string     `json:"phone" gorm:"uniqueIndex"`
	Session string     `json:"session" gorm:"uniqueIndex"`
	Code    string     `json:"code"`
}

func NewSession(userId uint, phone, code string) (*Session, error) {
	session, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	if code == "" {
		code = generateCode(4)
	}
	return &Session{
		UserID:  userId,
		Phone:   phone,
		Session: session.String(),
		Code:    code,
	}, nil
}

func generateCode(n int) string {
	var letterRunes = []rune("1234567890")
	code := make([]rune, n)
	for i := range code {
		code[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(code)
}
