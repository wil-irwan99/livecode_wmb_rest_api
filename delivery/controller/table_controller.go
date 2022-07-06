package controller

import (
	"wmb_rest_api/delivery/api"
	"wmb_rest_api/dto"
	"wmb_rest_api/model"
	"wmb_rest_api/usecase"
	"wmb_rest_api/utils"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	router  *gin.Engine
	ucTable usecase.TableUseCase
	api.BaseApi
}

func (t *TableController) tableAdd(c *gin.Context) {
	var newTableData dto.TableAddDto
	err := t.ParseRequestBody(c, &newTableData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTable.CreateTable(&model.Table{TableDescription: newTableData.TableDescription, IsAvailable: newTableData.IsAvailable})
	if err != nil {
		t.Failed(c, err)
		return
	}
	tableRes, _ := t.ucTable.GetLastestTableData()
	t.Success(c, tableRes)
}

func (t *TableController) tableFindById(c *gin.Context) {
	var newTableData dto.TableIdDto
	err := t.ParseRequestBody(c, &newTableData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	menuRes, err := t.ucTable.FindTableById(newTableData.Id)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, menuRes)
}

func (t *TableController) tableDeleteById(c *gin.Context) {
	var newTableData dto.TableIdDto
	err := t.ParseRequestBody(c, &newTableData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTable.DeleteTableByID(&model.Table{Id: newTableData.Id})
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, "Table Deleted")
}

func (t *TableController) updateTableUsingID(c *gin.Context) {
	var newTableData dto.TableUpdtDto
	err := t.ParseRequestBody(c, &newTableData)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTable.UpdateTableAvailable(&model.Table{Id: newTableData.Id}, map[string]interface{}{
		"table_description": newTableData.TableDescription,
		"is_available":      newTableData.IsAvailable,
	})

	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, "Table Updated")
}

func NewTableController(router *gin.Engine, ucTable usecase.TableUseCase) *TableController {

	controller := TableController{
		router:  router,
		ucTable: ucTable,
	}

	rCust := router.Group("/table")
	{
		rCust.POST("/add", controller.tableAdd)
		rCust.GET("/find", controller.tableFindById)
		rCust.DELETE("/delete", controller.tableDeleteById)
		rCust.PUT("/update", controller.updateTableUsingID)
	}
	return &controller

}
