package actions

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/Validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("Validation error")
		validator.AddMessage(Validation.GenMessage("id", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

}
