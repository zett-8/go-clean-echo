package services

import "github.com/zett-8/go-clean-echo/stores"

type Services struct {
	Author AuthorService
	Book   BookService
}

func New(s *stores.Stores) *Services {
	return &Services{
		Author: &authorService{stores: s},
		Book:   &bookService{stores: s},
	}
}
