package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/model"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type DiscountController struct {
	router *gin.Engine
	ucDisc usecase.DiscountUseCase
	api.BaseApi
}

func (d *DiscountController) discountAdd(c *gin.Context) {
	var newDiscData dto.DiscAddDto
	err := d.ParseRequestBody(c, &newDiscData)
	if err != nil {
		d.Failed(c, utils.RequiredError())
		return
	}
	err = d.ucDisc.CreateDiscount(&model.Discount{Description: newDiscData.Description, Pct: newDiscData.Pct})
	if err != nil {
		d.Failed(c, err)
		return
	}
	result, _ := d.ucDisc.GetLastestDiscountData()
	d.Success(c, result)
}

func (d *DiscountController) discountFindById(c *gin.Context) {
	var newDisceData dto.DiscIdDto
	err := d.ParseRequestBody(c, &newDisceData)
	if err != nil {
		d.Failed(c, utils.RequiredError())
		return
	}
	result, err := d.ucDisc.FindDiscountById(newDisceData.Id)
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, result)
}

func (d *DiscountController) discountDeleteById(c *gin.Context) {
	var newDisceData dto.DiscIdDto
	err := d.ParseRequestBody(c, &newDisceData)
	if err != nil {
		d.Failed(c, utils.RequiredError())
		return
	}
	err = d.ucDisc.DeleteDiscountById(&model.Discount{Id: newDisceData.Id})
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, "Discount Data Deleted")
}

func (d *DiscountController) updateDiscountUsingID(c *gin.Context) {
	var newDiscData dto.DiscUpdtDto
	err := d.ParseRequestBody(c, &newDiscData)
	if err != nil {
		d.Failed(c, utils.RequiredError())
		return
	}
	err = d.ucDisc.UpdateDiscountById(&model.Discount{Id: newDiscData.Id}, map[string]interface{}{
		"description": newDiscData.Description,
		"pct":         newDiscData.Pct,
	})

	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, "Discount Data Updated")
}

func NewDiscountController(router *gin.Engine, ucDisc usecase.DiscountUseCase) *DiscountController {

	controller := DiscountController{
		router: router,
		ucDisc: ucDisc,
	}

	rCust := router.Group("/discount")
	{
		rCust.POST("/add", controller.discountAdd)
		rCust.GET("/find", controller.discountFindById)
		rCust.DELETE("/delete", controller.discountDeleteById)
		rCust.PUT("/update", controller.updateDiscountUsingID)
	}
	return &controller

}
