package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/model"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	router      *gin.Engine
	ucMenu      usecase.MenuUseCase
	ucMenuPrice usecase.MenuPriceUseCase
	api.BaseApi
}

func (m *MenuController) menuAdd(c *gin.Context) {
	var newMenuData dto.MenuAddDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenu.CreateNewMenu(&model.Menu{MenuName: newMenuData.Name})
	if err != nil {
		m.Failed(c, err)
		return
	}
	menuRes, _ := m.ucMenu.GetLastestMenuData()
	m.Success(c, menuRes)
}

func (m *MenuController) menuFindById(c *gin.Context) {
	var newMenuData dto.MenuIdDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	menuRes, err := m.ucMenu.FindMenuById(newMenuData.Id)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, menuRes)
}

func (m *MenuController) deleteMenuUsingID(c *gin.Context) {
	var newMenuData dto.MenuIdDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenu.DeleteMenuByID(&model.Menu{Id: newMenuData.Id})
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, "Menu Deleted")
}

func (m *MenuController) updateMenuUsingID(c *gin.Context) {
	var newMenuData dto.MenuUpdtDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenu.UpdateMenuByID(&model.Menu{Id: newMenuData.Id}, map[string]interface{}{
		"menu_name": newMenuData.Name,
	})

	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, "Menu Updated")
}

func (m *MenuController) menuPriceAdd(c *gin.Context) {
	var newMenuData dto.MenuPriceAddDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenuPrice.CreateNewMenuPrice(&model.MenuPrice{MenuID: newMenuData.Id, Price: newMenuData.Price})
	if err != nil {
		m.Failed(c, err)
		return
	}
	menuRes, _ := m.ucMenuPrice.GetLastestMenuPrice()
	m.Success(c, menuRes)
}

func (m *MenuController) menuPriceFindById(c *gin.Context) {
	var newMenuData dto.MenuIdDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	menuRes, err := m.ucMenuPrice.FindMenuPriceById(newMenuData.Id)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, menuRes)
}

func (m *MenuController) deleteMenuPriceUsingID(c *gin.Context) {
	var newMenuData dto.MenuIdDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenuPrice.DeleteMenuPriceById(&model.MenuPrice{Id: newMenuData.Id})
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, "Menu Price Deleted")
}

func (m *MenuController) updateMenuPriceUsingID(c *gin.Context) {
	var newMenuData dto.MenuPriceUpdtDto
	err := m.ParseRequestBody(c, &newMenuData)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenuPrice.UpdateMenuPriceByID(&model.MenuPrice{Id: newMenuData.Id}, map[string]interface{}{
		"price": newMenuData.Price,
	})

	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, "Menu Price Updated")
}

func NewMenuController(router *gin.Engine, ucMenu usecase.MenuUseCase, ucMenuPrice usecase.MenuPriceUseCase) *MenuController {
	controller := MenuController{
		router:      router,
		ucMenu:      ucMenu,
		ucMenuPrice: ucMenuPrice,
	}

	rCust := router.Group("/menu")
	{
		rCust.POST("/add", controller.menuAdd)
		rCust.PUT("/update", controller.updateMenuUsingID)
		rCust.GET("/find", controller.menuFindById)
		rCust.DELETE("/delete", controller.deleteMenuUsingID)
	}

	rCust2 := router.Group("/menu/price")
	{
		rCust2.POST("/add", controller.menuPriceAdd)
		rCust2.PUT("/update", controller.updateMenuPriceUsingID)
		rCust2.GET("/find", controller.menuPriceFindById)
		rCust2.DELETE("/delete", controller.deleteMenuPriceUsingID)
	}
	return &controller

}
