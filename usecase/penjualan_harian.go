package usecase

import (
	"time"
	"wmb_rest_api/dto"
	"wmb_rest_api/repository"
)

type DailySell interface {
	PrintBillDaily() (dto.BillHarian, error)
}

type dailySell struct {
	billDetailRepo repository.BillDetailRepository
	menuPriceRepo  repository.MenuPriceRepository
}

func (d *dailySell) PrintBillDaily() (dto.BillHarian, error) {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	var laporanHarian dto.BillHarian

	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)
	result, err := d.billDetailRepo.FindByDate(yesterday, tomorrow)
	if err != nil {
		return dto.BillHarian{}, err
	}

	total := 0
	laporanHarian.Today = today
	for _, v := range result {
		menuPriceExist, err := d.menuPriceRepo.FindById(v.MenuPriceID)
		if err != nil {
			return dto.BillHarian{}, err
		}
		total += v.MenuPrice.Price * v.Qty
		laporanHarian.Orders = append(laporanHarian.Orders, dto.OrderList{
			Name:  menuPriceExist.Menu.MenuName,
			Qty:   v.Qty,
			Price: v.MenuPrice.Price,
		})
	}
	laporanHarian.TotalPrice = float32(total)

	return dto.BillHarian{}, nil
}

func NewDailySell(billDetailRepo repository.BillDetailRepository, menuPriceRepo repository.MenuPriceRepository) DailySell {
	return &dailySell{
		billDetailRepo: billDetailRepo,
		menuPriceRepo:  menuPriceRepo,
	}
}
