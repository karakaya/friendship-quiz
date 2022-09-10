package route

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	res := resource{}
	g.GET("/{id}", res.get)
	g.POST("/", res.create)

}

type resource struct{}

func (r resource) create(g *gin.Context) {
	g.JSON(200, "")
}

func (r resource) get(g *gin.Context) {
	g.JSON(200, "")
}
