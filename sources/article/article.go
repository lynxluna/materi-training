package main

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Article struct {
	ID      uuid.UUID
	Title   string
	Content string

	CreatedAt time.Time
}

func (a Article) IsNil() bool {
	return a.ID == uuid.Nil || (len(a.Title) == 0 && len(a.Content) == 0)
}

var (
	ErrEmptyTitle      = errors.New("title is empty")
	ErrEmptyContent    = errors.New("content is empty")
	ErrTitleTooShort   = errors.New("title too short")
	ErrTitleTooLong    = errors.New("title too long")
	ErrContentTooShort = errors.New("content too short")
)

func validateTitle(title string) error {
	const minTitleLength = 10
	const maxTitleLength = 500

	runeCount := utf8.RuneCountInString(title)

	if runeCount == 0 {
		return ErrEmptyTitle
	}

	if runeCount < minTitleLength {
		return ErrTitleTooShort
	}

	if runeCount > maxTitleLength {
		return ErrTitleTooLong
	}

	return nil
}
func validateContent(content string) error {
	const minContentLength = 200

	runeCount := utf8.RuneCountInString(content)

	if runeCount == 0 {
		return ErrEmptyContent
	}

	if runeCount < minContentLength {
		return ErrContentTooShort
	}

	return nil
}
func createArticleWithID(id uuid.UUID, title, content string) (Article, error) {
	var newArticle Article

	if err := validateTitle(title); err != nil {
		return Article{}, err
	}

	if err := validateContent(content); err != nil {
		return Article{}, err
	}

	newArticle = Article{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return newArticle, nil
}

func CreateArticle(title, content string) (Article, error) {
	newId, err := uuid.NewRandom()

	if err != nil {
		return Article{}, err
	}
	return createArticleWithID(newId, title, content)
}

func (a *Article) EditArticle(title, content string) error {
	if err := validateTitle(title); err != nil {
		return err
	}

	if err := validateContent(content); err != nil {
		return err
	}

	a.Title = title
	a.Content = content

	return nil
}
