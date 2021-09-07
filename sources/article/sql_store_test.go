package main

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	tableCreate = `CREATE TABLE IF NOT EXISTS articles (
  id uuid NOT NULL PRIMARY KEY,
  title VARCHAR(1024) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW());`

	pgConnStr = "dbname=articledbtest user=postgres password=postgres host=localhost sslmode=disable"
)

func TestSQLStorePostgres(t *testing.T) {

	ctx := context.Background()

	db, err := sql.Open("postgres", pgConnStr)
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(tableCreate)
	require.NoError(t, err)

	article, err := CreateArticle(validTitle, validContent)
	require.NoError(t, err)

	store, err := CreateSQLWithDB(db, squirrel.Dollar)
	require.NoError(t, err)

	err = store.SaveArticle(ctx, article)

	if assert.NoError(t, err) {
		a, err := store.FindArticleByID(ctx, article.ID)

		if assert.NoError(t, err) {
			assert.Equal(t, article, a)
		}
	}
}
