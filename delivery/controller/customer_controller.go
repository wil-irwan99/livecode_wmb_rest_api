package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router       *gin.Engine
	ucCustRegis  usecase.CustomerRegistration
	ucCustMember usecase.MemberActivation
	ucCust       usecase.CustomerUseCase
	api.BaseApi
}

func (v *CustomerController) customerRegistration(c *gin.Context) {
	var newCustData dto.CustRegisDto
	err := v.ParseRequestBody(c, &newCustData)
	if err != nil {
		v.Failed(c, utils.RequiredError())
		return
	}
	err = v.ucCustRegis.RegistrasiCust(newCustData.Name, newCustData.PhoneNo)
	if err != nil {
		v.Failed(c, err)
		return
	}
	custRes, _ := v.ucCust.GetLastDataCustomerList()
	v.Success(c, custRes)
}

func (v *CustomerController) customerMemberAct(c *gin.Context) {
	var newCustData dto.CustMemberActDto
	err := v.ParseRequestBody(c, &newCustData)
	if err != nil {
		v.Failed(c, utils.RequiredError())
		return
	}

	err = v.ucCustMember.AktivasiMember(newCustData.Id, newCustData.DiscId)
	if err != nil {
		v.Failed(c, err)
		return
	}
	custRes, _ := v.ucCust.FindCustomerById(newCustData.Id)
	v.Success(c, custRes)
}

func NewCustomerController(router *gin.Engine, ucCustRegis usecase.CustomerRegistration, ucCustMember usecase.MemberActivation, ucCust usecase.CustomerUseCase) *CustomerController {

	controller := CustomerController{
		router:       router,
		ucCustRegis:  ucCustRegis,
		ucCustMember: ucCustMember,
		ucCust:       ucCust,
	}

	rCust := router.Group("/customer")
	{
		rCust.POST("/registration", controller.customerRegistration)
		rCust.POST("/memberactivation", controller.customerMemberAct)
	}
	return &controller

}
