package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	database "github.com/zett-8/go-echo-without-orm/db"
	"github.com/zett-8/go-echo-without-orm/handlers"
	"github.com/zett-8/go-echo-without-orm/services"
	"github.com/zett-8/go-echo-without-orm/store"
	"os"
)

var GO_ENV = os.Getenv("GO_ENV")

func main() {

	fmt.Println("GO_ENV:", GO_ENV)

	db, err := database.Load()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authorStore := store.NewAuthorStore(db)
	bookStore := store.NewBooksStore(db)

	authorService := services.NewAuthorService(authorStore)
	bookService := services.NewBookService(bookStore)

	handlers.NewAuthorHandler(e, authorService)
	handlers.NewBookHandler(e, bookService)
	handlers.NewIndexHandler(e)

	e.Logger.Fatal(e.Start(":8080"))
}
