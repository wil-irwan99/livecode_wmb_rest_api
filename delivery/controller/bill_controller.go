package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type BillController struct {
	router      *gin.Engine
	ucPrintBill usecase.CetakBill
	api.BaseApi
}

func (b *BillController) PrintBillById(c *gin.Context) {
	var billIdData dto.PrintBillId
	err := b.ParseRequestBody(c, &billIdData)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	result, err := b.ucPrintBill.PrintBill(billIdData.BillId)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, result)
}

func NewBillController(router *gin.Engine, ucPrintBill usecase.CetakBill) *BillController {

	controller := BillController{
		router:      router,
		ucPrintBill: ucPrintBill,
	}

	rCust := router.Group("/bill")
	{
		rCust.GET("/print", controller.PrintBillById)
	}
	return &controller

}
