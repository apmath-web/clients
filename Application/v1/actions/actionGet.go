package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/clients/Application/v1/Validation"
	"github.com/apmath-web/clients/Application/v1/viewModels"
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
	fmt.Println(id)
	vm := new(viewModels.ClientViewModel)
	mockString := []byte(`{"firstName":"Lev","lastName":"Kovalenko","birthDate":"2019-01-02","Passport":{"series":2342,"number":434345},"Jobs":[{"name":"test1","wage":536446},{"name":"test4","wage":53644346}],"sex":"male","maritalStatus":"married","children":4}`)
	_ = json.Unmarshal(mockString, vm)
	c.JSON(http.StatusOK, vm)
}
