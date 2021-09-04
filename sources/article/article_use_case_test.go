package main

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateArticleUseCase(t *testing.T) {
	ctx := context.Background()
	mem := CreateMemStore()

	uc, err := NewArticleUseCase(mem)
	require.NoError(t, err)

	// tabel test
	tests := []struct {
		// nama test
		Name string
		// input
		Title   string
		Content string
		// output
		Err error
	}{
		{"EmptyTitleContent", "", "", ErrEmptyTitle},
		{"EmptyTitleOnly", "", validContent, ErrEmptyTitle},
		{"EmptyContent", validTitle, "", ErrEmptyContent},
		{"ShortTitle", "short", validContent, ErrTitleTooShort},
		{"ShortContent", validTitle, "short", ErrContentTooShort},
		{"TooLongTitle", longTitle, validContent, ErrTitleTooLong},
		{"ValidArticle", validTitle, validContent, nil},
	}

	// Test dijalankan satu-satu dari tabel
	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			article, err := uc.CreateArticle(ctx, item.Title, item.Content)
			assert.Equal(t, item.Err, err)

			if err != nil {
				assert.True(t, article.IsNil())
				return
			}

			result, err := mem.FindArticleByID(ctx, article.ID)

			if assert.NoError(t, err) {
				assert.Equal(t, article, result)
			}
		})
	}
}

func TestEditArticleUseCase(t *testing.T) {
	ctx := context.Background()
	mem := CreateMemStore()

	a, err := CreateArticle(validTitle, validContent)
	require.NoError(t, err)

	mem.FillArticle(a)

	uc, err := NewArticleUseCase(mem)
	require.NoError(t, err)

	existID := a.ID
	nonExistentID := uuid.MustParse("30339469-935b-4ab5-8816-d8a47450fe5f")

	replaceContent := validContent[20:]
	// tabel test
	tests := []struct {
		// nama test
		Name string
		// input
		ID      uuid.UUID
		Title   string
		Content string
		// output
		Err error
	}{
		{"NonExistentID", nonExistentID, validTitle, replaceContent, ErrArticleNotFound},
		{"EmptyTitleContent", existID, "", "", ErrEmptyTitle},
		{"EmptyTitleOnly", existID, "", validContent, ErrEmptyTitle},
		{"EmptyContent", existID, validTitle, "", ErrEmptyContent},
		{"ShortTitle", existID, "short", validContent, ErrTitleTooShort},
		{"ShortContent", existID, validTitle, "short", ErrContentTooShort},
		{"TooLongTitle", existID, longTitle, validContent, ErrTitleTooLong},
		{"ValidArticle", existID, validTitle, replaceContent, nil},
	}

	// Test dijalankan satu-satu dari tabel
	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			err := uc.EditArticle(ctx, item.ID, item.Title, item.Content)
			assert.Equal(t, item.Err, err)

			if err != nil {
				return
			}

			a, err := mem.FindArticleByID(ctx, item.ID)

			if !assert.NoError(t, err) {
				return
			}

			assert.Equal(t, item.Title, a.Title)
			assert.Equal(t, item.Content, a.Content)
		})
	}
}
