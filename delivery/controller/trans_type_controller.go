package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/model"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type TransTypeController struct {
	router      *gin.Engine
	ucTransType usecase.TransTypeUseCase
	api.BaseApi
}

func (t *TransTypeController) transTypeAdd(c *gin.Context) {
	var newTransTypeData dto.TransTypeAddDto
	err := t.ParseRequestBody(c, &newTransTypeData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTransType.CreateTransType(&model.TransType{Id: newTransTypeData.Id, Description: newTransTypeData.Description})
	if err != nil {
		t.Failed(c, err)
		return
	}
	result, _ := t.ucTransType.GetLastestTransTypeData()
	t.Success(c, result)
}

func (t *TransTypeController) transTypeFindById(c *gin.Context) {
	var newTransTypeData dto.TransTypeIdDto
	err := t.ParseRequestBody(c, &newTransTypeData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	menuRes, err := t.ucTransType.FindTransTypeById(newTransTypeData.Id)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, menuRes)
}

func (t *TransTypeController) transTypeDeleteById(c *gin.Context) {
	var newTransTypeData dto.TransTypeIdDto
	err := t.ParseRequestBody(c, &newTransTypeData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTransType.DeleteTransTypeById(&model.TransType{Id: newTransTypeData.Id})
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, "Trans Type Deleted")
}

func (t *TransTypeController) updateTransTypeUsingID(c *gin.Context) {
	var newTransTypeData dto.TransTypeAddDto
	err := t.ParseRequestBody(c, &newTransTypeData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTransType.UpdateTransTypeDescription(&model.TransType{Id: newTransTypeData.Id}, map[string]interface{}{
		"description": newTransTypeData.Description,
	})

	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, "Trans Type Updated")
}

func NewTransTypeController(router *gin.Engine, ucTransType usecase.TransTypeUseCase) *TransTypeController {

	controller := TransTypeController{
		router:      router,
		ucTransType: ucTransType,
	}

	rCust := router.Group("/transtype")
	{
		rCust.POST("/add", controller.transTypeAdd)
		rCust.GET("/find", controller.transTypeFindById)
		rCust.DELETE("/delete", controller.transTypeDeleteById)
		rCust.PUT("/update", controller.updateTransTypeUsingID)
	}
	return &controller

}
