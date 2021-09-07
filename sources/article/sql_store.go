package main

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type SQLStore struct {
	db *sql.DB

	ph sq.PlaceholderFormat
}

func CreateSQLStore(driver, connString string, ph sq.PlaceholderFormat) (*SQLStore, error) {
	db, err := sql.Open(driver, connString)

	if err != nil {
		return nil, err
	}

	return &SQLStore{
		db: db,
		ph: ph,
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

	_, err = sq.
		Insert("articles").Columns("id", "title", "content", "created_at").
		SetMap(updateMap).
		PlaceholderFormat(s.ph).RunWith(s.db).ExecContext(ctx)

	if err == nil {
		return err
	}

	idPredicate := sq.Eq{"id": article.ID}

	_, err = sq.
		Update("articles").Where(idPredicate).
		SetMap(updateMap).
		PlaceholderFormat(s.ph).RunWith(s.db).ExecContext(ctx)

	return err
}

func (s *SQLStore) FindArticleByID(ctx context.Context, id uuid.UUID) (Article, error) {
	var article Article
	var err error

	idPredicate := sq.Eq{"id": article.ID}

	err = sq.
		Select("id", "title", "content", "created_at").Where(idPredicate).
		RunWith(s.db).PlaceholderFormat(s.ph).
		ScanContext(ctx,
			&article.ID,
			&article.Title,
			&article.Content,
			&article.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return Article{}, ErrArticleNotFound
		}
		return Article{}, err
	}

	return article, nil
}
