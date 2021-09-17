package main

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type ArticleUseCase struct {
	store ArticleFinderSaver
}

var (
	ErrNilStore = errors.New("store cannot be nil")
)

func NewArticleUseCase(store ArticleFinderSaver) (*ArticleUseCase, error) {
	if store == nil {
		return nil, ErrNilStore
	}

	return &ArticleUseCase{store: store}, nil
}

func (uc *ArticleUseCase) CreateArticle(ctx context.Context, title, content string) (Article, error) {
	newArticle, err := CreateArticle(title, content)

	if err != nil {
		return Article{}, err
	}

	err = uc.store.SaveArticle(ctx, newArticle)

	if err != nil {
		return Article{}, err
	}

	return newArticle, nil
}

func (uc *ArticleUseCase) EditArticle(ctx context.Context, id uuid.UUID, newTitle, newContent string) error {
	article, err := uc.store.FindArticleByID(ctx, id)
	if err != nil {
		return err
	}

	if err = article.EditArticle(newTitle, newContent); err != nil {
		return err
	}

	return uc.store.SaveArticle(ctx, article)
}
