package usecase

import (
	"log"
	"time"
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type TableValidation interface {
	ValidasiMeja(id_cust int, id_transType string, id_meja int, id_menuPrice []int, qty []int) error
}

type tableValidation struct {
	custRepo       repository.CustomerRepository
	transTypeRepo  repository.TransTypeRepository
	tableRepo      repository.TableRepository
	billRepo       repository.BillRepository
	menuPriceRepo  repository.MenuPriceRepository
	billDetailRepo repository.BillDetailRepository
}

func (t *tableValidation) ValidasiMeja(id_cust int, id_transType string, id_meja int, id_menuPrice []int, qty []int) error {

	_, err := t.custRepo.FindById(id_cust)
	if err != nil {
		return err
	}

	_, err = t.transTypeRepo.FindById(id_transType)
	if err != nil {
		return err
	}

	tableExist, err := t.tableRepo.FindById(id_meja)
	if err != nil {
		return err
	}

	if id_transType == "EI" {
		if tableExist.IsAvailable == false {
			log.Println("table full...")
			return err
		} else {
			t.tableRepo.UpdateAvailable(&tableExist, map[string]interface{}{
				"is_available": false,
			})
		}
	} else {
		tableExist.Id = 1 //Id meja = 1 sengaja untuk TA, deskripsi meja id 1 adalah "bungkus"
	}

	for _, v := range id_menuPrice {
		_, err = t.menuPriceRepo.FindById(v)
		if err != nil {
			return err
		}
	}

	timeNow := time.Now()

	bill := model.Bill{
		TransDate:   timeNow,
		CustomerID:  id_cust,
		TableID:     tableExist.Id,
		TransTypeID: id_transType,
	}

	t.billRepo.Create(&bill)

	lastedBillCreated, err := t.billRepo.RetreiveLastedBillID()
	if err != nil {
		return err
	}

	for i, v := range id_menuPrice {

		billDetail := model.BillDetail{
			BillID:      lastedBillCreated.Id,
			MenuPriceID: v,
			Qty:         qty[i],
		}

		err = t.billDetailRepo.Create(&billDetail)
		if err != nil {
			return err
		}

	}

	return nil
}

func NewTableValidation(custRepo repository.CustomerRepository, transTypeRepo repository.TransTypeRepository, tableRepo repository.TableRepository, billRepo repository.BillRepository, menuPriceRepo repository.MenuPriceRepository, billDetailRepo repository.BillDetailRepository) TableValidation {
	return &tableValidation{
		custRepo:       custRepo,
		transTypeRepo:  transTypeRepo,
		tableRepo:      tableRepo,
		billRepo:       billRepo,
		menuPriceRepo:  menuPriceRepo,
		billDetailRepo: billDetailRepo,
	}
}
