package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustMarshalJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

var (
	invalidReqPayloadErrJSON = wrapError(ErrInvalidRequestPayload)
	emptyTitleErrJSON        = wrapError(ErrEmptyTitle)
	titleTooShortErrJSON     = wrapError(ErrTitleTooShort)
	contentTooShortErrJSON   = wrapError(ErrContentTooShort)
	titleTooLongErrJSON      = wrapError(ErrTitleTooLong)
)

func TestCreateArticleHandler(t *testing.T) {
	type testPayload struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	validPayload := testPayload{validTitle, validContent}
	validPayloadJSON := mustMarshalJSON(validPayload)
	validArticle, _ := CreateArticle(validTitle, validContent)
	validArticleJSON := mustMarshalJSON(validArticle)

	tests := []struct {
		name    string
		reqJSON []byte

		status int
		ret    []byte
	}{
		{"NilRequest", nil,
			http.StatusBadRequest, invalidReqPayloadErrJSON},
		{"BlankRequest", []byte(""),
			http.StatusBadRequest, invalidReqPayloadErrJSON},
		{"BlankJSONRequest", []byte("{}"),
			http.StatusUnprocessableEntity, emptyTitleErrJSON},
		{"WithoutMatchingKeys", []byte(`{"red":"yes"}`),
			http.StatusUnprocessableEntity, emptyTitleErrJSON},
		{"NilTitleAndContent", []byte(`{"title":nil,"content":nil}`),
			http.StatusBadRequest, invalidReqPayloadErrJSON},
		{"ShortTitle", mustMarshalJSON(testPayload{"short", validContent}),
			http.StatusUnprocessableEntity, titleTooShortErrJSON},
		{"ShortContent", mustMarshalJSON(testPayload{validTitle, "short"}),
			http.StatusUnprocessableEntity, contentTooShortErrJSON},
		{"LongTitle", mustMarshalJSON(testPayload{longTitle, validContent}),
			http.StatusUnprocessableEntity, titleTooLongErrJSON},
		{"ValidArticle", validPayloadJSON,
			http.StatusCreated, validArticleJSON},
	}

	s, err := NewHTTPServer()
	require.NoError(t, err)

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(item.reqJSON))
			req.Header.Set("content-type", "application/json")
			rr := httptest.NewRecorder()

			s.NewArticleHandler(rr, req)

			res := rr.Result()
			defer res.Body.Close()

			assert.Equal(t, item.status, res.StatusCode)

			var buf bytes.Buffer
			n, err := buf.ReadFrom(res.Body)
			require.NoError(t, err)
			assert.NotZero(t, n)

			isErrorCode := res.StatusCode/100 > 2

			if isErrorCode {
				assert.JSONEq(t, string(item.ret), buf.String())
				return
			}

			var ret struct {
				ID        string `json:"id"`
				CreatedAt string `json:"created_at"`
			}

			err = json.NewDecoder(&buf).Decode(&ret)

			if !assert.NoError(t, err) {
				return
			}

			_, err = uuid.Parse(ret.ID)
			assert.NoError(t, err)
			d, err := time.Parse(time.RFC3339, ret.CreatedAt)
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, 4*time.Minute, time.Now().Sub(d))

		})
	}
}
