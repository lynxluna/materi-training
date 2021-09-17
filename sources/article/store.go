package main

import (
	"context"

	"github.com/google/uuid"
)

type ArticleFinder interface {
	FindArticleByID(ctx context.Context, id uuid.UUID) (Article, error)
}

type ArticleSaver interface {
	SaveArticle(ctx context.Context, article Article) error
}

type ArticleFinderSaver interface {
	ArticleFinder
	ArticleSaver
}

// Read Models

type ArticleLister interface {
	ListArticles(ctx context.Context) ([]ArticleBrief, error)
}

type ArticleReader interface {
	ArticleFinder
	ArticleLister
}
