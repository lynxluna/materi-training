package main

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type SQLStore struct {
	db *sql.DB
}

func CreateSQLStore(driver, connString string) (*SQLStore, error) {
	db, err := sql.Open(driver, connString)

	if err != nil {
		return nil, err
	}

	return &SQLStore{
		db: db,
	}, nil
}

func (s *SQLStore) SaveArticle(ctx context.Context, article Article) error {
	var err error

	updateMap := map[string]interface{}{
		"id":         article.ID,
		"title":      article.Title,
		"content":    article.Content,
		"created_at": article.CreatedAt,
	}

	_, err = squirrel.
		Insert("articles").Columns("id", "title", "content", "created_at").
		SetMap(updateMap).RunWith(s.db).ExecContext(ctx)

	if err == nil {
		return err
	}

	idPredicate := squirrel.Eq{"id": article.ID}

	_, err = squirrel.
		Update("articles").Where(idPredicate).
		SetMap(updateMap).RunWith(s.db).ExecContext(ctx)

	return err
}

func (s *SQLStore) FindArticleByID(ctx context.Context, id uuid.UUID) (Article, error) {
	return Article{}, ErrNotImplemented
}
