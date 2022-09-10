package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karakaya/friendship-quiz/pkg/route"
)

type App struct {
	gin    *gin.Engine
	logger log.Logger
	port   int
}

func NewApp(logger log.Logger) *App {

	return &App{
		gin:    gin.Default(),
		logger: logger,
		port:   8080,
	}
}

func (a *App) Run() {
	a.routes()
	a.gin.Run(fmt.Sprintf(":%d", a.port))
	// http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.gin.Handler())
}

func (a *App) routes() {
	route.RegisterRoutes(&a.gin.RouterGroup)
}
