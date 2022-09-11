package route

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karakaya/friendship-quiz/pkg/request"
	"github.com/karakaya/friendship-quiz/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	uid, _ := r.service.Create(context.Background(), creqteQuizRequest)

	uidPrimitive, _ := uid.(primitive.Binary)
	userId, _ := uuid.FromBytes(uidPrimitive.Data)

	g.JSON(201, gin.H{"_id": userId})

}

func (r resource) get(g *gin.Context) {
	g.JSON(200, "")
}
