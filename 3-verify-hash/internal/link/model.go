package link

import (
	"gorm.io/gorm"
	"math/rand"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"unique_index"`
}

type CreateLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateLinkRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash" validate:"omitempty,len=10"`
}

func NewLink(url string) *Link {
	link := &Link{Url: url}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	const hashLen = 10
	link.Hash = RandStringRunes(hashLen)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
