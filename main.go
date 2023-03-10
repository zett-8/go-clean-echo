package main

import (
	database "github.com/zett-8/go-clean-echo/db"
	_ "github.com/zett-8/go-clean-echo/docs"
	"github.com/zett-8/go-clean-echo/handlers"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/middlewares"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/stores"
	"go.uber.org/zap"
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
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(GO_ENV == "development")
	if err != nil {
		logger.Fatal("failed to connect to the database", zap.Error(err))
	}
	defer db.Close()

	e := handlers.Echo()

	s := stores.New(db)
	ss := services.New(s)
	h := handlers.New(ss)

	jwtCheck, err := middlewares.JwtMiddleware()
	if err != nil {
		logger.Fatal("failed to set JWT middleware", zap.Error(err))
	}

	handlers.SetDefault(e)
	handlers.SetApi(e, h, jwtCheck)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8888"
	}

	logger.Fatal("failed to start server", zap.Error(e.Start(":"+PORT)))
}
