package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	echoSwagger "github.com/swaggo/echo-swagger"
	database "github.com/zett-8/go-clean-echo/db"
	_ "github.com/zett-8/go-clean-echo/docs"
	"github.com/zett-8/go-clean-echo/handlers"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/store"
	"log"
	"os"
)

var GO_ENV = os.Getenv("GO_ENV")

// @title Go clean echo API v1
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8888
// @BasePath /
// @schemes http
func main() {
	fmt.Println("GO_ENV:", GO_ENV)

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := handlers.New()
	v1 := e.Group("/api/v1")

	authorStore := store.NewAuthorStore(db)
	bookStore := store.NewBooksStore(db)

	authorService := services.NewAuthorService(authorStore)
	bookService := services.NewBookService(bookStore)

	handlers.NewAuthorHandler(v1, authorService)
	handlers.NewBookHandler(v1, bookService)
	handlers.NewIndexHandler(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8888"
	}

	log.Fatal(e.Start(":" + PORT))
}
