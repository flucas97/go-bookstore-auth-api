package app

import (
	"github.com/flucas97/bookstore/auth-api/src/http"
	"github.com/flucas97/bookstore/auth-api/src/repository/db"
	"github.com/flucas97/bookstore/auth-api/src/repository/rest"
	"github.com/flucas97/bookstore/auth-api/src/service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	usersRepository := rest.NewRepository()
	atService := service.NewService(usersRepository, dbRepository) // to use service I need to pass a repository
	atHandler := http.NewAccessTokenHandler(atService)             // to use controller I need to pass a service
	// atHandler := http.NewAccessTokenHandler(service.NewService(db.NewRepository()))

	MapURL(atHandler)

	router.Run(":8080")
}
