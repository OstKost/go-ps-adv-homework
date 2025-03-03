package hashes

import (
	"github.com/google/uuid"
	"sync"
)

type Hashes struct {
	m sync.Map
}

func New() *Hashes {
	return &Hashes{}
}

func (h *Hashes) Add() string {
	hash := uuid.New().String()
	h.m.Store(hash, true)
	return hash
}

func (h *Hashes) Exists(hash string) bool {
	_, ok := h.m.Load(hash)
	return ok
}

func (h *Hashes) Remove(hash string) {
	h.m.Delete(hash)
}

func (h *Hashes) GetAll() []string {
	var hashes []string
	h.m.Range(func(key, value interface{}) bool {
		if hash, ok := key.(string); ok {
			hashes = append(hashes, hash)
		}
		return true
	})
	return hashes
}
