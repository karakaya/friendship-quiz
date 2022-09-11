package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karakaya/friendship-quiz/pkg/db"
	"github.com/karakaya/friendship-quiz/pkg/repository"
	"github.com/karakaya/friendship-quiz/pkg/route"
	"github.com/karakaya/friendship-quiz/pkg/service"
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

	dbc, _ := db.ConnectDB()
	collection := dbc.Database("friendship_quiz").Collection("friendship_quiz")

	repo := repository.NewRepository(dbc, collection)
	serv := service.NewService(repo)

	a.routes(serv)

	a.gin.Run(fmt.Sprintf(":%d", a.port))
	// http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.gin.Handler())
}

func (a *App) routes(quizService service.QuizService) {
	route.RegisterRoutes(&a.gin.RouterGroup, quizService)
}
