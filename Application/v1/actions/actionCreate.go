package actions

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/Validation"
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	vm := new(viewModels.ClientViewModel)
	if err := c.BindJSON(vm); err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("Validation error")
		validator.AddMessage(Validation.GenMessage("json", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	if !vm.Validate() {
		validator := vm.GetValidation()
		validator.SetMessage("Validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": 1})
}
