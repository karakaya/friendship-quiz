package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karakaya/friendship-quiz/pkg/repository"
	"github.com/karakaya/friendship-quiz/pkg/request"
)

func RegisterRoutes(g *gin.RouterGroup) {
	res := resource{}
	g.GET("/{id}", res.get)
	g.POST("/", res.create)

}

type resource struct {
	repo repository.Repository
}

func (r resource) create(g *gin.Context) {
	var creqteQuizRequest request.CreateQuizRequest
	if err := g.BindJSON(&creqteQuizRequest); err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(creqteQuizRequest)
	g.JSON(200, "")
}

func (r resource) get(g *gin.Context) {
	g.JSON(200, "")
}
