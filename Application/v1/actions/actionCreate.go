package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}
