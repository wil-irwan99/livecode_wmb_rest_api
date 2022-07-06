package main

import "wmb_rest_api/delivery"

func main() {

	//cfg := config.NewConfig()
	//cfg.DbConn()
	// db := cfg.DbConn()
	// defer cfg.DbClose()

	// billDetailRepo := repository.NewBillDetailRepository(db)
	// tableRepo := repository.NewTableRpository(db)
	// menuRepo := repository.NewMenuRepository(db)
	// menuPriceRepo := repository.NewMenuPriceRepository(db)
	// customerRepo := repository.NewCustomerRepository(db)
	// billRepo := repository.NewBillRepository(db)

	// cetakBillUsecase := usecase.NewCetakBill(billDetailRepo, tableRepo, menuPriceRepo, menuRepo, customerRepo, billRepo)
	// cetakBillUsecase.PrintBill(2)

	// custRepo := repository.NewCustomerRepository(db)
	// transTypeRepo := repository.NewTransTypeRepository(db)
	// tableRepo := repository.NewTableRpository(db)
	// billRepo := repository.NewBillRepository(db)
	// menuPriceRepo := repository.NewMenuPriceRepository(db)
	// billDetailRepo := repository.NewBillDetailRepository(db)
	// tableValidUscs := usecase.NewTableValidation(custRepo, transTypeRepo, tableRepo, billRepo, menuPriceRepo, billDetailRepo)
	// menuPriceSlice := []int{1, 2}
	// menuQtySlice := []int{1, 1}
	// tableValidUscs.ValidasiMeja(2, "EI", 1, menuPriceSlice, menuQtySlice)

	delivery.Server().Run()
}
