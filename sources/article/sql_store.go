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

func CreateSQLWithDB(db *sql.DB, ph sq.PlaceholderFormat) (*SQLStore, error) {
	return &SQLStore{db: db, ph: ph}, nil
}

func CreateSQLStore(driver, connString string, ph sq.PlaceholderFormat) (*SQLStore, error) {
	db, err := sql.Open(driver, connString)

	if err != nil {
		return nil, err
	}

	return CreateSQLWithDB(db, ph)
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

	idPredicate := sq.Eq{"id": id.String()}

	err = sq.
		Select("id", "title", "content", "created_at").
		From("articles").
		Where(idPredicate).
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

func (s *SQLStore) ListArticles(ctx context.Context) ([]ArticleBrief, error) {
	var ret []ArticleBrief
	var err error
	var rows *sql.Rows

	const initialSize = 25

	rows, err = sq.
		Select("id", "title", "created_at").
		From("articles").
		RunWith(s.db).PlaceholderFormat(s.ph).
		QueryContext(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	defer rows.Close()

	ret = make([]ArticleBrief, 0, 25)

	for rows.Next() {
		var b ArticleBrief
		if err := rows.Scan(&b.ID, &b.Title, &b.CreatedAt); err != nil {
			return nil, err
		}

		ret = append(ret, b)
	}

	return ret, nil
}
