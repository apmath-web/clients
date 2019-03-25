package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/clients/Application/v1/validation"
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validator := validation.GenValidation()
		validator.SetMessage("validation error")
		validator.AddMessage(validation.GenMessage("id", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	vm := new(viewModels.ClientViewModel)
	if err := c.BindJSON(vm); err != nil {
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
	fmt.Println(id)
	c.String(http.StatusNoContent, "")

}
