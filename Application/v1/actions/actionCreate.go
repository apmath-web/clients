package actions

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/mapper"
	"github.com/apmath-web/clients/Application/v1/validation"
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/apmath-web/clients/Infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	vm := viewModels.ClientViewModel{}
	if err := c.BindJSON(&vm); err != nil {
		validator := validation.GenValidation()
		validator.SetMessage("validation error")
		validator.AddMessage(validation.GenMessage("json", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	if !vm.Validate() {
		validator := vm.GetValidation()
		validator.SetMessage("validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	dm := mapper.ClientViewMapper(vm)
	service := Infrastructure.GetServiceManager().GetClientService()
	id, err := service.Add(dm)
	if err != nil {
		validator := validation.GenValidation()
		validator.SetMessage(err.Error())
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id.Get()})
}
