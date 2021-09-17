package main

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/google/go-cmp/cmp"
)

const (
	tableCreate = `CREATE TABLE IF NOT EXISTS articles (
  id uuid NOT NULL PRIMARY KEY,
  title VARCHAR(1024) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW());

	TRUNCATE TABLE articles;
	`

	pgConnStr = "dbname=articledbtest user=postgres password=postgres host=localhost sslmode=disable"
)

func setup(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", pgConnStr)
	require.NoError(t, err)

	_, err = db.Exec(tableCreate)
	require.NoError(t, err)

	return db
}

func cleanup(t *testing.T, db *sql.DB) {
	db.Exec("DROP TABLE articles;")
	db.Close()
}

func TestSQLStorePostgres(t *testing.T) {
	if testing.Short() {
		t.Skip() // <1>
	}
	ctx := context.Background()

	db := setup(t)
	defer cleanup(t, db)

	article, err := CreateArticle(validTitle, validContent)
	require.NoError(t, err)

	store, err := CreateSQLWithDB(db, squirrel.Dollar)
	require.NoError(t, err)

	err = store.SaveArticle(ctx, article) // <2>

	if assert.NoError(t, err) {
		a, err := store.FindArticleByID(ctx, article.ID) // <3>

		if assert.NoError(t, err) {
			assert.Equal(t, article.ID, a.ID)
			assert.Equal(t, article.Title, a.Title)
			assert.Equal(t, article.Content, a.Content)
			assert.GreaterOrEqual(t, 1*time.Minute, article.CreatedAt.Sub(a.CreatedAt))
		}
	}

	nonExistentID, _ := uuid.NewRandom()

	a, err := store.FindArticleByID(ctx, nonExistentID) // <4>

	assert.Equal(t, Article{}, a)
	assert.ErrorIs(t, ErrArticleNotFound, err)

	articleList := []ArticleBrief{
		{article.ID, article.Title, article.CreatedAt},
	}

	res, err := store.ListArticles(ctx)

	if assert.NoError(t, err) {
		assert.NotNil(t, res)
		if !assert.True(t, cmp.Equal(articleList, res)) {
			t.Log(cmp.Diff(articleList, res))
		}
	}
}
