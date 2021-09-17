package main

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemStore(t *testing.T) {
	memStore := CreateMemStore() // <1>

	article, err := CreateArticle(validTitle, validContent)

	require.NoError(t, err)

	ctx := context.Background()

	err = memStore.SaveArticle(ctx, article) // <2>

	if assert.NoError(t, err) {
		a, err := memStore.FindArticleByID(ctx, article.ID) // <3>

		if assert.NoError(t, err) {
			assert.Equal(t, article, a)
		}
	}

	nonExistentID, _ := uuid.NewRandom()

	a, err := memStore.FindArticleByID(ctx, nonExistentID) // <4>

	assert.Equal(t, Article{}, a)
	assert.ErrorIs(t, ErrArticleNotFound, err)

	articleList := []ArticleBrief{
		{article.ID, article.Title, article.CreatedAt},
	}

	res, err := memStore.ListArticles(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	if !assert.True(t, cmp.Equal(articleList, res)) {
		t.Log(cmp.Diff(articleList, res))
	}
}
