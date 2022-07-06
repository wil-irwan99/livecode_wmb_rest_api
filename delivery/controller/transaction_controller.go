package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router          *gin.Engine
	ucTableValTrans usecase.TableValidation
	api.BaseApi
}

func (t *TransactionController) transactionValidation(c *gin.Context) {
	var newTransData dto.TransAddDto
	err := t.ParseRequestBody(c, &newTransData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTableValTrans.ValidasiMeja(newTransData.CustId, newTransData.TransTypeId, newTransData.MejaId, newTransData.MenuPriceId, newTransData.Qty)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, "Transaction Success")
}

func NewTransactionController(router *gin.Engine, ucTableValTrans usecase.TableValidation) *TransactionController {

	controller := TransactionController{
		router:          router,
		ucTableValTrans: ucTableValTrans,
	}

	rCust := router.Group("/transaction")
	{
		rCust.POST("/add", controller.transactionValidation)
	}
	return &controller

}
