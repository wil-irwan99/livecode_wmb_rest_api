package delivery

import (
	"wmb_rest_api/config"
	"wmb_rest_api/delivery/controller"
	"wmb_rest_api/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	managerUscs manager.UseCaseManager
	engine      *gin.Engine
	host        string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(&appConfig)
	managerRepo := manager.NewRepositoryManager(infra)
	managerUseCase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.Url
	return &appServer{
		managerUscs: managerUseCase,
		engine:      r,
		host:        host,
	}
}

func (a *appServer) initControllers() {
	controller.NewCustomerController(a.engine, a.managerUscs.CustomerRegistration(), a.managerUscs.MemberActivation(), a.managerUscs.CustomerUseCase())
	controller.NewMenuController(a.engine, a.managerUscs.MenuUseCase(), a.managerUscs.MenuPriceUseCase())
	controller.NewTableController(a.engine, a.managerUscs.TableUseCase())
	controller.NewTransTypeController(a.engine, a.managerUscs.TransTypeUseCase())
	controller.NewDiscountController(a.engine, a.managerUscs.DiscountUseCase())
	controller.NewTransactionController(a.engine, a.managerUscs.TableValidation())
	controller.NewBillController(a.engine, a.managerUscs.CetakBill())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
