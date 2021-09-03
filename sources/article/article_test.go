package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	validContent = `Sollicitudin ac orci phasellus egestas tellus. Ultricies mi quis hendrerit dolor
  magna eget est lorem ipsum. Et netus et malesuada fames ac. Euismod quis viverra nibh cras
  pulvinar mattis nunc sed blandit. Aliquam vestibulum morbi blandit cursus risus at. Amet risus
  nullam eget felis eget nunc lobortis. Amet volutpat consequat mauris nunc congue nisi vitae. Sem
  viverra aliquet eget sit amet tellus cras adipiscing enim. Aliquam ultrices sagittis orci a
  scelerisque purus semper eget duis. Interdum velit laoreet id donec ultrices tincidunt.
  Sollicitudin aliquam ultrices sagittis orci a. Aliquet eget sit amet tellus. Quis enim lobortis
  scelerisque fermentum dui faucibus. Dolor sit amet consectetur adipiscing elit ut. Vulputate enim
  nulla aliquet porttitor lacus.`

	validTitle = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
  incididunt ut labore et dolore magna aliqua.`

	longTitle = `et netus et malesuada fames ac turpis egestas integer eget aliquet nibh praesent
  tristique magna sit amet purus gravida quis blandit turpis cursus in hac habitasse platea dictumst
  quisque sagittis purus sit amet volutpat consequat mauris nunc congue nisi vitae suscipit tellus
  mauris a diam maecenas sed enim ut sem viverra aliquet eget sit amet tellus cras adipiscing enim
  eu turpis egestas pretium aenean pharetra magna ac placerat vestibulum lectus mauris ultrices eros
  in cursus turpis massa tincidunt dui ut ornare lectus sit amet est placerat in egestas erat
  imperdiet sed euismod nisi porta lorem mollis aliquam ut porttitor leo`
)

func TestCreateArticle(t *testing.T) {
	// article yang dianggap nil (karena bukan pointer)
	var nilArticle Article
	require.True(t, nilArticle.IsNil())

	// id yang pasti dianggap valid
	id := uuid.MustParse("836f6aa2-ed56-437a-aee1-eff92cf4ee4d")

	validArticle, err := createArticleWithID(id, validTitle, validContent)
	require.NoError(t, err)

	// tabel test
	tests := []struct {
		// nama test
		Name string
		// input
		Title   string
		Content string
		// output
		Result Article
		Err    error
	}{
		{"EmptyTitleContent", "", "", nilArticle, ErrEmptyTitle},
		{"EmptyTitleOnly", "", validContent, nilArticle, ErrEmptyTitle},
		{"EmptyContent", validTitle, "", nilArticle, ErrEmptyContent},
		{"ShortTitle", "short", validContent, nilArticle, ErrTitleTooShort},
		{"ShortContent", validTitle, "short", nilArticle, ErrContentTooShort},
		{"TooLongTitle", longTitle, validContent, nilArticle, ErrTitleTooLong},
		{"ValidArticle", validTitle, validContent, validArticle, nil},
	}

	// Test dijalankan satu-satu dari tabel
	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			article, err := createArticleWithID(id, item.Title, item.Content)
			assert.Equal(t, item.Err, err)

			if item.Err == nil {
				assert.Equal(t, item.Result.ID, article.ID)
				assert.Equal(t, item.Result.Title, article.Title)
				assert.Equal(t, item.Result.Content, article.Content)
				assert.LessOrEqual(t, time.Now().Sub(article.CreatedAt), 5*time.Second)
				assert.False(t, item.Result.IsNil())
			}
		})
	}
}

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		Name  string
		Title string
		Err   error
	}{
		{"Empty", "", ErrEmptyTitle},
		{"TooShort", "too-short", ErrTitleTooShort},
		{"TooLong", longTitle, ErrTitleTooLong},
		{"Valid", validTitle, nil},
	}

	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			assert.Equal(t, item.Err, validateTitle(item.Title))
		})
	}
}

func TestValidateContent(t *testing.T) {
	tests := []struct {
		Name string

		Content string
		Err     error
	}{
		{"Empty", "", ErrEmptyContent},
		{"TooShort", "short-content", ErrContentTooShort},
		{"Valid", validContent, nil},
	}

	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			assert.Equal(t, item.Err, validateContent(item.Content))
		})
	}
}
