package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPServer struct {
	port uint16
	host string

	router *chi.Mux

	uc *ArticleUseCase
}

func NewHTTPServer(options ...func(*HTTPServer) error) (*HTTPServer, error) {
	uc, err := NewArticleUseCase(CreateMemStore())
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	httpServer := &HTTPServer{
		host:   "127.0.0.1",
		port:   8000,
		router: r,
		uc:     uc,
	}

	httpServer.setupRoute()

	if len(options) == 0 {
		return httpServer, nil
	}

	for _, opt := range options {
		if err := opt(httpServer); err != nil {
			return nil, err
		}
	}

	return httpServer, nil
}

func (s *HTTPServer) setupRoute() {
	r := s.router

	r.Post("/articles", s.NewArticleHandler)
	r.Put("/articles/{articleID}", s.EditArticleHandler)
}

func (s *HTTPServer) Start() {
	listen := fmt.Sprintf("%s:%d", s.host, s.port)

	http.ListenAndServe(listen, s.router)
}

// wrapError

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	wrapper := struct {
		Message string `json:"message"`
	}{Message: err.Error()}
	json.NewEncoder(w).Encode(wrapper)
}

func (s *HTTPServer) NewArticleHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, ErrNotImplemented)
}

func (s *HTTPServer) EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, ErrNotImplemented)
}
