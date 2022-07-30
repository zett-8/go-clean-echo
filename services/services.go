package services

import "github.com/zett-8/go-clean-echo/stores"

type Services struct {
	AuthorService
	BookService
}

func New(s *stores.Stores) *Services {
	return &Services{
		AuthorService: &AuthorServiceContext{s.AuthorStore},
		BookService:   &BookServiceContext{s.BookStore},
	}
}
