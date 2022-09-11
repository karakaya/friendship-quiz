package route

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karakaya/friendship-quiz/pkg/request"
	"github.com/karakaya/friendship-quiz/pkg/service"
)

func RegisterRoutes(g *gin.RouterGroup, service service.QuizService) {

	res := resource{
		service: service,
	}

	g.GET("/{id}", res.get)
	g.POST("/", res.create)

}

type resource struct {
	service service.QuizService
}

func (r resource) create(g *gin.Context) {
	var creqteQuizRequest request.CreateQuizRequest
	if err := g.BindJSON(&creqteQuizRequest); err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	r.service.Create(context.Background(), creqteQuizRequest)
	g.JSON(200, "")
}

func (r resource) get(g *gin.Context) {
	g.JSON(200, "")
}
