package actions

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/validation"
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/apmath-web/clients/Domain/models"
	"github.com/apmath-web/clients/Infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validator := validation.GenValidation()
		validator.SetMessage("validation error")
		validator.AddMessage(validation.GenMessage("id", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	IdModel := models.GenId(id)
	service := Infrastructure.GetServiceManager().GetClientService()
	dm, err := service.Get(IdModel)
	if err != nil {
		validator := validation.GenValidation()
		validator.SetMessage(err.Error())
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	vm := new(viewModels.ClientViewModel)
	vm.Hydrate(dm)
	c.JSON(http.StatusOK, vm)
}
