package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zett-8/go-echo-without-orm/db"
	"github.com/zett-8/go-echo-without-orm/handlers"
	"github.com/zett-8/go-echo-without-orm/services"
	"github.com/zett-8/go-echo-without-orm/store"
	"os"
)

var GO_ENV = os.Getenv("GO_ENV")

func main() {

	fmt.Println("GO_ENV:", GO_ENV)

	db, err := db.Load()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authorStore := store.NewAuthorStore(db)
	authorService := services.NewAuthorService(authorStore)
	handlers.NewAuthorHandler(e, authorService)

	handlers.Set(e)

	e.Logger.Fatal(e.Start(":8080"))
}
