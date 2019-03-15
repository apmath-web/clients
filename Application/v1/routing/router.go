package routing

import (
	"github.com/apmath-web/clients/Application/v1/actions"
	"github.com/gin-gonic/gin"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/", actions.Create)
		v1.PUT("/:id", actions.Update)
	}
	return router
}
