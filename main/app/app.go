package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karakaya/friendship-quiz/pkg/route"
)

type App struct {
	gin    *gin.Engine
	logger log.Logger
	port   int
}

func NewApp(logger log.Logger, port int) *App {
	return &App{
		gin:    gin.Default(),
		logger: logger,
		port:   port,
	}
}

func (a *App) Run() {
	a.routes()
	http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.gin.Handler())
}

func (a *App) routes() {
	route.RegisterRoutes(&a.gin.RouterGroup)
}