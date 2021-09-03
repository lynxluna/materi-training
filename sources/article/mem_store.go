package main

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type MemStore struct {
	articles map[uuid.UUID]Article

	lock *sync.RWMutex
}

var (
	ErrNilArticle      = errors.New("cannot save nil article")
	ErrNotImplemented  = errors.New("method not yet implemented")
	ErrArticleNotFound = errors.New("article not found")
)

// Membuat Memory Store
func CreateMemStore() *MemStore {
	return &MemStore{
		articles: make(map[uuid.UUID]Article),
		lock:     &sync.RWMutex{},
	}
}

// Fill With Items

func (s *MemStore) FillArticle(articles ...Article) {
	if len(articles) == 0 {
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	for _, a := range articles {
		s.articles[a.ID] = a
	}
}

// Menyimpan satu artikel
func (s *MemStore) SaveArticle(ctx context.Context, article Article) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if article.IsNil() {
		return ErrNilArticle
	}

	s.articles[article.ID] = article

	return nil
}

// Mencari satu artikel berdasarkan ID
func (s *MemStore) FindArticleByID(ctx context.Context, id uuid.UUID) (Article, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	article, ok := s.articles[id]

	if !ok {
		return Article{}, ErrArticleNotFound
	}

	return article, nil
}
