package main

import (
	"fmt"
	database "github.com/zett-8/go-clean-echo/db"
	_ "github.com/zett-8/go-clean-echo/docs"
	"github.com/zett-8/go-clean-echo/handlers"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/stores"
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

	db, err := database.New(GO_ENV == "development")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := handlers.Echo()

	s := stores.New(db)
	ss := services.New(s)
	h := handlers.New(ss)

	handlers.Set(e, h)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8888"
	}

	log.Fatal(e.Start(":" + PORT))
}
